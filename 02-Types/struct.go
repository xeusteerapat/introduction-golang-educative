package main

import (
	"fmt"
	"time"
)

type Bootcamp struct {
	// Latitude of the event
	Lat float64
	// Longitude of the event
	Lon float64
	// Date of the event
	Date time.Time
}

type Point struct {
	X, Y int
}

func main() {
	fmt.Println(Bootcamp{
		Lat:  34.012836,
		Lon:  -118.495338,
		Date: time.Now(),
	})

	// Accessing using dot notation
	event := Bootcamp{
		Lat: 34.012836,
		Lon: -118.495338,
	}

	fmt.Printf("Event on %s, location (%f, %f)\n",
		event.Date, event.Lat, event.Lon)

	// Struct literal
	var (
		p = Point{1, 2}  // has type Point
		q = &Point{1, 2} // has type *Point
		r = Point{X: 1}  // Y:0 is implicit
		s = Point{}      // X:0 and Y:0

	)

	fmt.Println(p, q, r, s)
}
