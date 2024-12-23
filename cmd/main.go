package main

import (
	"log"
	"net/http"
	"os"

	"github.com/ashtotakoe/calculator-web-service/internal/server"
)

func main() {
	port := os.Args[1]

	runWithDetailedValidation := false

	if len(os.Args) > 2 {
		runWithDetailedValidation = os.Args[2] == "detailed-validation"
	}

	s := server.NewServer(runWithDetailedValidation)

	log.Printf("Server is running on port %s \n detailed validation = %t\n", port, runWithDetailedValidation)
	err := http.ListenAndServe(":"+port, s)
	if err != nil {
		log.Println(err)
	}
}
