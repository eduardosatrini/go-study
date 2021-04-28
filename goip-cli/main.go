package main

import (
	"goip-cli/app"
	"log"
	"os"
)

func main() {

	start := app.Generate()
	if err := start.Run(os.Args); err != nil {
		log.Fatal(err)
	}

}
