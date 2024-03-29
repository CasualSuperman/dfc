package terrain

type Fluid uint8

func (f Fluid) Clean() bool {
	return (uint8(f) & 0x08) == 0x08
}

func (f Fluid) Type() uint8 {
	return (uint8(f) & 0xF0) >> 4
}

func (f Fluid) Depth() uint8 {
	return (uint8(f) & 0x07)
}
