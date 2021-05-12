package main

import "fmt"

func main() {
	fmt.Println("Start!")
	defer fmt.Println("Start! Defer!")
	panic("Something is wrong.")
	fmt.Println("End.")
}
