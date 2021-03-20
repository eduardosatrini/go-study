package main

import "fmt"

func sumTotalNumbers(n ...int) int {
	var total int
	for _, v := range n {
		total += v
	}

	return total
}

func joinTextNumber(txt string, n ...int) {
	for _, v := range n {
		fmt.Println(txt, v)
	}
}

func main() {
	t := sumTotalNumbers(6, 2, 7, 3)
	fmt.Println(t)

	joinTextNumber("Sequence:", 2, 5, 3, 1, 4)
}
