package main

import (
	"fmt"
)

func add(msg string, a ...int) {

	var total int
	for _, v := range a {
		total += v
	}

	fmt.Println("sum is: ", total)

}

func main() {
	add("add is: ", 3, 3, 2)
}
