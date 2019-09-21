package git

import "github.com/xanzy/go-gitlab"

var client *gitlab.Client
var projectUrl string

func NewClient(endpoint string, token string) {
	client = gitlab.NewClient(nil, token)
	client.SetBaseURL(endpoint)
}

func ProjectUrl(url string) {
	projectUrl = url
}
