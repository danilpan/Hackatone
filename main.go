package main

import (
	"log"

	"github.com/madxiii/hackatone/app"
)

func main() {
	a, err := app.New()
	if err != nil {
		log.Fatalf("application err: %v\n", err)
	}

	errRun := a.Run()
	if errRun != nil {
		log.Printf("run: %v", errRun)
	}
}
