package main

import "fmt"

func major(n1, n2 int) bool {
	return n1 > n2
}

var minor = func(n1, n2 int) bool {
	return n1 < n2
}

func lenName(name string) (string, int) {
	return name, len(name)
}

func main() {
	r := major(5, 7)
	fmt.Println(r)

	fmt.Println(minor(4, 6))

	fmt.Println(lenName("Eduardo"))

	name, lenght := lenName("Henrique")
	print(name, " ", lenght, "\n")

	_, lenght2 := lenName("Doe")
	fmt.Println(lenght2)
}
