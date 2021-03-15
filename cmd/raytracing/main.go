package main

import (
	"log"
	"os"

	"github.com/y-taka-23/raytracing-go"
)

func main() {
	if err := raytracing.Run(os.Stdout, os.Stderr); err != nil {
		log.Fatal(err)
	}
}
