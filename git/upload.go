// Copyright (c) 2019.
// Author: Quan TRAN

package git

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/xanzy/go-gitlab"
)

func Upload(projectUrl string, project string, file string) (fileUrl string, err error) {
	if !fileExists(file) {
		return "", errors.New(fmt.Sprintf("File %s does not exist", file))
	}

	f, err := os.Open(file)
	if err != nil {
		return fileUrl, err
	}

	uf, _, err := client.Projects.UploadFile(project, f, filepath.Base(file))
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

func UploadAsset(projectUrl string, project string, tag string, name string, file string, url string) error {
	if len(file) > 0 {
		var err error
		fileUrl, err := Upload(projectUrl, project, file)
		if err != nil {
			return err
		}
		url = fileUrl
	} else if len(url) == 0 {
		return errors.New("either file or url must be given")
	}

	if len(name) == 0 {
		parts := strings.Split(url, "/")
		name = parts[len(parts)-1]
	}

	opt := gitlab.CreateReleaseLinkOptions{
		Name: gitlab.String(name),
		URL:  gitlab.String(url),
	}
	rl, _, err := client.ReleaseLinks.CreateReleaseLink(project, tag, &opt)
	if err != nil {
		return err
	}
	if rl == nil {
		return errors.New(fmt.Sprintf("could not create release link for file %s", url))
	} else {
		fmt.Printf("File %s was created as release link of tag %s\n", url, tag)
	}

	return nil
}

func fileExists(filename string) bool {
	info, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return !info.IsDir()
}
