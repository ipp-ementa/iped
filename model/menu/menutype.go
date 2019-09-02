package menu

// MenuType is a enum representation of a menu type
// A UML overview of this enum can be found at https://github.com/ipp-ementa/iped-documentation/wiki/Architecture#models-structure
type MenuType int

// A menu is either available at lunch or dinner
const (
	Lunch MenuType = iota
	Dinner
)

// String is the implementation of stringer for MenuType enum
// If the MenuType being converted is not valid the string "nil" will be returned
func (Type MenuType) String() string {

	switch Type {
	case Lunch:
		return "lunch"
	case Dinner:
		return "dinner"
	default:
		return "nil"
	}
}

// Validate check if a given integer is a valid menu type
func Validate(menutype int) bool {
	return menutype == 0 || menutype == 1
}
