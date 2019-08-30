package model

// DishType is a enum representation of a dish type
// A UML overview of this enum can be found at https://github.com/ipp-ementa/iped-documentation/wiki/Architecture
type DishType int

// Possible dish types are the following
const (
	Meat DishType = iota
	Fish
	Vegetarian
	Diet
)
