package main

import (
	"fmt"
	"time"
)

func main() {
	go greeting("my routine!")
	greeting("other routine!")
}

func greeting(txt string) {
	for {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}
