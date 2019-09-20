// Copyright (c) 2019.
// Author: Quan TRAN

package git

import (
	"errors"
	"fmt"

	"github.com/xanzy/go-gitlab"
)

func Upload(project string, file string) error {
	uf, _, err := client.Projects.UploadFile(project, file)
	if err != nil {
		return err
	}
	if uf == nil {
		return errors.New(fmt.Sprintf("could not upload file %s for project %s", file, project))
	} else {
		fmt.Printf("File %s was uploaded to %s\nMarkdown: %s\n", file, uf.URL, uf.Markdown)
	}

	return nil
}

func UploadAsset(project string, tag string, file string) error {
	uf, _, err := client.Projects.UploadFile(project, file)
	if err != nil {
		return err
	}
	if uf == nil {
		return errors.New(fmt.Sprintf("could not upload file %s for project %s", file, project))
	} else {
		fmt.Printf("File %s was uploaded to %s\nMarkdown: %s\n", file, uf.URL, uf.Markdown)
	}

	opt := gitlab.CreateReleaseLinkOptions{}
	opt.URL = &uf.URL
	rl, _, err := client.ReleaseLinks.CreateReleaseLink(project, tag, &opt)
	if err != nil {
		return err
	}
	if rl == nil {
		return errors.New(fmt.Sprintf("could not create release link for file %s", uf.URL))
	} else {
		fmt.Printf("File %s was created as release link of tag %s\n", uf.URL, tag)
	}

	return nil
}
