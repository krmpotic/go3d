package main

import (
	_ "embed"
	"encoding/csv"
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"io"
	"log"
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
				c.north = degToRad(deg(d))
			case "lng":
				d, _ := strconv.ParseFloat(f, 64)
				c.east = degToRad(deg(d))
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

func main() {
	var points []northeast
	for i := range cities {
		xyz := cities[i].XYZ(1.0) // take radius of Earth as unit
		if xyz.x < 0 {
			continue
		}
		xyz.x += 5
		_, ne := xyz.NorthEast()
		points = append(points, ne)
	}

	const size = 500
	view := degToRad(deg(20))

	img := image.NewPaletted(
		image.Rect(0, 0, 2*size+1, 2*size+1),
		[]color.Color{color.White, color.Black})

	for _, p := range points {
		x := p.east / view
		y := p.north / view
		img.SetColorIndex(size + int(x*size+0.5), size - int(y*size+0.5), 1)
	}

	jpeg.Encode(os.Stdout, img, nil)
	fmt.Fprintf(os.Stderr, "done")
}
