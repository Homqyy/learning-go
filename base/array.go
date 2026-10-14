package main

import "fmt"

func main() {
	var b [3]int
	a := [5]int{1, 2, 3, 4, 5}

	b[0] = a[0]
	b[1] = a[1]

	for i, v := range b {
		fmt.Printf("[%v] %v\n", i, v)
	}

	// panic
	// b = a
}
