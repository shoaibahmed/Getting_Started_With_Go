package main

import "fmt"

func main() {
	s := make([] string, 3)
	fmt.Println("Emp:", s)

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("Set:", s)
	fmt.Println("Get:", s[2])

	fmt.Println("Length:", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("Append:", s)

	c := make([] string, len(s))
	copy(c, s)
	fmt.Println("Copy:", c)

	l := s[2:5]
	fmt.Println("Slice # 01:", l)

	l = s[:5]
	fmt.Println("Slice # 02:", l)

	l = s[2:]
	fmt.Println("Slice # 03:", l)

	t := [] string {"g", "h", "i"}
	fmt.Println("Declaration:", t)

	twoD := make([][] int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([] int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2D array:", twoD)
}
