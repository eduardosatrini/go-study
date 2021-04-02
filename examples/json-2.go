package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type cat struct {
	Name   string  `json:"name"`
	Age    uint8   `json:"age"`
	Weight float64 `json:"weight"`
}

func main() {
	catInJSON := `{"name":"Katto", "age": 3, "weight": 4.5}`

	var c cat
	if err := json.Unmarshal([]byte(catInJSON), &c); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c)

	cat2InJSON := `{"name":"Kitte", "age": "5", "weight": "6.5"}` // ???

	c2 := make(map[string]string)
	if err := json.Unmarshal([]byte(cat2InJSON), &c2); err != nil {
		log.Fatal(err)
	}
	fmt.Println(c2)

}
