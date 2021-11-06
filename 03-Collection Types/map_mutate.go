package main

import "fmt"

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.6843, -74.3997},
	"Google":    {37.4220, -122.0841},
}

func main() {
	m["Splice"] = Vertex{34.0564, -74.3997} // mutate with new key-value
	fmt.Println(m["Splice"])

	delete(m, "Splice") // delete element
	fmt.Printf("%v\n", m)

	name, ok := m["Splice"] // check to see if element is present
	fmt.Printf("key 'Splice' is present?: %t - value: %v\n", ok, name)

	name, ok = m["Google"]
	fmt.Printf("key 'Google' is present?: %t - value: %v\n", ok, name)
}
