package main

import (
	"adconnector/adapi"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println("Hello World!")

	boo, err := adapi.NewADHelperWithClientSecret(os.Getenv("TENANT_ID"), os.Getenv("CLIENT_ID"), os.Getenv("CLIENT_SECRET"))
	if err != nil {
		log.Fatal(err)

	}
	users, err := boo.GetUsers()
	if err != nil {
		log.Fatal(err)

	}
	for _, user := range users.GetValue() {
		fmt.Println(*user.GetDisplayName())
	}

}
