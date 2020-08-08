import { Err, Ok, Result } from "../deps.ts";

class NonEmptyString extends String {
  constructor(value: string) {
    super(value);
  }

  static create(value: string): Result<NonEmptyString, string> {
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

export { HasDuplicates, Err, NonEmptyString, Ok, Result };
