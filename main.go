package main

import (
	"log"
	"os"

	"github.com/jhonnydsl/payment-API/src/repository"
	"github.com/jhonnydsl/payment-API/src/utils/apperrors"
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

	apperrors.CheckErr(repo.CreateTablePayments(), "error creating table payments") 
	apperrors.CheckErr(repo.CreateTableUsers(), "error creating table users") 
	apperrors.CheckErr(repo.CreateTablePaymentEvents(), "error creating table payment_events") 
	apperrors.CheckErr(repo.CreateTablePaymentMethods(), "error creating table payment_methods") 
}