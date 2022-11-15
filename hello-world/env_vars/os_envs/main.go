package main

import (
	"fmt"
	"os"
)

var myEnv map[string]string

/** Print a Hello World message using a .env file to set/read environment variables */
func main() {
	city := os.Getenv("CITY")

	fmt.Println("<<< Hello World from", city, ">>>")
}

// run: CITY=Lisbon go run main.go
// output: <<< Hello World from Lisbon >>>
