package main

import (
	"fmt"
    "log"
    "os"

    "github.com/joho/godotenv"
)

func main() {
	appId := os.Getenv("APPID")
	secretKey := os.Getenv("SECRET_KEY")
	fmt.Printf("appId:%s\r\n", appId)
	fmt.Printf("secretKey:%s\r\n", secretKey)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	appId = os.Getenv("APPID")
	secretKey = os.Getenv("SECRET_KEY")
	fmt.Printf("appId:%s\r\n", appId)
	fmt.Printf("secretKey:%s\r\n", secretKey)
}
