interface AvailableCanteensItem {
  id: number;
  name: string;
}

interface GeographicalLocation {
  latitude: number;
  longitude: number;
}

export interface AvailableCanteens extends Array<AvailableCanteensItem> {}

export interface CreateCanteen {
  name: string;
  location: GeographicalLocation;
}

export interface CreatedCanteen {
  id: string;
  name: string;
  location: GeographicalLocation;
}

export interface DetailedCanteenInformation {
  id: string;
  name: string;
  location: GeographicalLocation;
}
