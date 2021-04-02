package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

func main() {
	fmt.Println("Number favority:", rand.Intn(10))

	var a, b, c = "Python", "Go", "PHP"
	d := "C++"
	fmt.Println(a, b, c, d)
	fmt.Println(reflect.TypeOf(d))

	const fistName, lastName = "Eduardo", "Henrique"
	fmt.Println(reflect.TypeOf(fistName))

	for i := 10; i >= 0; i -= 2 {
		time.Sleep(time.Second)
		fmt.Printf(" %d ", i)

		for j := 0; j <= len(b)-1; j++ {
			fmt.Printf("%s", string(b[j]))
		}
	}

}
