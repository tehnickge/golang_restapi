package main 

import (
	"log"
	"github.com/tehnickge/golang_restapi/internal/app/apiserver"
)

func main() {
	s := apiserver.New()
	if err := s.Start(); err != nil {
		log.Fatal(err)
	}
}
