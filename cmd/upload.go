package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/itscaro/gitlab-utils/git"
	"github.com/spf13/cobra"
)

var uploadCmdOpts struct {
	configFile   string
	projectUrl   string
	project      string
	fileToUpload string
}

func createUploadCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:           "upload",
		Short:         "",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE:          runUploadCmd,
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			printMemory()
		},
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Could not determine working directory")
	}
	cmd.Flags().StringVarP(&uploadCmdOpts.configFile, "config", "c", filepath.Join(dir, "label.yml"), "Project")
	cmd.Flags().StringVarP(&uploadCmdOpts.project, "project", "p", "", "Project")
	_ = cmd.MarkPersistentFlagRequired("project")
	cmd.Flags().StringVarP(&uploadCmdOpts.projectUrl, "project-url", "u", "", "Project Url")
	_ = cmd.MarkPersistentFlagRequired("project-url")
	cmd.Flags().StringVarP(&uploadCmdOpts.fileToUpload, "file", "f", "", "The file to upload")
	_ = cmd.MarkPersistentFlagRequired("file")

	return cmd
}

func runUploadCmd(cmd *cobra.Command, args []string) error {
	createClient()

	_, err := git.Upload(uploadCmdOpts.projectUrl, uploadCmdOpts.project, uploadCmdOpts.fileToUpload)
	return err
}
