import { GeographicalLocation } from "./geographical_location.ts";
import { NonEmptyString, Equatable } from "./common.ts";

class Canteen implements Equatable {
  location: GeographicalLocation;

  name: string;

  private constructor(location: GeographicalLocation, name: NonEmptyString) {
    this.location = location;
    this.name = name.valueOf();
  }

  equals(obj: any): boolean {
    return obj &&
      ((obj.name === this.name) ||
        this.location.equals(obj.location));
  }

  public static create(location: GeographicalLocation, name: NonEmptyString) {
    return new Canteen(location, name);
  }

  public static fromJson(object: any): Canteen {
    return new Canteen(
      GeographicalLocation.fromJson(object.location),
      object.name,
    );
  }
}

export { Canteen };
