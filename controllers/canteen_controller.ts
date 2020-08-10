import {
  SchoolRepository,
  querySchoolById,
} from "./school_controller.ts";

import {
  AvailableCanteens,
  Error,
  InternalServerError,
  NotFound,
  AvailableCanteensItem,
  CreatedCanteen,
  CreateCanteen,
  BadRequest,
  DetailedCanteenInformation,
} from "../views/views.ts";

import {
  Err,
  Ok,
  Result,
  Canteen,
  GeographicalLocation,
  NonEmptyString,
} from "../models/models.ts";

export async function availableCanteens(
  schoolRepository: SchoolRepository,
  schoolId: string,
): Promise<Result<AvailableCanteens, Error>> {
  const schoolResult = await querySchoolById(schoolRepository, schoolId);

  if (schoolResult.isErr()) {
    return Err(schoolResult.unwrapErr());
  } else {
    const school = schoolResult.unwrap();

    const canteens = school.canteens;

    const availableCanteensView = canteens.map<AvailableCanteensItem>(
      (function (c): AvailableCanteensItem {
        return {
          id: c.name.valueOf(),
          name: c.name.valueOf(),
        };
      }),
    );

    return Ok(availableCanteensView);
  }
}

export async function createCanteen(
  schoolRepository: SchoolRepository,
  schoolId: string,
  canteenToCreate: CreateCanteen,
): Promise<Result<CreatedCanteen, Error>> {
  const schoolResult = await querySchoolById(schoolRepository, schoolId);

  if (schoolResult.isErr()) {
    return Err(schoolResult.unwrapErr());
  } else {
    const school = schoolResult.unwrap();

    const canteenLocation = GeographicalLocation.create(
      canteenToCreate.location.latitude,
      canteenToCreate.location.longitude,
    );

    const canteenName = NonEmptyString.create(canteenToCreate.name);

    const results = [canteenLocation, canteenName];

    const firstError = results.find((r) => r.isErr());

    if (firstError) {
      return Err(
        <BadRequest> ({
          message: firstError.unwrapErr(),
        }),
      );
    } else {
      const canteen = Canteen.create(
        canteenLocation.unwrap(),
        canteenName.unwrap(),
      );

      const addCanteenResult = school.addCanteen(canteen);

      if (addCanteenResult.isErr()) {
        return Err(
          <BadRequest> ({
            message: addCanteenResult.unwrapErr(),
          }),
        );
      } else {
        const schoolUpdateResult = await schoolRepository.update(school);

        if (schoolUpdateResult.isErr()) {
          return Err(<InternalServerError> {});
        } else {
          const updatedSchool = schoolUpdateResult.unwrap();

          const createdCanteen = updatedSchool.canteens.find((c) =>
            c.name == canteen.name
          );

          const createdCanteenView = <CreatedCanteen> {
            id: createdCanteen?.name.valueOf(),
            location: createdCanteen?.location,
            name: createdCanteen?.name.valueOf(),
          };

          return Ok(createdCanteenView);
        }
      }
    }
  }
}

export async function detailedCanteenInformation(
  schoolRepository: SchoolRepository,
  schoolId: string,
  canteenId: string,
): Promise<Result<DetailedCanteenInformation, Error>> {
  const canteenResult = await queryCanteenById(
    schoolRepository,
    schoolId,
    canteenId,
  );

  if (canteenResult.isErr()) {
    return Err(canteenResult.unwrapErr());
  } else {
    const canteen = canteenResult.unwrap();

    const detailedCanteenInformationView = <DetailedCanteenInformation> {
      id: canteen.name.valueOf(),
      location: canteen.location,
      name: canteen.name.valueOf(),
    };

    return Ok(detailedCanteenInformationView);
  }
}

export async function queryCanteenById(
  repository: SchoolRepository,
  schoolId: string,
  canteenId: string,
): Promise<Result<Canteen, Error>> {
  const schoolResult = await querySchoolById(repository, schoolId);

  if (schoolResult.isErr()) {
    return Err(schoolResult.unwrapErr());
  } else {
    const school = schoolResult.unwrap();

    const canteen = school.canteens.find((c) => c);

    if (canteen) {
      return Ok(canteen);
    } else {
      return Err(<NotFound> {});
    }
  }
}
