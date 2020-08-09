import { NonEmptyString } from "./common.ts";

class Dish {
  description: NonEmptyString;

  type: DishType;

  private constructor(description: NonEmptyString, type: DishType) {
    this.description = description;
    this.type = type;
  }
}

enum DishType {
  meat,
  fish,
  vegetarian,
}

export { Dish, DishType };
