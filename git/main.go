package git

import "github.com/xanzy/go-gitlab"

var client *gitlab.Client

func NewClient(endpoint string, token string) {
	client = gitlab.NewClient(nil, token)
	client.SetBaseURL(endpoint)
}
