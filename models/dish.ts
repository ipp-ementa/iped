import { NonEmptyString, Equatable, Result, Ok, Err } from "./common.ts";

class Dish implements Equatable {
  description: string;

  type: DishType;

  private constructor(description: NonEmptyString, type: DishType) {
    this.description = description.valueOf();
    this.type = type;
  }

  equals(obj: any): boolean {
    return obj && obj.description === this.description;
  }

  public static create(description: NonEmptyString, type: DishType): Dish {
    return new Dish(description, type);
  }

  public static typeFromString(type: string): Result<DishType, string> {
    switch (type) {
      case DishType[0]:
        return Ok(DishType.meat);
      case DishType[1]:
        return Ok(DishType.diet);
      case DishType[2]:
        return Ok(DishType.fish);
      case DishType[3]:
        return Ok(DishType.vegetarian);
      default:
        return Err(`type: ${type} is not valid`);
    }
  }
}

enum DishType {
  meat,
  diet,
  fish,
  vegetarian,
}

export { Dish, DishType };
