package main

import "fmt"

func main() {

	func() {
		fmt.Println("Closures")
	}()

	r := func(txt string) string {
		return fmt.Sprintf("Get -> %s", txt)
	}("Hi!")

	fmt.Println(r)

	n := func(n ...int) int {
		var t int
		for _, v := range n {
			t += v
		}

		return t
	}(4, 2, 3, 1)

	fmt.Println(n)
}
