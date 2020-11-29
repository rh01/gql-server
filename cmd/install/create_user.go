package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"report/internal/graphql/models"
	"github.com/joho/godotenv"
	"time"
)

// GetConfig by using godotenv
func GetConfig() {
	godotenv.Load()
}

// Run this file by using

// go run ./cmd/install.go

func main() {
	// get .env config
	GetConfig()

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	user := models.User{
		Name:     "admin",
		Username: os.Getenv("DEFAULT_USERNAME"),
		Password: os.Getenv("DEFAULT_PASSWORD"),
		Location: os.Getenv("DEFAULT_LOCATION"),
		Abbr:     os.Getenv("DEFAULT_ABBR"),
		Email:    os.Getenv("DEFAULT_EMAIL"),
		Openhab:  os.Getenv("DEFAULT_OPENHAB"),
	}

	user.HashPassword(user.Password)


	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		log.Panic()
	}
	fmt.Println("CREATED USER admin")
}
