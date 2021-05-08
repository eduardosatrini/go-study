package main

import (
	"fmt"
)

func main() {

	city := map[string]int{
		"Belém":          20,
		"São Paulo":      30,
		"Rio de Janeiro": 15,
		"Belo Horizonte": 25,
	}

	city["Fortaleza"] = 11
	delete(city, "Rio de Janeiro")

	fmt.Println("\nSize map:", len(city))
	fmt.Println(city)

	pcity := city
	delete(pcity, "Belo Horizonte")
	fmt.Println(pcity)

	fmt.Println("\nExists: ")
	value, ok := city["Rio de Janeiro"]
	fmt.Println(value, ok) // 0, false

}
