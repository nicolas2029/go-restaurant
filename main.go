package main

import (
	"log"

	"github.com/nicolas2029/go-restaurant/pkg/address"
	"github.com/nicolas2029/go-restaurant/storage"
)

func main() {
	storage.NewPostgresDB()
	storageAddress := storage.NewPsqlAddress(storage.Pool())
	serviceAddress := address.NewService(storageAddress)
	ms, err := serviceAddress.GetAll()

	if err != nil {
		log.Fatalf("Fatal en storage.GetAll(): %v", err)
	}

	log.Printf("%s", ms)
}
