package main

import "github.com/CasualSuperman/dfc/terrain"

type world struct {
	chunks  [][]terrain.Chunk
	dwarves []Dwarf
}
