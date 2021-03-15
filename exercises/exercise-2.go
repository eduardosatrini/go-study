package main

import (
	"fmt"
)

func studentGrade(n1, n2 float64) {
	fmt.Println("First note:", n1)
	fmt.Println("Second note:", n2)
	fmt.Println("Media:", (n1+n2)/2)
}

func main() {
	studentGrade(8, 6.5)
}
