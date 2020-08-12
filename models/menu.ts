import { Dish } from "./dish.ts";
import { HasDuplicates, Ok, Result, Err } from "./common.ts";
import { v4 } from "../deps.ts";

class Menu {
  id: string;

  dishes: Dish[];

  type: MenuType;

  private constructor(dishes: Dish[], type: MenuType, id: string) {
    this.dishes = dishes;
    this.type = type;
    this.id = id;
  }

  public static create(dishes: Dish[], type: MenuType): Result<Menu, string> {
    if (dishes.length == 0) {
      return Err("must provide at least one dish");
    } else if (HasDuplicates<Dish>(dishes)) {
      return Err("cannot have duplicate dish");
    } else {
      return Ok(new Menu(dishes, type, `${MenuType[type]}_${v4.generate()}`));
    }
  }

  public static fromJson(object: any): Menu {
    return new Menu(
      object.dishes,
      object.type,
      object.id,
    );
  }

  public static typeFromString(type: string): Result<MenuType, string> {
    switch (type) {
      case MenuType[0]:
        return Ok(MenuType.lunch);
      case MenuType[1]:
        return Ok(MenuType.dinner);
      default:
        return Err(`type: ${type} is not valid`);
    }
  }
}

enum MenuType {
  lunch,
  dinner,
}

export { Menu, MenuType };
