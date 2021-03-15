package main

import (
	"fmt"
	"math"
)

func doubleTripleSqrt(n float64) {
	fmt.Println("Double:", n*2)
	fmt.Println("Triple:", n*3)
	fmt.Println("Square root:", math.Sqrt(n))
}

func main() {
	doubleTripleSqrt(85)
}
