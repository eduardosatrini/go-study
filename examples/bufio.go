package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	input := ConvertInt(Read("Enter a value: "))
	fmt.Printf("%T %d\n", input, input)

	input2 := ConvertFloat(Read("Enter a value: "))
	fmt.Printf("%T %f\n", input2, input2)

	input3 := Read("Enter a value: ")
	fmt.Printf("%T %s", input3, input3)
}

func Read(txt string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print(txt)
	scanner.Scan()
	input := scanner.Text()

	return input
}

func ConvertInt(n string) int64 {
	value, err := strconv.ParseInt(n, 10, 64)

	if err != nil {
		log.Fatal(err)
	}

	return value
}

func ConvertFloat(n string) float64 {
	value, err := strconv.ParseFloat(n, 64)

	if err != nil {
		log.Fatal(err)
	}

	return value
}
