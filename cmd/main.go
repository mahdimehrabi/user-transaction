package main

import (
	"bbdk/app/api"
	"bbdk/cmd/seeder"
	"flag"
)

func main() {
	seed := flag.Bool("seed", false, "Seed the database with initial data")

	flag.Parse()

	if *seed {
		seeder.Seed()
		return
	}

	api.Boot()
}
