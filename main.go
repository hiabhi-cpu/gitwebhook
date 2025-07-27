package main

import (
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
	user_PAT := os.Getenv("GIT_PAT")
	gitRepo := "github.com/hiabhi-cpu/webHookTry"
	fmt.Println("Hello")
	err = CreateWebHook(gitRepo, user_PAT)
	if err != nil {
		panic(err)
	}
}
