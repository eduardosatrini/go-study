package main

import (
	"bufio"
	"fmt"
	"os"
)

func scanner() {
	s := bufio.NewScanner(os.Stdin)
	fmt.Print("input: ")
	for s.Scan() {
		fmt.Print("output: ")
		fmt.Println(s.Text())
	}
}

func main() {
	scanner()
}
