package main

import "fmt"

func main() {

	for i := 1; i <= 10; i++ {
		func(i int) {
			fmt.Println(i)
		}(i)
	}

	var myFunc func(int, string) bool
	fmt.Printf("%T ---- %v", myFunc, myFunc)

}
