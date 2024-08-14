package main

import (
	_ "embed"
	"fmt"
	"math"
)

const Tau = 2 * math.Pi

type northeast struct {
	north, east float64
}

type xyz struct {
	x, y, z float64
}

func (c xyz) String() string {
	return fmt.Sprintf("[%f %f %f]", c.x, c.y, c.z)
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

func (c northeast) East(r float64) northeast {
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

func (c northeast) String() string {
	return fmt.Sprintf("N%v E%v", radToDeg(c.north), radToDeg(c.east))
}

func toPiRange(a float64) float64 {
	a = math.Mod(a, Tau)
	if a <= -Tau/2 {
		a += Tau
	} else if a > Tau/2 {
		a -= Tau
	}
	return a
}
