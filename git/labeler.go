// Copyright (c) 2019.
// Author: Quan TRAN

package git

import (
	"errors"
	"fmt"
	"io/ioutil"

	"github.com/gobwas/glob"
	"github.com/xanzy/go-gitlab"
	"gopkg.in/yaml.v2"
)

func Label(project string, mergeRequest int) error {
	var config map[string][]string
	if data, err := ioutil.ReadFile("label.yml"); err != nil {
		return err
	} else {
		if err := yaml.Unmarshal(data, &config); err != nil {
			return err
		}
	}
	fmt.Printf("Config %s\n", config)

	mr, _, _ := client.MergeRequests.GetMergeRequestChanges(project, mergeRequest)
	if mr == nil {
		return errors.New(fmt.Sprintf("could not fetch merge request %d for project %s", mergeRequest, project))
	}

	var labelsToApply []string

	var g glob.Glob
	counter := 0
	for _, change := range mr.Changes {
		counter++
		for label, paths := range config {
			for _, p := range paths {

				g = glob.MustCompile(p)
				if g.Match(change.OldPath) || g.Match(change.NewPath) {
					fmt.Printf("Apply label %s because of %s or %s\n", label, change.OldPath, change.NewPath)
					labelsToApply = append(labelsToApply, label)
					delete(config, label)
				}
			}
		}
		if len(config) == 0 {
			// All labels applied, stop the loop
			break
		}
	}

	fmt.Printf("Done after %d iterations / %d changes\n", counter, len(mr.Changes))
	fmt.Printf("Going to apply %s\n", labelsToApply)

	// Set new labels
	opts := gitlab.UpdateMergeRequestOptions{
		Labels: labelsToApply,
	}
	if _, _, err := client.MergeRequests.UpdateMergeRequest(project, mergeRequest, &opts); err != nil {
		return err
	}

	return nil
}
