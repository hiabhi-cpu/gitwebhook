package gitlib

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func GetWebHook(repoUrl, user_pat string) ([]GitJsonReply, error) {

	getWebHookUrl := "https://api.github.com/repos/OWNER/REPO/hooks"

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	parts := strings.Split(repoUrl, "/")
	owner := parts[1]
	repo := strings.ReplaceAll(parts[2], ".git", "")

	getWebHookUrl = strings.ReplaceAll(getWebHookUrl, "OWNER", owner)
	getWebHookUrl = strings.ReplaceAll(getWebHookUrl, "REPO", repo)

	req, err := http.NewRequest("GET", getWebHookUrl, nil)

	req.Header.Set("Authorization", "Bearer "+user_pat)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var gitJsonReply []GitJsonReply

	decoder := json.NewDecoder(resp.Body)

	if err := decoder.Decode(&gitJsonReply); err != nil {
		fmt.Println("error decoding response body")
		return nil, err
	}

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading res body")
		return nil, err
	}
	fmt.Println("hello abhi")
	fmt.Println(string(resBody))

	return gitJsonReply, nil
}
