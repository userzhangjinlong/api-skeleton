package main

import (
	"api-skeleton/app/Console"
	"api-skeleton/bootstrap"
	"log"
)

func main() {
	bootstrap.InitDB()
	err := Console.Execute()

	if err != nil {
		log.Fatalf("cmd.Execute err: %v", err)
	}
}
