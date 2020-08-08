import { GeographicalLocation } from "./geographical_location.ts";
import { NonEmptyString } from "./common.ts";

class Canteen {
  location: GeographicalLocation;

  name: NonEmptyString;

  constructor(location: GeographicalLocation, name: string) {
    this.location = location;
    this.name = name;
  }
}

export { Canteen };
