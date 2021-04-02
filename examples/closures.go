package main

import "fmt"

func txt() func() {
	t := "function txt."

	inside := func() {
		fmt.Println(t)
	}

	return inside
}

func main() {
	t := "function main."
	fmt.Println(t)

	tx := txt()
	tx()
}
