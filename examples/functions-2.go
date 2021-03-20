package main

import "fmt"

func doubleAndTriple(n int) (double, triple int) {
	double = n * 2
	triple = n * 3
	return // return double and tiple variable
}

func main() {
	d, t := doubleAndTriple(7)
	fmt.Println(d, t)

}
