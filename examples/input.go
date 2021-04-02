package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
)

func main() {
	fmt.Println("Reading input")

	input := bufio.NewReader(os.Stdin)

	fmt.Print("Enter a name")
	txt, v := input.ReadString('\n')
	fmt.Println(reflect.TypeOf(txt), txt, v)
}
