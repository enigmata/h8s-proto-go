package main

import (
	"fmt"

	"godel/cmd"
	"godel/version"
)

var versionCmd = &cmd.Command{
	Name:    "version",
	Handler: versionPrint,
	UsageHelp: `version

Print all version information.`,
	Description: "Print all version information",
}

func versionPrint(cmd *cmd.Command, args cmd.Args) {
	fmt.Println("Client:")
	version.GetVersion().Print("  ")
	fmt.Println("Daemon:")
	return
}

func init() {
	rootCmd.AddSubCommand(versionCmd)
}
