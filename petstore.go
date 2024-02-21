package main

import (
	"flag"
	"log"

	"github.com/jvansan/example-opapi-codegen/server"
	"github.com/jvansan/example-opapi-codegen/server/api"
)

func main() {
	port := flag.String("port", "8080", "Port for test HTTP server")
	flag.Parse()
	// Create an instance of our handler which satisfies the generated interface
	petStore := api.NewPetStore()
	s := server.NewGinPetServer(petStore, *port)
	// And we serve HTTP until the world ends.
	log.Fatal(s.ListenAndServe())
}
