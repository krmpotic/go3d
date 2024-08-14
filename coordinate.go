package main

import (
	_ "embed"
	"math"
)

type northeast struct {
	north, east float64
}

type xyz struct {
	x, y, z float64
}

func (c northeast) XYZ(r float64) xyz {
	north, east := c.north, c.east
	if north > math.Pi/2 || north < -math.Pi/2 {
		panic("north (latitude) out of [-Pi,+Pi] bounds")

	}
	z := r * math.Sin(north)
	r_xy := r * math.Cos(math.Abs(north))
	y := r_xy * math.Sin(east)
	x := r_xy * math.Cos(east)
	return xyz{x: x, y: y, z: z}
}

func (c northeast) rotateEast(r float64) northeast {
	return northeast{north: c.north, east: toPiRange(c.east + r)}
}

func (c xyz) NorthEast() (r float64, ne northeast) {
	x, y, z := c.x, c.y, c.z
	r = math.Sqrt(x*x + y*y + z*z)
	north := math.Asin(z / r)
	r_xy := r * math.Cos(math.Abs(north))
	east := math.Asin(y / r_xy)
	if x < 0 {
		east = math.Pi - east
	}
	east = toPiRange(east)

	ne.north, ne.east = north, east
	return
}

func toPiRange(a float64) float64 {
	a = math.Mod(a, 2*math.Pi)
	if a <= -math.Pi {
		a += 2 * math.Pi
	} else if a > math.Pi {
		a -= 2 * math.Pi
	}
	return a
}
