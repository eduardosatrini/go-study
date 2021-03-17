package main

import (
	"errors"
	"fmt"
)

func main() {
	// INTEGERS
	var n0 int
	var n1 int8
	var n2 int16
	var n3 int32
	var n4 int64

	var n5 uint8
	var n6 uint16
	var n7 uint32
	var n8 uint64

	var n9 rune // int84 = rune

	var n10 byte // uint8 = byte

	fmt.Println("Integers",
		n0, n1, n2, n3, n4,
		n5, n6, n7, n8, n9,
		n10,
	)

	// FLOATS
	var f1 float32
	var f2 float64

	fmt.Println("Floats", f1, f2)

	// STRINGS
	var s1 string = "ABC"
	var ch = 'G' // '' = char
	fmt.Println("Strings", s1)
	fmt.Println(ch)

	// BOOl
	var b1 bool
	var b2 = true
	print(b1, " ", b2, "\n")

	// ERROR
	var err1 error
	var err2 error = errors.New("type error")
	fmt.Println(err1)
	fmt.Println(err2)
}
