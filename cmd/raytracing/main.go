package main

import (
	"log"

	"github.com/y-taka-23/raytracing-go"
)

func main() {
	if err := raytracing.Run(); err != nil {
		log.Fatal(err)
	}
}
