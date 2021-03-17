package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Slices")

	mySliceInt := []int{50, 100, 150}
	mySliceString := []string{"G", "O"}

	// type with reflect lib
	fmt.Println(reflect.TypeOf(mySliceInt))
	fmt.Println(reflect.TypeOf(mySliceString))

	fmt.Println(mySliceInt, mySliceString)

	// append new values
	mySliceInt = append(mySliceInt, 175)
	mySliceString = append(mySliceString, "!")

	fmt.Println(mySliceInt, mySliceString)

	// slice a size
	mySliceInt2 := mySliceInt[0:2]
	mySliceString2 := mySliceString[1:]

	fmt.Println(mySliceInt2, mySliceString2)

	// change values
	mySliceInt2[0] = 123
	mySliceString2[1] = "G"
	fmt.Println(mySliceInt2, mySliceString2)
}
