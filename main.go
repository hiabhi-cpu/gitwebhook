package main

import (
	"fmt"
	"log"
	"os"
	"strings"

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
	gitJsonReply, err := GetWebHook(gitRepo, user_PAT)
	if err != nil {
		panic(err)
	}
	rev_url := os.Getenv("REV_URL")
	fmt.Println(len(gitJsonReply))
	if len(gitJsonReply) == 0 {
		err = CreateWebHook(gitRepo, user_PAT, rev_url)
		if err != nil {
			panic(err)
		}
	}
	cnt := 0
	for _, r := range gitJsonReply {
		fmt.Println(r.Config.URL)
		fmt.Println(rev_url)
		if strings.Trim(r.Config.URL, " ") == rev_url {
			cnt++
		}
	}
	fmt.Println(cnt, "cnt")
	if cnt == 0 {
		err = CreateWebHook(gitRepo, user_PAT, rev_url)
		if err != nil {
			panic(err)
		}
	}
}
