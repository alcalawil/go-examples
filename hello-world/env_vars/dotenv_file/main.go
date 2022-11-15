package main

import (
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/joho/godotenv/autoload"
)

var myEnv map[string]string

/** Print a Hello World message using a .env file to set/read environment variables */
func main() {
	myEnv, err := godotenv.Read()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	city := myEnv["CITY"]
	fmt.Println("<<< Hello World from", city, ">>>")
}
