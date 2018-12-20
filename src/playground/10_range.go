package main

import "fmt"

func main() {
	nums := [] int {2, 3, 4}
	
	// Range iterates over the elements
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("Sum:", sum)

	kvs := map[string] string {"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s\n", k ,v)
	}

	for k := range kvs {
		fmt.Println("Key:", k)
	}

	for i, c := range "go" {
		fmt.Println(i, c)
	}
}
