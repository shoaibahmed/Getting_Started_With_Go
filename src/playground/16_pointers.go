package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Initial:", i)

	zeroval(i)
	fmt.Println("Zero val:", i)

	zeroptr(&i)
	fmt.Println("Zero ptr:", i)

	fmt.Println("Pointer:", &i)
}
