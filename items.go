package main

type Item interface {
	ItemCategory
	Name() string                 // An artifact might have a special name
	Material() Material           // The material the item is made of
	Quality() Quality             // The quality of the item, if manufactured
	SetTemperature(t Temperature) // Set the item's temperature
	Temperature() Temperature     // This item's temperature
	Value() uint                  // The value of the item in monies
	Weight() uint                 // The weight, calculated from the type and the material
}

type ItemCategory interface {
	Character() rune // The character displayed in the interface
	Type() string    // The item type (sword, cup, etc.)
}
