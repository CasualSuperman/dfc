package main

type Dwarf struct {
	hunger    int
	thirst    int
	awakeness int
	happiness int
	relations []Relation
}

type Relation struct {
	relation_type int
	other         *Dwarf
}
