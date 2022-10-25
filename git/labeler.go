// Copyright (c) 2019.
// Author: Quan TRAN

package git

import (
	"errors"
	"fmt"

	"github.com/gobwas/glob"
	"github.com/itscaro/gitlab-utils/utils"
	"github.com/xanzy/go-gitlab"
)

func Label(config map[string][]string, project string, mergeRequest int) error {
	fmt.Printf("Config %s\n", config)

	mr, _, err := client.MergeRequests.GetMergeRequestChanges(project, mergeRequest, nil)
	if err != nil {
		return err
	}
	if mr == nil {
		return errors.New(fmt.Sprintf("could not fetch merge request %d for project %s", mergeRequest, project))
	}

	var labelsToApply gitlab.Labels

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
					break
				}
			}
		}
		if len(config) == 0 {
			// All labels applied, stop the loop
			break
		}
	}

	fmt.Printf("Done after %d iterations / %d changes\n", counter, len(mr.Changes))

	// Set new labels
	if len(labelsToApply) == 0 {
		fmt.Printf("No label to apply\n")
	} else {
		fmt.Printf("Going to apply %s\n", labelsToApply)
		mr, _, err := client.MergeRequests.GetMergeRequest(project, mergeRequest, nil)
		if err != nil {
			return err
		}
		labels := gitlab.Labels(utils.Unique(append(mr.Labels, labelsToApply...)))
		opts := gitlab.UpdateMergeRequestOptions{
			Labels: &labels,
		}
		if _, _, err := client.MergeRequests.UpdateMergeRequest(project, mergeRequest, &opts); err != nil {
			return err
		}
	}

	return nil
}
