export interface AvailableSchoolsItem {
  id: string;
  acronym: string;
  name: string;
}

interface GeographicalLocation {
  latitude: number;
  longitude: number;
}

interface Canteen {
  name: string;
  location: GeographicalLocation;
}

interface CreatedCanteen {
  id: string;
  name: string;
}

export interface AvailableSchools extends Array<AvailableSchoolsItem> {}

export interface CreateSchool {
  acronym: string;
  name: string;
  canteens: Canteen[];
}

export interface CreatedSchool {
  id: string;
  acronym: string;
  name: string;
  canteens: CreatedCanteen[];
}

export interface DetailedSchoolInformation {
  id: string;
  acronym: string;
  name: string;
  canteens: CreatedCanteen[];
}
