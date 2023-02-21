package main

import (
	"testing"
	"math"
)

func within(p, a, b float64) bool {
	return (1-p)*a < b && b < (1+p)*a
}

func TestToPiRange(t *testing.T) {
	cases := []struct{ in , want float64 } {
		{ math.Pi/2, math.Pi/2 },
		{ math.Pi/4, math.Pi/4 },
		{ 2*math.Pi, 0 },
		{ -5*math.Pi, math.Pi },
	}

	for _, c := range cases {
		if out := toPiRange(c.in); out != c.want {
			t.Errorf("to180range(%f) = %f, want %f\n", c.in, out, c.want)
		}
	}
}

func TestDegToRad(t *testing.T) {
	cases := []struct { in deg; want float64 } {
		{ deg(180), math.Pi },
		{ deg(360), 2*math.Pi },
		{ deg(720), 4*math.Pi },
		{ deg(90), math.Pi/2 },
		{ deg(-180), -math.Pi },
		{ deg(-360), -2*math.Pi },
		{ deg(-720), -4*math.Pi },
		{ deg(-90), -math.Pi/2 },
	};

	for _, c := range cases {
		if out := degToRad(c.in); out != c.want {
			t.Errorf("degToRad(%f) = %f, want %f\n", c.in, out, c.want);
		}
	}
}

func TestRadToDeg(t *testing.T) {
	cases := []struct { want deg; in float64 } {
		{ deg(180), math.Pi },
		{ deg(360), 2*math.Pi },
		{ deg(720), 4*math.Pi },
		{ deg(90), math.Pi/2 },
		{ deg(-180), -math.Pi },
		{ deg(-360), -2*math.Pi },
		{ deg(-720), -4*math.Pi },
		{ deg(-90), -math.Pi/2 },
	};

	for _, c := range cases {
		if out := radToDeg(c.in); out != c.want {
			t.Errorf("radToDeg(%f) = %f, want %f\n", c.in, out, c.want);
		}
	}
}
