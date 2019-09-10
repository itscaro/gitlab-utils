package main

import (
	"fmt"

	"github.com/itscaro/gitlab-labeler/cmd"
	"github.com/itscaro/gitlab-labeler/utils"
)

func main() {
	printVersion()
	cmd.Execute()
}

func printVersion() {
	fmt.Printf(
		"Gitlab Labeler (%s-%s) (Go %s)\n",
		utils.GetVersion(),
		utils.GetCommit(),
		utils.GetRuntimeVersion(),
	)
}
