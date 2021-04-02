package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a letter: ")
	chr, _, err := r.ReadRune()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(chr)

	switch chr {
	case 'a':
		fmt.Println("Key pressed a")
	case 'A':
		fmt.Println("Key pressed A")
	default:
		fmt.Println("Something pressed.")
	}
}
