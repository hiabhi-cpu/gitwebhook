package gitlib

import (
	"fmt"
	"strings"
)

func GetOrCreateWebhook(gitRepo, user_PAT, rev_url string) error {

	fmt.Println("Hello")
	gitJsonReply, err := GetWebHook(gitRepo, user_PAT)
	if err != nil {
		return err
	}

	fmt.Println(len(gitJsonReply))
	if len(gitJsonReply) == 0 {
		err = CreateWebHook(gitRepo, user_PAT, rev_url)
		if err != nil {
			return err
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
			return err
		}
	}
	return nil
}
