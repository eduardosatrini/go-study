package main

import "fmt"

func div(n1, n2 float32) (float32, error) {

	if n1 == 0.0 || n2 == 0.0 {
		return 0, fmt.Errorf("Cannot divide by zero")
	}

	return n1 / n2, nil
}

func main() {

	result, err := div(4, 0)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Result:", result)
}
