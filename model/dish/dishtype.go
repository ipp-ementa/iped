package dish

// DishType is a enum representation of a dish type
// A UML overview of this enum can be found at https://github.com/ipp-ementa/iped-documentation/wiki/Architecture#models-structure
type DishType int

// Possible dish types are the following
const (
	Meat DishType = iota
	Fish
	Vegetarian
	Diet
)

// Validate check if a given integer is a valid dish type
func Validate(dishtype int) bool {
	return dishtype >= 0 && dishtype <= 3
}
