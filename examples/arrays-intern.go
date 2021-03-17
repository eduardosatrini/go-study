package main

import "fmt"

func main() {
	// make(type, start, end)
	slice := make([]uint, 3, 4)
	fmt.Println(slice)
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))

	print("\n")

	slice = append(slice, 8)
	slice = append(slice, 1)
	fmt.Println(slice)
	fmt.Println("Length:", len(slice))
	fmt.Println("Capacity:", cap(slice))
}
