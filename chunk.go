package main

type octree struct {
	ranges   [3]int
	subtrees [2][2][2]Chunk
	base     Tile
}
