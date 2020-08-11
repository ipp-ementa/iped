import { GeographicalLocation } from "./geographical_location.ts";

import {
  NonEmptyString,
  Equatable,
  Option,
  Some,
  None,
} from "./common.ts";

import { Menu } from "./menu.ts";

class Canteen implements Equatable {
  location: GeographicalLocation;

  name: string;

  menusMap: Map<Date, Menu[]>;

  private constructor(
    location: GeographicalLocation,
    name: NonEmptyString,
    menusMap: Map<Date, Menu[]>,
  ) {
    this.location = location;
    this.name = name.valueOf();
    this.menusMap = menusMap;
  }

  todayMenus(): Option<Menu[]> {
    const todayDate = new Date();
    todayDate.setHours(0, 0, 0, 0);

    const menus = this.menusMap.get(todayDate);

    if (menus) {
      return Some(menus);
    } else {
      return None;
    }
  }

  addMenu(menu: Menu): void {
    const todayDate = new Date();
    todayDate.setHours(0, 0, 0, 0);

    let menus = this.menusMap.get(todayDate);

    if (!menus) {
      menus = [];
    }

    menus.push(menu);

    this.menusMap.set(todayDate, menus);
  }

  equals(obj: any): boolean {
    return obj &&
      ((obj.name === this.name) ||
        this.location.equals(obj.location));
  }

  public static create(location: GeographicalLocation, name: NonEmptyString) {
    return new Canteen(location, name, new Map<Date, Menu[]>());
  }

  public static fromJson(object: any): Canteen {
    const localMenusMap = new Map<Date, Menu[]>();

    const objectMenusMap = Object.assign(
      new Map<Date, Menu[]>(),
      object.menusMap,
    );

    objectMenusMap.forEach((menus: Menu[], date: Date) =>
      localMenusMap.set(date, menus.map((m) => Menu.fromJson(m)))
    );

    return new Canteen(
      GeographicalLocation.fromJson(object.location),
      object.name,
      localMenusMap,
    );
  }
}

export { Canteen };
