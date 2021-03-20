package main

import "fmt"

func recoveryExec() {
	if r := recover(); r != nil {
		fmt.Println("Recovery program..")
	}
}

func mediaStudent(n1, n2 float32) bool {
	defer recoveryExec()
	m := (n1 + n2) / 2

	if m > 5 {
		return true
	} else if m < 5 {
		return false
	}

	panic("Media (5) error!")
}

func main() {
	fmt.Println("PANIC and RECOVER")
	fmt.Println(mediaStudent(6, 8))
}
