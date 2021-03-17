package main

import "fmt"

type user struct {
	name   string
	age    int
	salary float64
}

func main() {
	fmt.Println("Struct")

	// 1º form
	var u user
	u.name = "Eduardo"
	u.age = 25
	u.salary = 3000
	fmt.Println(u)

	// 2º form
	u2 := user{"Henrique", 24, 3400}
	fmt.Println(u2)

	// 3º form
	u3 := user{name: "Marie"}
	fmt.Println(u3)
}
