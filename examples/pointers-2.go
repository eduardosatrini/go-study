package main

import (
	"fmt"
)

func main() {
	var a int = 5
	var b *int = &a
	fmt.Println(&a, b)

	a = 6
	fmt.Println(a, *b) // 6 6

	*b = 1
	fmt.Println(a, *b) // 1 1
}
