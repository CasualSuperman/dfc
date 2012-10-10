package main

const chunk_horiz_dim = 32
const chunk_vert_dim = 128

/*
type Chunk struct {
	tiles [chunk_vert_dim][chunk_horiz_dim][chunk_horiz_dim]Tile
}
*/

type octree struct {
	ranges   [3]int
	subtrees [2][2][2]Chunk
	base     Tile
}

type Chunk interface {
	At(x, y, z int) Tile
	From(x, y, z int) interface {
		To(x, y, z int) Chunk
	}
}
