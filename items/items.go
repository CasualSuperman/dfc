package items

type Item interface {
	ItemCategory
	Name() string                 // An artifact might have a special name
	Material() Material           // The material the item is made of
	Quality() Quality             // The quality of the item, if manufactured
	SetTemperature(t Temperature) // Set the item's temperature
	Temperature() Temperature     // This item's temperature
	Value() uint                  // The value of the item in monies
	Weight() float32              // The weight, calculated from the type and the material
}

type ItemCategory interface {
	Character() rune // The character displayed in the interface
	Type() string    // The item type (sword, cup, etc.)
}

type Quality uint8

func NewItemCategory(r rune, name string, volume float32, value uint) ItemCategory {
	return itemCategory{r, name, volume, value}
}

type itemCategory struct {
	character rune
	name      string
	volume    float32
	value     uint
}

func (i itemCategory) Character() rune {
	return i.character
}

func (i itemCategory) Type() string {
	return i.name
}

type item struct {
	*itemCategory
	name        string
	material    *Material
	quality     Quality
	temperature Temperature
}

func (i item) Value() uint {
	return i.material.Value * i.itemCategory.value
}

func (i item) Weight() float32 {
	return i.itemCategory.volume * float32(i.material.Density.Solid)
}
