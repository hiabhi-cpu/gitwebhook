package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func GetWebHookConfig(repoUrl, user_pat, HOOK_ID string) error {

	getWebHookConfigUrl := "https://api.github.com/repos/OWNER/REPO/hooks/HOOK_ID"

	client := &http.Client{
		Timeout: time.Second * 10,
	}
	parts := strings.Split(repoUrl, "/")
	owner := parts[1]
	repo := strings.ReplaceAll(parts[2], ".git", "")

	getWebHookConfigUrl = strings.ReplaceAll(getWebHookConfigUrl, "OWNER", owner)
	getWebHookConfigUrl = strings.ReplaceAll(getWebHookConfigUrl, "REPO", repo)
	getWebHookConfigUrl = strings.ReplaceAll(getWebHookConfigUrl, "HOOK_ID", HOOK_ID)

	req, err := http.NewRequest("GET", getWebHookConfigUrl, nil)

	req.Header.Set("Authorization", "Bearer "+user_pat)
	req.Header.Set("Accept", "application/vnd.github+json")
	req.Header.Set("X-GitHub-Api-Version", "2022-11-28")
	req.Header.Set("Content-Type", "application/json")

	resp, err := client.Do(req)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	resBody, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error reading res body")
		return err
	}
	fmt.Println(string(resBody))

	return nil
}
