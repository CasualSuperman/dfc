package main

//import "github.com/CasualSuperman/dfc/terrain"

//type octree struct {
//	ranges   [3]int
//	subtrees [2][2][2]terrain.Chunk
//	base     Tile
//}

type Chunk interface {
}

type chunk struct {
	position struct {
		x, y, z int
	}
	size struct {
		x, y, z int
	}
	creatures []*Dwarf
}

func NewChunk() Chunk {
	return nil
}
