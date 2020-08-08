interface AvailableDishesItem {
  id: string;
  type: string;
  description: string;
}

export interface AvailableDishes extends Array<AvailableDishesItem> {}

export interface DetailedDishInformation {
  id: number;
  type: string;
  description: string;
}
