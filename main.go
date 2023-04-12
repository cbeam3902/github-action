package main

import (
	"log"

	"github.com/cbeam3902/github-action/microservice"
)

func main() {
	s := microservice.NewServer("", "8000")
	log.Fatal(s.ListenAndServe())
}
