package main

import (
	"fmt"
)

func main() {
	var i uint8
	for i <= 4 {
		fmt.Printf("%d ", i)
		i++
	}

	fmt.Println()

	for i := 0; i <= 4; i++ {
		fmt.Printf("%d ", i)
	}

	fmt.Println()

	n := [5]int{5, 10, 15, 20, 25}
	for i := range n {
		fmt.Printf("%d ", i)
	}
}
