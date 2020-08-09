import { Err, Ok, Result } from "./common.ts";

class GeographicalLocation {
  latitude: number;

  longitude: number;

  private constructor(latitude: number, longitude: number) {
    this.latitude = latitude;

    this.longitude = longitude;
  }

  public static create(
    latitude: number,
    longitude: number,
  ): Result<GeographicalLocation, string> {
    if (latitude > 90 || latitude < -90) {
      return Err("latitude must range [-90, 90]");
    } else if (longitude > 90 || longitude < -90) {
      return Err("longitude must range [-180, 180]");
    } else {
      return Ok(new GeographicalLocation(latitude, longitude));
    }
  }
}

export { GeographicalLocation };
