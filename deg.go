package main

import (
	"fmt"
)

type deg float64

func (d deg) String() string {
	return fmt.Sprintf("%+ 6.4f°", float64(d))
}

func radToDeg(r float64) deg {
	return deg(r / Tau * 360)
}

func (d deg) Rad() float64 {
	return float64(d) / 360 * Tau
}
