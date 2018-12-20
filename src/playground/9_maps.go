package main

import "fmt"

func main() {
	m := make(map[string] int) // Map[Key type] Vaue type

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("Map:", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("Length:", len(m))

	delete(m, "k2")
	fmt.Println("Map:", m)

	val, present := m["k2"] // Second return value indicates if the value was found in the map
	fmt.Println("Found:", present, "| Val:", val)

	n := map[string] int {"foo": 1, "bar": 2}
	fmt.Println("Map:", n)
}
