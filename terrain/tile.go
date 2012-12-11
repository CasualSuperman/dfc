package terrain

import "github.com/CasualSuperman/dfc/items"

type Tile struct {
	material    *items.Material
	fluid       Fluid
	temperature items.Temperature
}
