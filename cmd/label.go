package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/itscaro/gitlab-utils/git"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var labelCmdOpts struct {
	configFile     string
	project        string
	mergeRequestID int
}

func createLabelCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:           "label",
		Short:         "",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE:          runLabelCmd,
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			printMemory()
		},
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Could not determine working directory")
	}
	cmd.Flags().StringVarP(&labelCmdOpts.configFile, "config", "f", filepath.Join(dir, "label.yml"), "Project")
	cmd.Flags().StringVarP(&labelCmdOpts.project, "project", "p", "", "Project")
	_ = cmd.MarkPersistentFlagRequired("project")
	cmd.Flags().IntVarP(&labelCmdOpts.mergeRequestID, "merge-request-id", "i", 0, "Merge Request ID")
	_ = cmd.MarkPersistentFlagRequired("merge-request-id")

	return cmd
}

func runLabelCmd(cmd *cobra.Command, args []string) error {
	createClient()

	var config map[string][]string
	if data, err := ioutil.ReadFile(labelCmdOpts.configFile); err != nil {
		return err
	} else {
		if err := yaml.Unmarshal(data, &config); err != nil {
			return err
		}
	}

	return git.Label(config, labelCmdOpts.project, labelCmdOpts.mergeRequestID)
}
