package cmd

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"runtime"

	"github.com/itscaro/gitlab-labeler/git"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v2"
)

var rootCmdOpts struct {
	configFile     string
	project        string
	mergeRequestID int
}

func createRootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:           "labeler",
		Short:         "",
		SilenceUsage:  true,
		SilenceErrors: true,
		PreRunE:       prerunCmd,
		RunE:          runCmd,
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			printMemory()
		},
	}

	dir, err := os.Getwd()
	if err != nil {
		log.Fatalln("Could not determine working directory")
	}
	cmd.Flags().StringVarP(&rootCmdOpts.configFile, "config", "f", filepath.Join(dir, "label.yml"), "Project")
	cmd.Flags().StringVarP(&rootCmdOpts.project, "project", "p", "", "Project")
	cmd.Flags().IntVarP(&rootCmdOpts.mergeRequestID, "mergerRequestID", "i", 0, "Merge Request ID")

	return cmd
}

func Execute() {
	if err := createRootCmd().Execute(); err != nil {
		fmt.Printf("%s\n", err)
		os.Exit(1)
	}
}

func prerunCmd(cmd *cobra.Command, args []string) error {
	if len(rootCmdOpts.project) == 0 {
		return errors.New("project is not valid")
	}
	if rootCmdOpts.mergeRequestID == 0 {
		return errors.New("merge request ID is not valid")
	}
	return nil
}

func runCmd(cmd *cobra.Command, args []string) error {
	endpoint := os.Getenv("GITLAB_ENDPOINT")
	if len(endpoint) == 0 {
		log.Fatalln("GITLAB_ENDPOINT must be set")
	}
	token := os.Getenv("GITLAB_TOKEN")
	if len(token) == 0 {
		log.Fatalln("GITLAB_ENDPOINT must be set")
	}

	var config map[string][]string
	if data, err := ioutil.ReadFile(rootCmdOpts.configFile); err != nil {
		return err
	} else {
		if err := yaml.Unmarshal(data, &config); err != nil {
			return err
		}
	}

	git.NewClient(endpoint, token)

	return git.Label(config, rootCmdOpts.project, rootCmdOpts.mergeRequestID)
}

func printMemory() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory (total allocation) %.2f MB\n", float64(m.TotalAlloc)/1000000)
}
