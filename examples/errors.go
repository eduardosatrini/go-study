package main

import (
	"errors"
	"fmt"
	"log"
)

func add(n1, n2 int) (int, error) {
	if n1 == n2 {
		return -1, errors.New("Numbers equals.")
	}

	return n1 + n2, nil
}

func main() {
	result, err := add(5, 5)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)
}
