package main

import "github.com/CasualSuperman/dfc/terrain"

type octree struct {
	ranges   [3]int
	subtrees [2][2][2]terrain.Chunk
	base     Tile
}
