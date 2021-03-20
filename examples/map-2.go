package main

import "fmt"

func main() {
	fmt.Println("Maps")

	perfil := map[string]string{
		"name":       "Eduardo",
		"profession": "Programmer",
		"status":     ":D",
	}
	fmt.Println(perfil)
	fmt.Println("Profession:", perfil["profession"])

	prices := map[string]float32{
		"playstation_5":   4000,
		"xbox_x":          4200,
		"nintendo_switch": 3200,
	}
	delete(prices, "xbox_x")
	fmt.Println(prices)
}
