package main

import "fmt"

func main() {
	names := map[string]string{
		"firstName": "Eduardo",
		"lastName":  "Henrique",
	}

	fmt.Print(names["firstName"], " ")
	fmt.Print(names["lastName"], "\n")

	numbers := make(map[int]int)
	numbers[11] = 111
	fmt.Println(numbers)

	animals := map[bool]bool{}
	animals[true] = 5 > 5
	fmt.Println(animals)
}
