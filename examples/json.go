package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type user struct {
	Name   string  `json:"name"`
	Age    uint8   `json:"age"`
	Salary float64 `json:"salary"`
	Status bool    `json:"status"`
}

func main() {
	fmt.Println("Json")

	u := user{"Edu", 25, 5000, true}
	fmt.Println(u)

	userInJSON, err := json.Marshal(u)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("In bytes:", userInJSON)
	fmt.Println("Convert:", bytes.NewBuffer(userInJSON))

	u2 := map[string]string{
		"name": "Henri", "age": "26", "salary": "5500", "status": "false",
	}

	user2InJSON, err := json.Marshal(u2)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("In bytes:", user2InJSON)
	fmt.Println("Convert:", bytes.NewBuffer(user2InJSON))

}
