package main

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"image"
	"image/color"

	//	"image/jpeg"
	"image/gif"
	"io"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

const minPopulation = 0

//go:embed "worldcities.csv"
var cities_csv string

type city struct {
	city       string
	city_ascii string
	northeast
	country    string
	iso3       string
	admin_name string
	population int
}

var cities []city

func init() {
	r := csv.NewReader(strings.NewReader(cities_csv))
	header := make(map[int]string)
	rec, err := r.Read()
	if err != nil {
		log.Fatal(err)
	}
	for i := range rec {
		header[i] = rec[i]
	}

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		var c city
		for i, f := range record {
			switch header[i] {
			case "city":
				c.city = f
			case "city_ascii":
				c.city_ascii = f
			case "lat":
				d, _ := strconv.ParseFloat(f, 64)
				c.north = deg(d).Rad()
			case "lng":
				d, _ := strconv.ParseFloat(f, 64)
				c.east = deg(d).Rad()
			case "country":
				c.country = f
			case "iso3":
				c.iso3 = f
			case "admin_name":
				c.admin_name = f
			case "population":
				c.population, _ = strconv.Atoi(f)
			}
		}
		if c.population >= minPopulation {
			cities = append(cities, c)
		}
	}
}

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
		nframes = 720
		delay   = 5
	)
	anim := gif.GIF{LoopCount: nframes}

	for f := 0; f < nframes; f++ {
		rot := float64(f) / 360 * 2 * math.Pi
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
