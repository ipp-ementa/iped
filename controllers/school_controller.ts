import {
  AvailableSchools,
  Error,
  AvailableSchoolsItem,
  InternalServerError,
  NotFound,
  CreateSchool,
  CreatedSchool,
  BadRequest,
  CreatedCanteen,
  DetailedSchoolInformation,
} from "../views/views.ts";

import {
  Err,
  Ok,
  Result,
  School,
  Option,
  NonEmptyString,
  Canteen,
  GeographicalLocation,
  ObjectId,
  Some,
  None,
} from "../models/models.ts";

import { Collection, QuerySelector } from "../deps.ts";

export async function availableSchools(
  repository: SchoolRepository,
): Promise<Result<AvailableSchools, Error>> {
  const availableSchools = await repository.schools();

  if (availableSchools.isOk()) {
    const schools = availableSchools.unwrap();

    if (schools.length > 0) {
      const availableSchoolsView = schools.map<AvailableSchoolsItem>(
        function (s): AvailableSchoolsItem {
          return {
            id: s.id,
            acronym: s.acronym,
            name: s.name,
          };
        },
      );
      return Ok(availableSchoolsView);
    } else {
      return Err(new NotFound());
    }
  } else {
    return Err(new InternalServerError());
  }
}

export async function createSchool(
  repository: SchoolRepository,
  schoolToCreate: CreateSchool,
): Promise<Result<CreatedSchool, Error>> {
  const schoolFoundByAcronym = await repository.school(
    { acronym: schoolToCreate.acronym },
  );

  if (schoolFoundByAcronym.isOk() && schoolFoundByAcronym.unwrap().isNone()) {
    const acronym = NonEmptyString.create(schoolToCreate.acronym);

    const name = NonEmptyString.create(schoolToCreate.name);

    const canteens = schoolToCreate.canteens.map(
      function (c): Result<Canteen, string> {
        const location = GeographicalLocation.create(
          c.location.latitude,
          c.location.longitude,
        );

        const name = NonEmptyString.create(c.name);

        if (location.isErr()) {
          return Err(location.unwrapErr());
        } else if (name.isErr()) {
          return Err(name.unwrapErr());
        } else {
          return Ok(Canteen.create(location.unwrap(), name.unwrap()));
        }
      },
    );

    const results = [acronym, ...canteens, name];

    const firstError = results.find((r) => r.isErr());

    if (firstError) {
      return Err(
        new BadRequest(
          firstError.unwrapErr(),
        ),
      );
    } else {
      const school = School.create(
        acronym.unwrap(),
        canteens.map((r) => r.unwrap()),
        name.unwrap(),
      );

      if (school.isErr()) {
        return Err(
          new BadRequest(
            school.unwrapErr(),
          ),
        );
      }

      const createdSchool = await repository.create(
        school.unwrap(),
      );

      if (createdSchool.isErr()) {
        return Err(new InternalServerError());
      } else {
        const unwrapSchool = createdSchool.unwrap();

        const createdCanteens = unwrapSchool.canteens.map(
          function (c): CreatedCanteen {
            return {
              id: c.name,
              location: c.location,
              name: c.name,
            };
          },
        );
        return Ok(
          <CreatedSchool> {
            id: unwrapSchool.id,
            acronym: unwrapSchool.acronym,
            name: unwrapSchool.name,
            canteens: createdCanteens,
          },
        );
      }
    }
  } else {
    return Err(
      new BadRequest("School already exists"),
    );
  }
}

export async function detailedSchoolInformation(
  repository: SchoolRepository,
  id: string,
): Promise<Result<DetailedSchoolInformation, Error>> {
  const schoolResult = await querySchoolById(repository, id);

  if (schoolResult.isErr()) {
    return Err(schoolResult.unwrapErr());
  } else {
    const school = schoolResult.unwrap();

    const schoolCanteensAsCreatedCanteens = school.canteens.map(
      function (c): CreatedCanteen {
        return {
          id: c.name,
          location: c.location,
          name: c.name,
        };
      },
    );

    return Ok(
      <DetailedSchoolInformation> {
        id: school.id,
        acronym: school.acronym,
        canteens: schoolCanteensAsCreatedCanteens,
        name: school.name,
      },
    );
  }
}

export async function querySchoolById(
  repository: SchoolRepository,
  id: string,
): Promise<Result<School, Error>> {
  const schoolQuery = <SchoolQuery> { id: id };

  const schoolResult = await repository.school(schoolQuery);

  if (schoolResult.isErr()) {
    return Err(new InternalServerError());
  } else {
    const schoolOption = schoolResult.unwrap();

    if (schoolOption.isNone()) {
      return Err(new NotFound());
    } else {
      const school = schoolOption.unwrap();

      return Ok(school);
    }
  }
}

export interface SchoolRepository {
  create(school: School): Promise<Result<School, Error>>;
  schools(): Promise<Result<School[], Error>>;
  school(query: SchoolQuery): Promise<Result<Option<School>, Error>>;
  update(school: School): Promise<Result<School, Error>>;
}

export class MongoSchoolRepository implements SchoolRepository {
  collection: Collection<School>;

  public constructor(collection: Collection<School>) {
    this.collection = collection;
  }

  async create(school: School): Promise<Result<School, Error>> {
    try {
      const result = await this.collection.insertOne(school);

      school._id = result;

      return Ok(school);
    } catch (error) {
      return Err(new InternalServerError());
    }
  }

  async schools(): Promise<Result<School[], Error>> {
    try {
      const result = await this.collection.find({ name: { $ne: null } });

      return Ok(result);
    } catch (error) {
      return Err(new InternalServerError());
    }
  }
  async school(query: SchoolQuery): Promise<Result<Option<School>, Error>> {
    try {
      const mongoQuery = {
        _id: <QuerySelector<ObjectId>> {},
        acronym: {
          $eq: query.acronym,
        },
      };

      if (query.id) {
        mongoQuery._id = {
          $eq: ObjectId(query.id),
        };
      } else {
        delete mongoQuery._id;
      }

      const result = await this.collection.findOne(mongoQuery);

      if (result) {
        return Ok(Some(result));
      } else {
        return Ok(None);
      }
    } catch (error) {
      return Err(new InternalServerError());
    }
  }

  async update(school: School): Promise<Result<School, Error>> {
    try {
      const result = await this.collection.updateOne({
        _id: school._id,
      }, {
        $set: {
          acronym: school.acronym,
          canteens: school.canteens,
          name: school.name,
        },
      });

      if (result.modifiedCount > 0) {
        return Ok(result);
      } else {
        return Err(new InternalServerError());
      }
    } catch (error) {
      return Err(new InternalServerError());
    }
  }
}

export interface SchoolQuery {
  id?: string;
  acronym?: string;
}
