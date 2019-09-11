package main

import (
	"fmt"

	"github.com/itscaro/gitlab-utils/cmd"
	"github.com/itscaro/gitlab-utils/utils"
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
