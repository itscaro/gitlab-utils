package git

import "github.com/xanzy/go-gitlab"

var client *gitlab.Client

func NewClient(endpoint string, token string) {
	client, _ = gitlab.NewClient(token, gitlab.WithBaseURL(endpoint))
}
