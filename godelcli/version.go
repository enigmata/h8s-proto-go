package main

import (
	"fmt"
	"runtime"

	"godel/cmd"
)

var (
	cliVersion string = "build-import-cliversion"
	gitCommit  string = "build-import-gitcommit"
	buildTime  string = "build-import-buildtime"
)

var versionCmd = &cmd.Command{
	Name:    "version",
	Handler: version,
	UsageHelp: `version

Print all version information.`,
	Description: "Print all version information",
}

func version(cmd *cmd.Command, args cmd.Args) {
	fmt.Printf("Version:          %s\n", cliVersion)
	fmt.Printf("Git Commit:       %s\n", gitCommit)
	fmt.Printf("Build Time:       %s\n", buildTime)
	fmt.Printf("Architecture:     %s\n", runtime.GOARCH)
	fmt.Printf("Operating System: %s\n", runtime.GOOS)
	return
}

func init() {
	rootCmd.AddSubCommand(versionCmd)
}
