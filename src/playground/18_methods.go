package main

import "fmt"

type rect struct {
	width, height, areaVar, pariVar int
}

// Methods defined for the struct rect
func (r *rect) area() int {
	a := r.width * r.height
	r.areaVar = a
	return a
}

func (r rect) perimeter() int {
	p := 2 * r.width + 2 * r.height
	r.pariVar = p
	return p
}

func main() {
	r := rect{width: 10, height: 5}

	fmt.Println("Area:", r.area())
	fmt.Println("Area variable:", r.areaVar)
	fmt.Println("Perimeter:", r.perimeter())
	fmt.Println("Perimeter variable:", r.pariVar)

	rp := &r
	fmt.Println("Area:", rp.area())
	fmt.Println("Area variable:", rp.areaVar)
	fmt.Println("Perimeter:", rp.perimeter())
	fmt.Println("Perimeter variable:", rp.pariVar)
}
