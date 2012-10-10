package main

type Fluid uint8


func (f Fluid) Clean() bool {
	return (uint8(f) & 0x80) == 0x80
}

func (f Fluid) Type() uint8 {
	return (uint8(f) & 0x70) >> 4
}

func (f Fluid) Depth() uint8 {
	return (uint8(f) & 0x0F)
}
