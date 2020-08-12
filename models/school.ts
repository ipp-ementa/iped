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
    acronym: string,
    canteens: Canteen[],
    name: string,
    id?: { $oid: string },
  ) {
    this._id = id;
    this.acronym = acronym;
    this.canteens = canteens;
    this.name = name;
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

  public updateCanteen(canteen: Canteen): Result<void, string> {
    const canteenFiltered = this.canteens.filter((c) => c.name == canteen.name);

    const hasCanteen = canteenFiltered.length < this.canteens.length;

    if (hasCanteen) {
      canteenFiltered.push(canteen);
      return Ok(undefined);
    } else {
      return Err("canteen not found");
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
      return Ok(new School(acronym.valueOf(), canteens, name.valueOf()));
    }
  }

  public static fromJson(object: any): School {
    return new School(
      object.acronym,
      object.canteens.map((c: Canteen) => Canteen.fromJson(c)),
      object.name,
      object._id,
    );
  }
}

export { School };
