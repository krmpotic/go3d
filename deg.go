package main

import (
	"math"
)

type deg float64

func radToDeg(r float64) deg {
	return deg(r / 2 / math.Pi * 360)
}

func (d deg) Rad() float64 {
	return float64(d) / 360 * 2 * math.Pi
}
