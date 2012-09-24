package main

type Tile struct {
	material_id int
	fluid       Fluid
	dwarves     []Dwarf
	animals     []interface{}
	items       []interface{}
}
