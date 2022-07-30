package main

import (
	"log"

	"github.com/winebarrel/render"
)

func init() {
	log.SetFlags(0)
}

func main() {
	flags := parseFlags()

	for _, f := range flags.files {
		err := render.Render(f)

		if err != nil {
			log.Fatal(err)
		}
	}
}
