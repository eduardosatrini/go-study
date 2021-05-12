package main

import (
	"fmt"
	"log"
)

func main() {

	fmt.Println("Start")

	defer func() {
		if err := recover(); err != nil {
			log.Println("Error:", err)
		}
	}()

	panic("something wrong")

}
