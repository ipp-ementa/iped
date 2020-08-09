import { GeographicalLocation } from "./geographical_location.ts";
import { NonEmptyString } from "./common.ts";

class Canteen {
  location: GeographicalLocation;

  name: NonEmptyString;

  private constructor(location: GeographicalLocation, name: NonEmptyString) {
    this.location = location;
    this.name = name;
  }

  public static create(location: GeographicalLocation, name: NonEmptyString) {
    return new Canteen(location, name);
  }
}

export { Canteen };
