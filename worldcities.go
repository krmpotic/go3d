package main

import (
	_ "embed"
	"encoding/csv"
	"io"
	"log"
	"strconv"
	"strings"
)

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
		cities = append(cities, c)
	}
}
