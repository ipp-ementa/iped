import { Canteen } from "./canteen.ts";

import { HasDuplicates, Err, NonEmptyString, Ok, Result } from "./common.ts";

class School {
  acronym: NonEmptyString;

  canteens: Canteen[];

  name: NonEmptyString;

  constructor(
    acronym: NonEmptyString,
    canteens: Canteen[],
    name: NonEmptyString,
  ) {
    this.acronym = acronym;
    this.canteens = canteens;
    this.name = name;
  }

  static create(
    acronym: NonEmptyString,
    canteens: Canteen[],
    name: NonEmptyString,
  ): Result<School, string> {
    if (canteens.length == 0) {
      return Err("must provide at least one canteen");
    } else if (HasDuplicates<Canteen>(canteens)) {
      return Err("cannot have duplicate canteen");
    } else {
      return Ok(new School(acronym, canteens, name));
    }
  }
}

export { Canteen };
