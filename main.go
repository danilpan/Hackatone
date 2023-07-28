package main

import (
	"fmt"
	"log"

	"github.com/madxiii/hackatone/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatalf("application err: %v\n", err)
	}
	fmt.Println("Cfgg", a)
}
