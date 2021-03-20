package main

import "fmt"

type animal struct {
	color  string
	weight float64
	height float64
}

type cat struct {
	name string
	age  int
	animal
}

func main() {
	a := animal{"black", 2.0, 20.0}
	c := cat{"Katto", 3, a}

	fmt.Println(c)
	fmt.Println(c.name)

}
