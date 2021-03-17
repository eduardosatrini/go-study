package main

import "fmt"

func main() {
	if n := 5; n%2 == 0 {
		fmt.Println("EVEN!")
	} else {
		fmt.Println("ODD!")
	}
}
