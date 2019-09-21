package cmd

import (
	"fmt"
	"log"
	"os"
	"runtime"

	"github.com/itscaro/gitlab-utils/git"
	"github.com/spf13/cobra"
)

func Execute() {
	if err := createRootCmd().Execute(); err != nil {
		fmt.Printf("Error: %s\n", err)
		os.Exit(1)
	}
}

func createRootCmd() *cobra.Command {
	var cmd = &cobra.Command{
		Use:           "gitlab-utils",
		Short:         "",
		SilenceUsage:  true,
		SilenceErrors: true,
		RunE:          runRootCmd,
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			printMemory()
		},
	}

	cmd.AddCommand(
		createLabelCmd(),
		createUploadCmd(),
		createUploadAssetCmd(),
	)
	return cmd
}

func runRootCmd(cmd *cobra.Command, args []string) error {
	return cmd.Help()
}

func createClient() {
	endpoint := os.Getenv("GITLAB_ENDPOINT")
	if len(endpoint) == 0 {
		log.Fatalln("GITLAB_ENDPOINT must be set")
	}
	token := os.Getenv("GITLAB_TOKEN")
	if len(token) == 0 {
		log.Fatalln("GITLAB_TOKEN must be set")
	}

	git.NewClient(endpoint, token)
}

func printMemory() {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	fmt.Printf("Memory (total allocation) %.2f MB\n", float64(m.TotalAlloc)/1000000)
}
