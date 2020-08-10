import { Canteen } from "./canteen.ts";

import {
  HasDuplicates,
  Err,
  NonEmptyString,
  Ok,
  Result,
} from "./common.ts";

class School {
  public get id(): string {
    return this._id?.$oid || "undefined";
  }

  _id?: { $oid: string };

  acronym: string;

  canteens: Canteen[];

  name: string;

  private constructor(
    acronym: NonEmptyString,
    canteens: Canteen[],
    name: NonEmptyString,
  ) {
    this._id = undefined;
    this.acronym = acronym.valueOf();
    this.canteens = canteens;
    this.name = name.valueOf();
  }

  public addCanteen(canteen: Canteen): Result<void, string> {
    const canteensCopy = this.canteens.slice();
    canteensCopy.push(canteen);

    if (HasDuplicates<Canteen>(canteensCopy)) {
      return Err("cannot add duplicate canteen");
    } else {
      this.canteens.push(canteen);
      return Ok(undefined);
    }
  }

  public static create(
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

export { School };
