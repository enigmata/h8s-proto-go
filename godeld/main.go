package main

import (
	"godel/cmd"
)

func main() {
	rootCmd.Run()
}

var rootCmd = cmd.Command{
	Name:    "godeld",
	Handler: nil,
	Flags:   nil,
	UsageHelp: `Daemon of the Gödel system.

The Gödel system is a mesh network of devices comprising an automation
environment for the betterment of lifestyle. The mesh is formed by the
interconnection of Gödel daemons.`,
	Description: "Daemon of the Gödel system",
}
