import { Err, Ok, Option, Result } from "../deps.ts";

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

function HasDuplicates<T>(array: T[]) {
  return array.filter((e, i, a) => a.indexOf(e) !== i).length != 0;
}

export { HasDuplicates, Err, NonEmptyString, Ok, Option, Result };
