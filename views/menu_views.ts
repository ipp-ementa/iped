export interface AvailableMenusItem {
  id: string;
  type: string;
}

interface Dish {
  type: string;
  description: string;
}

export interface CreatedDish {
  id: string;
  type: string;
  description: string;
}

export interface AvailableMenus extends Array<AvailableMenusItem> {}

export interface CreateMenu {
  type: string;
  dishes: Dish[];
}

export interface CreatedMenu {
  id: string;
  type: string;
  dishes: CreatedDish[];
}

export interface DetailedMenuInformation {
  id: string;
  type: string;
  dishes: CreatedDish[];
}
