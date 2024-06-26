package main

import (
	"log"

	"github.com/Salikhov079/rent_car/storage/postgres"
)

func main() {
	_, err := postgres.NewPostgresStorage()
	if err != nil {
		log.Fatal("Error while connection on db: ", err.Error())
	}
}
