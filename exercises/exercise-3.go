package main

import "fmt"

func tableCalc(n int) {
	for i := 1; i <= 10; i++ {
		fmt.Println(n, "x", i, "=", n*i)
	}
}

func main() {
	tableCalc(8)
}
