import { Dish } from "./dish.ts";
import { HasDuplicates, Ok, Result, Err } from "./common.ts";

class Menu {
  dishes: Dish[];

  type: MenuType;

  private constructor(dishes: Dish[], type: MenuType) {
    this.dishes = dishes;
    this.type = type;
  }

  public static create(dishes: Dish[], type: MenuType): Result<Menu, string> {
    if (dishes.length == 0) {
      return Err("must provide at least one dish");
    } else if (HasDuplicates<Dish>(dishes)) {
      return Err("cannot have duplicate dish");
    } else {
      return Ok(new Menu(dishes, type));
    }
  }
}

enum MenuType {
  lunch,
  dinner,
}

export { Menu, MenuType };
