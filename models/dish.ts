import { NonEmptyString, Equatable } from "./common.ts";

class Dish implements Equatable {
  description: NonEmptyString;

  type: DishType;

  private constructor(description: NonEmptyString, type: DishType) {
    this.description = description;
    this.type = type;
  }

  equals(obj: any): boolean {
    return obj && obj.description === this.description;
  }
}

enum DishType {
  meat,
  fish,
  vegetarian,
}

export { Dish, DishType };
