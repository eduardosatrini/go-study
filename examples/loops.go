package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Loops with for")

	// 1° form
	i := 1
	for i <= 3 {
		fmt.Println("count", i)
		time.Sleep(time.Second)
		i++
	}

	// 2º form
	for i := 1; i <= 3; i++ {
		fmt.Println("count", i)
		time.Sleep(time.Second)
	}

	// 3º form
	values := [3]int{5, 8, 6}
	for k, v := range values {
		fmt.Println("key", k, "value", v)
		time.Sleep(time.Second)
	}

	// 3.1º form
	names := [3]string{"Kin", "Kon", "Kat"}
	for _, v := range names {
		fmt.Println(v)
		time.Sleep(time.Second)
	}

	// 3.2º form
	for k, v := range "Eduardo" {
		fmt.Println(k, string(v))
		time.Sleep(time.Second)
	}

}
