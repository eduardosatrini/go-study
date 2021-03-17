package main

import "fmt"

type animal struct {
	age    int
	weight float64
	height float64
}

type dog struct {
	name string
	animal
}

type cat struct {
	name string
	animal
}

func main() {
	animal1 := animal{5, 4.6, 60}
	d := dog{"Rufus", animal1}
	fmt.Println("Struct dog:", d)

	animal2 := animal{3, 2.0, 40}
	c := cat{"Nita", animal2}
	fmt.Println("Struct cat:", c)

	fmt.Println(d.name, "and", c.name)
}
