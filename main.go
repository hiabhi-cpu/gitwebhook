package main

import (
	"log"
	"os"

	gitlib "github.com/hiabhi-cpu/gitwebhook/gitLib"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	user_PAT := os.Getenv("GIT_PAT")
	gitRepo := "github.com/hiabhi-cpu/webHookTry"
	gitlib.GetOrCreateWebhook(gitRepo, user_PAT)
}
