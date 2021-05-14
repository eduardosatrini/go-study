package main

import "fmt"

func main() {
	a := [3]int{5, 10, 15}
	b := &a[0]
	c := &a[1]

	fmt.Printf("%v %p %p\n", a, b, c)
}
