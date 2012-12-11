package terrain

import "sync"
import "github.com/larspensjo/Go-simplex-noise/simplexnoise"

const chunk_horiz_dim = 32
const chunk_vert_dim = 128

type chunk struct {
	tiles     [chunk_horiz_dim][chunk_horiz_dim][chunk_vert_dim]Tile
	tickCount uint8
}

type Chunk interface {
	Tick(*sync.WaitGroup)
}

type Tile interface{}

var _ = simplexnoise.Noise1(0)
