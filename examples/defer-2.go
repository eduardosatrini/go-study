package main

import "fmt"

func studentStatus(h float32) bool {
	defer fmt.Println("Verify...")
	if h < 6 {
		return false
	}

	return true
}

func main() {
	s := studentStatus(4)
	fmt.Println(s)
}
