package main

type Dwarf struct {
	name      string
	hunger    int
	thirst    int
	awakeness int
	happiness int
	age       int
	relations []Relation
}

type Relation struct {
	relation_type int
	relative      *Dwarf
}

type Position_List []struct {
	x, y, z int
	thing *Dwarf
}
