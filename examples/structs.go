package main

import "fmt"

type Movie struct {
	name    string
	year    uint
	reviews []float64
}

func main() {
	heroes := Movie{
		name:    "Spider Man 1",
		year:    2003,
		reviews: []float64{8.8, 9.0, 7.5},
	}

	fmt.Println(heroes.name, heroes.year)
	fmt.Println(heroes.reviews)
}
