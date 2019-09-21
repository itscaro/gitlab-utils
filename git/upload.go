// Copyright (c) 2019.
// Author: Quan TRAN

package git

import (
	"errors"
	"fmt"
	"strings"

	"github.com/xanzy/go-gitlab"
)

func Upload(projectUrl string, project string, file string) (fileUrl string, err error) {
	uf, _, err := client.Projects.UploadFile(project, file)
	if err != nil {
		return fileUrl, err
	}
	if uf == nil {
		return fileUrl, errors.New(fmt.Sprintf("could not upload file %s for project %s", file, project))
	} else {
		fileUrl = projectUrl + uf.URL
		fmt.Printf("File %s was uploaded to %s\nMarkdown: %s\n", file, fileUrl, uf.Markdown)
	}

	return fileUrl, nil
}

func UploadAsset(projectUrl string, project string, tag string, name string, file string) error {
	fileUrl, err := Upload(projectUrl, project, file)
	if err != nil {
		return err
	}

	if len(name) == 0 {
		parts := strings.Split("/", fileUrl)
		name = parts[len(parts)-1]
	}

	opt := gitlab.CreateReleaseLinkOptions{
		Name: gitlab.String(name),
		URL:  gitlab.String(fileUrl),
	}
	rl, _, err := client.ReleaseLinks.CreateReleaseLink(project, tag, &opt)
	if err != nil {
		return err
	}
	if rl == nil {
		return errors.New(fmt.Sprintf("could not create release link for file %s", fileUrl))
	} else {
		fmt.Printf("File %s was created as release link of tag %s\n", fileUrl, tag)
	}

	return nil
}
