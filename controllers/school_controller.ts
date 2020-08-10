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

import { Collection } from "../deps.ts";

export async function availableSchools(
  repository: SchoolRepository,
): Promise<Result<AvailableSchools, Error>> {
  const availableSchools = await repository.schools();

  if (availableSchools.isOk()) {
    const schools = availableSchools.unwrap();

    if (schools.length > 0) {
      const availableSchoolsView = schools.map<AvailableSchoolsItem>(
        function (s, i, a): AvailableSchoolsItem {
          return {
            id: s._id.$oid,
            acronym: s.acronym.valueOf(),
            name: s.name.valueOf(),
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
  const schoolFoundByName = await repository.school(
    { name: schoolToCreate.name },
  );

  if (schoolFoundByName.isErr()) {
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
              id: c.name.valueOf(),
              location: c.location,
              name: c.name.valueOf(),
            };
          },
        );
        return Ok(
          <CreatedSchool> {
            id: unwrapSchool._id.$oid,
            acronym: unwrapSchool.acronym.valueOf(),
            name: unwrapSchool.name.valueOf(),
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
          id: c.name.valueOf(),
          location: c.location,
          name: c.name.valueOf(),
        };
      },
    );

    return Ok(
      <DetailedSchoolInformation> {
        id: school._id.$oid,
        acronym: school.acronym.valueOf(),
        canteens: schoolCanteensAsCreatedCanteens,
        name: school.name.valueOf(),
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
        _id: <ObjectId> {},
        name: <String> {},
      };

      if (query.id) {
        mongoQuery._id = ObjectId(query.id);
      } else if (query.name) {
        mongoQuery.name = query.name;
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
  name?: string;
}
