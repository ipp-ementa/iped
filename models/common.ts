import { Err, None, ObjectId, Ok, Option, Result, Some } from "../deps.ts";

class NonEmptyString extends String {
  public constructor(value: string) {
    super(value);
  }

  public static create(value: string): Result<NonEmptyString, string> {
    if (value?.trim().length > 0) {
      return Ok(new NonEmptyString(value));
    } else {
      return Err("string cannot be empty");
    }
  }
}

interface Equatable {
  equals(obj: any): boolean;
}

function HasDuplicates<T extends Equatable>(array: T[]) {
  if (array.length < 2) {
    return false;
  }

  for (let i = 0; i < array.length; i++) {
    for (let j = i + 1; j < array.length; j++) {
      if (array[i].equals(array[j])) {
        return true;
      }
    }
  }

  return false;
}

function ValidObjectId(oid: string): boolean {
  return /^[0-9a-fA-F]{24}$/.test(oid);
}

export {
  HasDuplicates,
  Equatable,
  Err,
  None,
  NonEmptyString,
  ObjectId,
  Ok,
  Option,
  Result,
  Some,
  ValidObjectId,
};
