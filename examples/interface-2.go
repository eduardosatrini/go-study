package main

import "fmt"

type Writer interface {
	Write([]byte) (int, error)
}

type Printer struct{}

func (pr Printer) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func main() {
	var p Writer = Printer{}
	p.Write([]byte("Hello, Go!"))
}
