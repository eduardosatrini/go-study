package main

import "fmt"

func option() func() {
	txt := "Inside function option!"

	f := func() {
		fmt.Println(txt)
	}

	return f
}

func main() {
	txt := "Inside function main!"
	fmt.Println(txt)

	newOption := option()
	newOption()
}
