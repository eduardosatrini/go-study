package main

import "fmt"

type animal struct {
	age    uint8
	weight float64
}

type Cat struct {
	animal
	nick  string
	noise string
}

func (c Cat) MakeNoise() {
	fmt.Println("Miau!")
}

func (c *Cat) Birth() {
	c.age++
}

func main() {
	a := animal{age: 5, weight: 4.3}
	c := Cat{a, "Katto", "Miau!"}

	fmt.Println(c.nick, c.age)
	c.MakeNoise()

	c.Birth()
	fmt.Println(c.age)
}
