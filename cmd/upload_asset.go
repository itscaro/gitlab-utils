package cmd

import (
	"log"
	"os"
	"path/filepath"

	"github.com/itscaro/gitlab-utils/git"
	"github.com/spf13/cobra"
)

var uploadAssetCmdOpts struct {
	configFile   string
	project      string
	tag          string
	fileToUpload string
}

func createUploadAssetCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:           "upload-asset",
		Short:         "",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE:          runUploadAssetCmd,
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			printMemory()
		},
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Could not determine working directory")
	}
	cmd.Flags().StringVarP(&uploadAssetCmdOpts.configFile, "config", "f", filepath.Join(dir, "label.yml"), "Project")
	cmd.Flags().StringVarP(&uploadAssetCmdOpts.project, "project", "p", "", "Project")
	_ = cmd.MarkPersistentFlagRequired("project")
	cmd.Flags().StringVarP(&uploadAssetCmdOpts.tag, "tag", "t", "", "Tag")
	_ = cmd.MarkPersistentFlagRequired("tag")
	cmd.Flags().StringVarP(&uploadAssetCmdOpts.fileToUpload, "file", "i", "", "The file to upload")
	_ = cmd.MarkPersistentFlagRequired("file")

	return cmd
}

func runUploadAssetCmd(cmd *cobra.Command, args []string) error {
	createClient()

	return git.UploadAsset(uploadAssetCmdOpts.project, uploadAssetCmdOpts.tag, uploadAssetCmdOpts.fileToUpload)
}