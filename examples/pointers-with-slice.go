package main

import "fmt"

func main() {

	a := []int{2, 4, 6} // slice and maps is used with reference
	b := a
	fmt.Println(a, b)

	a[0] = 10
	fmt.Println(a, b)

}
