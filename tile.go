package main

type Tile struct {
	material_id int
	fluid       uint8
}

func (t Tile) getFluidType() uint8 {
	return (t.fluid & 0xF0) >> 4
}

func (t Tile) getFluidDepth() uint8 {
	return (t.fluid & 0x0F)
}
