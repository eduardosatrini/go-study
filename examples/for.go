package main

import "fmt"

func main() {

	// stmt := true
	// for stmt {
	// fmt.Print("OK! ")
	// stmt = false
	// }
	//
	// for i := 5; i >= 1; i-- {
	// fmt.Print(i)
	// }
	//
	// sl := []float64{2, 5, 10, 1.1, 0}
	// for i, v := range sl {
	// fmt.Print(i, " ", v)
	// }

	for x, y, z := 0, 6, 10; x <= 5; x, y, z = x+1, y-1, z*2 { // OMG! wtf?
		fmt.Println(x, y, z)
	}

}
