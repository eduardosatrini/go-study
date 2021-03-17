package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"firstName": "Eduardo",
		"lastName":  "Henrique",
		"status":    ":|",
	}

	fmt.Println(user)
	fmt.Println(user["firstName"], user["lastName"])

	print("\n")

	market := map[string]map[string]float64{
		"fruits": {
			"Banana":     5,
			"Apple":      6,
			"Watermelon": 2,
		},
		"vegetables": {
			"Briccoli":   2,
			"Carrots":    5,
			"Green Peas": 6,
		},
	}

	print("\n")

	fmt.Println(market)

	delete(market, "vegetables")
	fmt.Println(market)

	market["employee"] = map[string]float64{
		"Eduardo":  5500,
		"Henrique": 6700,
	}

	fmt.Println(market["employee"])
}
