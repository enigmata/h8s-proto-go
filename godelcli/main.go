package main

import (
	"godel/cmd"
)

func main() {
	rootCmd.Run()
}

var rootCmd = cmd.Command{
	Name:    "godel",
	Handler: nil,
	UsageHelp: `Command-line interface to the Gödel system.

Provides a means to interact with the Gödel system from a terminal
session, or enables scripted automation of various aspects
and behaviours of the Gödel system.

The Gödel system is a mesh network of devices comprising an automation
environment for the betterment of lifestyle.`,
	Description: "Command-line interface to the Gödel system",
	Examples: `# Retrieve all the thingys
godel get ... `,
}
