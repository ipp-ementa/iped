import {
  SchoolRepository,
} from "./school_controller.ts";

import {
  queryMenuById,
} from "./menu_controller.ts";

import {
  Error,
  NotFound,
  AvailableDishes,
  AvailableDishesItem,
  DetailedDishInformation,
} from "../views/views.ts";

import {
  Err,
  Ok,
  Result,
  Canteen,
  Menu,
  Dish,
  DishType,
  School,
} from "../models/models.ts";

export async function availableDishes(
  schoolRepository: SchoolRepository,
  schoolId: string,
  canteenId: string,
  menuId: string,
): Promise<Result<AvailableDishes, Error>> {
  const menuResult = await queryMenuById(
    schoolRepository,
    schoolId,
    canteenId,
    menuId,
  );

  if (menuResult.isErr()) {
    return Err(menuResult.unwrapErr());
  } else {
    const menu = menuResult.unwrap().menu;

    const dishes = menu.dishes;

    if (dishes.length == 0) {
      return Err(new NotFound());
    } else {
      const availableDishesView = dishes.map<AvailableDishesItem>(
        (function (d): AvailableDishesItem {
          return {
            id: d.description,
            description: d.description,
            type: DishType[d.type],
          };
        }),
      );

      return Ok(availableDishesView);
    }
  }
}

export async function detailedDishInformation(
  schoolRepository: SchoolRepository,
  schoolId: string,
  canteenId: string,
  menuId: string,
  dishId: string,
): Promise<Result<DetailedDishInformation, Error>> {
  const dishResult = await queryDishById(
    schoolRepository,
    schoolId,
    canteenId,
    menuId,
    dishId,
  );

  if (dishResult.isErr()) {
    return Err(dishResult.unwrapErr());
  } else {
    const dish = dishResult.unwrap().dish;

    const detailedDishInformation = <DetailedDishInformation> {
      id: dish.description,
      description: dish.description,
      type: DishType[dish.type],
    };

    return Ok(detailedDishInformation);
  }
}

export async function queryDishById(
  repository: SchoolRepository,
  schoolId: string,
  canteenId: string,
  menuId: string,
  dishId: string,
): Promise<
  Result<{ school: School; canteen: Canteen; menu: Menu; dish: Dish }, Error>
> {
  const menuResult = await queryMenuById(
    repository,
    schoolId,
    canteenId,
    menuId,
  );

  if (menuResult.isErr()) {
    return Err(menuResult.unwrapErr());
  } else {
    const menu = menuResult.unwrap().menu;

    const dish = menu.dishes.find((d) => d.description == dishId);

    if (dish) {
      return Ok(
        {
          school: menuResult.unwrap().school,
          canteen: menuResult.unwrap().canteen,
          menu: menuResult.unwrap().menu,
          dish: dish,
        },
      );
    } else {
      return Err(new NotFound());
    }
  }
}
