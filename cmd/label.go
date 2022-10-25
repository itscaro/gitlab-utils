package cmd

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strconv"

	"github.com/itscaro/gitlab-utils/git"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
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
		PreRunE:       preRunLabelCmd,
		RunE:          runLabelCmd,
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			printMemory()
		},
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Could not determine working directory")
	}
	cmd.Flags().StringVarP(&labelCmdOpts.configFile, "config", "c", filepath.Join(dir, "label.yaml"), "Project")
	cmd.Flags().StringVarP(&labelCmdOpts.project, "project", "p", "", "Project [default: $CI_PROJECT_PATH]")
	_ = cmd.MarkPersistentFlagRequired("project")
	cmd.Flags().IntVarP(&labelCmdOpts.mergeRequestID, "merge-request-id", "i", 0, "Merge Request ID [default: $CI_MERGE_REQUEST_IID]")
	_ = cmd.MarkPersistentFlagRequired("merge-request-id")

	return cmd
}

func preRunLabelCmd(cmd *cobra.Command, args []string) error {
	var err error
	if labelCmdOpts.project == "" {
		labelCmdOpts.project = os.Getenv("CI_PROJECT_PATH")
	}

	if labelCmdOpts.mergeRequestID == 0 {
		labelCmdOpts.mergeRequestID, err = strconv.Atoi(os.Getenv("CI_MERGE_REQUEST_IID"))
	}

	return err
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
