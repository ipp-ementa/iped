import {
  SchoolRepository,
} from "./school_controller.ts";

import {
  queryCanteenById,
} from "./canteen_controller.ts";

import {
  Error,

  NotFound,

  BadRequest,
  AvailableMenus,
  AvailableMenusItem,
  CreateMenu,
  CreatedMenu,
  CreatedDish,
  DetailedMenuInformation,
} from "../views/views.ts";

import {
  Err,
  Ok,
  Result,
  Canteen,
  NonEmptyString,
  MenuType,
  Menu,
  Dish,
  DishType,
  School,
} from "../models/models.ts";

export async function availableMenus(
  schoolRepository: SchoolRepository,
  schoolId: string,
  canteenId: string,
): Promise<Result<AvailableMenus, Error>> {
  const canteenResult = await queryCanteenById(
    schoolRepository,
    schoolId,
    canteenId,
  );

  if (canteenResult.isErr()) {
    return Err(canteenResult.unwrapErr());
  } else {
    const canteen = canteenResult.unwrap().canteen;

    const menusOption = canteen.todayMenus();

    if (menusOption.isNone()) {
      return Err(new NotFound());
    } else {
      const menus = menusOption.unwrap();

      const availableMenusView = menus.map<AvailableMenusItem>(
        (function (m): AvailableMenusItem {
          return {
            id: m.id,
            type: MenuType[m.type],
          };
        }),
      );

      return Ok(availableMenusView);
    }
  }
}

export async function createMenu(
  schoolRepository: SchoolRepository,
  schoolId: string,
  canteenId: string,
  menuToCreate: CreateMenu,
): Promise<Result<CreatedMenu, Error>> {
  const canteenResult = await queryCanteenById(
    schoolRepository,
    schoolId,
    canteenId,
  );

  if (canteenResult.isErr()) {
    return Err(canteenResult.unwrapErr());
  } else {
    const school = canteenResult.unwrap().school;

    const canteen = canteenResult.unwrap().canteen;

    const menuTypeResult = Menu.typeFromString(menuToCreate.type);

    const dishesResults = menuToCreate.dishes.map(
      function (d): Result<Dish, Error> {
        const dishTypeResult = Dish.typeFromString(d.type);

        const descriptionResult = NonEmptyString.create(d.description);

        if (dishTypeResult.isErr()) {
          return Err(new BadRequest(dishTypeResult.unwrapErr()));
        } else if (descriptionResult.isErr()) {
          return Err(new BadRequest(descriptionResult.unwrapErr()));
        } else {
          return Ok(
            Dish.create(descriptionResult.unwrap(), dishTypeResult.unwrap()),
          );
        }
      },
    );

    const results = [menuTypeResult, ...dishesResults];

    const firstError = results.find((r) => r.isErr());

    if (firstError) {
      return Err(firstError.unwrapErr());
    } else {
      const menuResult = Menu.create(
        dishesResults.map((dr) => dr.unwrap()),
        menuTypeResult.unwrap(),
      );

      if (menuResult.isErr()) {
        return Err(new BadRequest(menuResult.unwrapErr()));
      } else {
        const menu = menuResult.unwrap();

        canteen.addMenu(menu);

        const localCanteenUpdateResult = school.updateCanteen(canteen);

        if (localCanteenUpdateResult.isErr()) {
          return Err(new BadRequest(localCanteenUpdateResult.unwrapErr()));
        } else {
          const updateResult = await schoolRepository.update(school);

          if (updateResult.isErr()) {
            return Err(updateResult.unwrapErr());
          } else {
            const createdMenuView = <CreatedMenu> {
              id: menu.id,
              type: MenuType[menu.type],
              dishes: menu.dishes.map(function (d): CreatedDish {
                return {
                  id: d.description,
                  description: d.description,
                  type: DishType[d.type],
                };
              }),
            };
            return Ok(createdMenuView);
          }
        }
      }
    }
  }
}

export async function detailedMenuInformation(
  schoolRepository: SchoolRepository,
  schoolId: string,
  canteenId: string,
  menuId: string,
): Promise<Result<DetailedMenuInformation, Error>> {
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

    const detailedMenuInformation = <DetailedMenuInformation> {
      id: menu.id,
      type: MenuType[menu.type],
      dishes: menu.dishes.map(function (d): CreatedDish {
        return {
          id: d.description,
          description: d.description,
          type: DishType[d.type],
        };
      }),
    };

    return Ok(detailedMenuInformation);
  }
}

export async function queryMenuById(
  repository: SchoolRepository,
  schoolId: string,
  canteenId: string,
  menuId: string,
): Promise<Result<{ school: School; canteen: Canteen; menu: Menu }, Error>> {
  const canteenResult = await queryCanteenById(
    repository,
    schoolId,
    canteenId,
  );

  if (canteenResult.isErr()) {
    return Err(canteenResult.unwrapErr());
  } else {
    const school = canteenResult.unwrap().school;

    const canteen = canteenResult.unwrap().canteen;

    const menu = Array.from(canteen.menusMap.values()).flatMap((m) => m).find((
      m,
    ) => m.id == menuId);

    if (menu) {
      return Ok({ school: school, canteen: canteen, menu: menu });
    } else {
      return Err(new NotFound());
    }
  }
}
