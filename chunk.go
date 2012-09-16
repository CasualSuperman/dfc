package main

const chunk_horiz_dim = 32
const chunk_vert_dim = 128

type Chunk struct {
	tiles [chunk_vert_dim][chunk_horiz_dim][chunk_horiz_dim]Tile
}
