package main

import (
	"api-skeleton/app/Console"
	"log"
)

func main() {
	err := Console.Execute()

	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
