package main

import (
	"fmt"
	"image"
	"image/color"

	//	"image/jpeg"
	"image/gif"
	"os"
)

func northEastImage(points []northeast) *image.Paletted {
	const size = 500
	view := deg(20).Rad()

	img := image.NewPaletted(
		image.Rect(0, 0, 2*size+1, 2*size+1),
		[]color.Color{color.Black, color.White})

	for _, p := range points {
		x := p.east / view
		y := p.north / view
		img.SetColorIndex(size+int(x*size+0.5), size-int(y*size+0.5), 1)
	}
	return img
}

func main() {
	const (
		nframes     = 360
		degPerFrame = deg(1)
		delay       = 5
	)
	anim := gif.GIF{LoopCount: nframes}

	for f := 0; f < nframes; f++ {
		rot := float64(f) * degPerFrame.Rad()
		var points []northeast
		for i := range cities {
			xyz := cities[i].rotateEast(rot).XYZ(1.0) // take radius of Earth as unit
			if xyz.x < 0 {
				continue
			}
			xyz.x += 5
			_, ne := xyz.NorthEast()
			points = append(points, ne)
		}
		anim.Delay = append(anim.Delay, delay)
		anim.Image = append(anim.Image, northEastImage(points))
		fmt.Fprintf(os.Stderr, "\r[%.0f%%]", 100*float64(f)/nframes)
	}
	fmt.Fprintf(os.Stderr, "\r")

	gif.EncodeAll(os.Stdout, &anim)
	//jpeg.Encode(os.Stdout, northEastImage(points), nil)
}
