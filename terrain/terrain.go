package terrain

import "sync"
import "github.com/larspensjo/Go-simplex-noise/simplexnoise"
import "github.com/CasualSuperman/dfc/items"

const chunk_horiz_dim = 64
const chunk_vert_dim = 128

type chunk struct {
	// Temperature is at a slightly lower res than discrete tiles are.
	temperature [chunk_horiz_dim / 4][chunk_horiz_dim / 4][chunk_vert_dim / 4]items.Temperature

	fluids    [chunk_horiz_dim][chunk_horiz_dim][chunk_vert_dim]Fluid
	tickCount uint8
}

type Chunk interface {
	Tick(*sync.WaitGroup)
}

var _ = simplexnoise.Noise1(0)
