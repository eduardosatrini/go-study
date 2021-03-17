package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Array and Slices")

	var values [3]int // array
	fmt.Println(values)

	values[0] = 10
	values[1] = 20
	values[2] = 30
	fmt.Println(reflect.TypeOf(values))
	fmt.Println(values)
	fmt.Println("================")

	values2 := [3]string{"A", "B", "C"}
	fmt.Println(reflect.TypeOf(values2))
	fmt.Println(values2)
	fmt.Println("================")

	values3 := [...]float64{4, 8, 2}
	fmt.Println(reflect.TypeOf(values3))
	fmt.Println(values3)
}
