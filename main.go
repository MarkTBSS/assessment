package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//log.Fatal("Error loading .env file: ", godotenv.Load("development.env"))
	err := godotenv.Load("development.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	test := os.Getenv("ENVIRONMENT")
	fmt.Println(test)
}
