// Copyright (c) 2019.
// Author: Quan TRAN

package utils

import "runtime"

var (
	GitCommit string
	Version   string
)

func GetVersion() string {
	version := Version
	if len(version) == 0 {
		version = "dev"
	}

	return version
}

func GetCommit() string {
	commit := GitCommit
	if len(commit) == 0 {
		commit = "dev"
	} else {
		commit = commit[:8]
	}

	return commit
}

func GetRuntimeVersion() string {
	return runtime.Version()
}
