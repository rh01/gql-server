package main

import (
	"log"
	"os"
	"report/internal/graphql"
)

func main() {

	if err := graphql.Run(); err != nil {
		log.Fatalln(err)
		os.Exit(1)
	}
}
