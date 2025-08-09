package gitlib

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func CreateWebHook(repoUrl, user_pat, reportUrl string) error {
	createWebHookUrl := "https://api.github.com/repos/OWNER/REPO/hooks"

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	parts := strings.Split(repoUrl, "/")
	owner := parts[1]
	repo := strings.ReplaceAll(parts[2], ".git", "")

	createWebHookUrl = strings.ReplaceAll(createWebHookUrl, "OWNER", owner)
	createWebHookUrl = strings.ReplaceAll(createWebHookUrl, "REPO", repo)

	fmt.Println(owner)
	fmt.Println(repo)
	fmt.Println(createWebHookUrl)

	requestBody := WebHookRequest{
		Owner:  owner,
		Repo:   repo,
		Name:   "web",
		Active: true,
		Events: []string{"push"},
		Config: Config{
			URL:         reportUrl,
			ContentType: "json",
			InsecureSSL: "1",
		},
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Println("Couldn't send request")
		return err
	}

	req, err := http.NewRequest("POST", createWebHookUrl, bytes.NewBuffer(jsonData))

	req.Header.Set("Authorization", "Bearer "+user_pat)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("Made request")
	// var recivedJson WebHookRequest
	// if err = json.Unmarshal(); err != nil {

	// }
	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading res body")
		return err
	}
	fmt.Println(string(resBody))
	return nil
}
