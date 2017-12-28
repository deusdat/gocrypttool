package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	if len(os.Args) != 3 {
		log.Fatal("Please provide the password and the cost.")
	}

	password := os.Args[1]

	cost, err := strconv.Atoi(os.Args[2])
	if err != nil {
		log.Fatal("Please provide a cost, as a number as second argument")
	}
	hash, err := hashPassword(password, cost)
	fmt.Println("Password is: ", hash)
}

func hashPassword(password string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
