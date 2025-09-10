package main

import (
	"log"
	"os"

	"github.com/jhonnydsl/payment-API/src/repository"
	"github.com/joho/godotenv"
)

func main() {
	if os.Getenv("DB_HOST") == "" {
		err := godotenv.Load()
		if err != nil {
			log.Println("no .env file found, continuing...")
		}
	}

	err := repository.Connect()
	if err != nil {
		log.Fatalf("error connecting to the database: %s", err)
	} else {
		log.Println("connection established")
	}
	defer repository.DB.Close()

	repo := &repository.TableRepository{}

	err = repo.CreateTablePayments()
	if err != nil {
		log.Fatalf("error creating table payments: %v", err)
	}

	err = repo.CreateTableUsers()
	if err != nil {
		log.Fatalf("error creating table users: %v", err)
	}

	err = repo.CreateTablePaymentEvents()
	if err != nil {
		log.Fatalf("error creating table payment_events: %v", err)
	}

	err = repo.CreateTablePaymentMethods()
	if err != nil {
		log.Fatalf("error creating table payment_methods")
	}
}