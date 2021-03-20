package main

import "fmt"

func firstExec() {
	fmt.Println("First exec function!")
}

func secondExec() {
	fmt.Println("Second exec function!")
}

func main() {
	defer firstExec()
	secondExec()
}
