// Copyright (c) 2019.
// Author: Quan TRAN

package git

import (
	"errors"
	"fmt"
)

func UploadAsset(project string, file string) error {
	res, _, err := client.Projects.UploadFile(project, file)
	if err != nil {
		return err
	}
	if res == nil {
		return errors.New(fmt.Sprintf("could not upload file %s for project %s", file, project))
	} else {
		fmt.Printf("File %s was uploaded to %s\n", file, res.URL)
	}

	return nil
}
