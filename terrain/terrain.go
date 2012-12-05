package terrain

import "sync"
import "github.com/larspensjo/Go-simplex-noise/simplexnoise"

const chunk_horiz_dim = 32
const chunk_vert_dim = 128

type chunk [chunk_horiz_dim][chunk_horiz_dim][chunk_vert_dim]Tile

type Chunk interface {
	Interacter
	At(x, y, z int) Tile
	From(x, y, z int) interface {
		To(x, y, z int) Interacter
	}
	Tick(*sync.WaitGroup)
}

type Interacter interface {
}

type Tile interface{}

var _ = simplexnoise.Noise1(0)
