package main

import (
	c "image/color"
)

type Material struct {
	Name    string
	Tile    int
	Temp    temps
	Density density
	Color   colors
	Value   int
}

type temps struct {
	/* Self-explanatory */
	Ignite, Melt, Boil, Fixed Temperature
	/* What temperature the material emits, like lava emitting heat. Rarely used. Ignored if 0. */
	Emit Temperature
	/* How readily the material changes temperature (how good it is at holding heat, for example.) */
	Rate float32
	/* Terperatures at which the material starts to take damage from heat and cold. */
	HeatDam, ColdDam Temperature
}

type density struct {
	/* Density when the material is in a liquid state. */
	Liquid int
	/* Density when the material is in a solid state. */
	Solid int
}

type colors struct {
	Display  c.RGBA
	Tile     c.RGBA
	Building c.RGBA
}

/* We operate on an absolute scale, like Kelvin */
type Temperature uint16
