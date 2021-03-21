package main

import "fmt"

type user struct {
	name   string
	age    uint8
	status bool
}

func (u user) greeting() {
	fmt.Printf("You're welcome %s!\n", u.name)
}

func (u user) major() bool {
	return u.age >= 18
}

func (u *user) birth() {
	u.age++
}

func main() {
	u := user{"Eduardo", 17, true}
	u.greeting()
	m := u.major()
	fmt.Println(m)
	u.birth()
	fmt.Println(u)

	print("\n")

	u2 := user{"Henrique", 26, false}
	u2.greeting()
	m2 := u2.major()
	fmt.Println(m2)
	fmt.Println(u2)
}
