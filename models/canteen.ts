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

  menusMap: Map<string, Menu[]>;

  private constructor(
    location: GeographicalLocation,
    name: NonEmptyString,
    menusMap: Map<string, Menu[]>,
  ) {
    this.location = location;
    this.name = name.valueOf();
    this.menusMap = menusMap;
  }

  todayMenus(): Option<Menu[]> {
    const menusObject = Object.fromEntries(this.menusMap);

    const todayDate = new Date();
    todayDate.setHours(0, 0, 0, 0);

    const menus = menusObject[todayDate.toString()];

    if (menus) {
      return Some(menus);
    } else {
      return None;
    }
  }

  addMenu(menu: Menu): void {
    const todayDate = new Date();
    todayDate.setHours(0, 0, 0, 0);

    let menus = this.menusMap.get(todayDate.toString());

    if (!menus) {
      menus = [];
    }

    menus.push(menu);

    this.menusMap.set(todayDate.toString(), menus);
  }

  equals(obj: any): boolean {
    return obj &&
      ((obj.name === this.name) ||
        this.location.equals(obj.location));
  }

  public static create(location: GeographicalLocation, name: NonEmptyString) {
    return new Canteen(location, name, new Map<string, Menu[]>());
  }

  public static fromJson(object: any): Canteen {
    return new Canteen(
      GeographicalLocation.fromJson(object.location),
      object.name,
      new Map<string, Menu[]>(Object.entries(object.menusMap)),
    );
  }
}

export { Canteen };
