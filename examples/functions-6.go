package main

import (
	"fmt"
)

func hello(f, l *string) {
	fmt.Println(*f, *l)
	*f = "Edu"
	*l = "Henry"
}

func main() {
	firstName := "Eduardo"
	lastName := "Henrique"
	hello(&firstName, &lastName)

	fmt.Println(firstName, lastName)
}
