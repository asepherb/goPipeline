package main

import (
	"os"
	"fmt"
)

func getUserName() string {
	username := os.Getenv("USER")

	if username == "" {
		username = os.Getenv("USERNAME")
	}

	return username
}

func main()  {
	fmt.Printf("Hello, %s\n", getUserName())
}
