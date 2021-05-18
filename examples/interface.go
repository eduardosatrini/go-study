package main

import "fmt"

type Writer interface {
	Write([]byte) (int, error)
}

type ConWriter struct{}

func (w ConWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

func main() {

	var w Writer = ConWriter{}
	w.Write([]byte("AAA1"))

}
