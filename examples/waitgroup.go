package main

import (
	"fmt"
	"sync"
	"time"
)

func greeting(txt string) {
	for i := 0; i < 4; i++ {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}

func main() {
	var w sync.WaitGroup
	w.Add(3)

	go func() {
		greeting("Waitgroup!")
		w.Done()
	}()

	go func() {
		greeting("Other waitgroup!")
		w.Done()
	}()

	go func() {
		greeting("Other other waitgroup!")
		w.Done()
	}()

	w.Wait()
}
