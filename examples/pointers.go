package main

import "fmt"

func main() {
	fmt.Println("Pointers")

	var value1 int = 25
	var value2 int = value1

	fmt.Println(value1, value2)

	value1 += 1
	fmt.Println(value1, value2)

	// pointer is a memory reference
	var num1 int
	var num2 *int
	fmt.Println(num1, num2)

	num1 = 100
	num2 = &num1
	fmt.Println(num1, num2)

	num1 = 120
	fmt.Println(num1, *num2)
}
