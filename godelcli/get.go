package main

import (
	"godel/cmd"
	"godel/controller"
)

var getCmd = &cmd.Command{
	Name:    "get",
	Handler: get,
	UsageHelp: `get <thing>

Retrieve various aspects fo the Gödel system, including
state and configuration.`,
	Description: "Retrieve various aspects of the Gödel system",
	Examples: `# Retrieve all the thingys
get`,
}

func get(cmd *cmd.Command, args cmd.Args) {

	isy := controller.NewClient(controller.ISY994)
	devices := isy.Devices()
	for _, dev := range devices {
		dev.Print("")
	}

	return
}

func init() {
	rootCmd.AddSubCommand(getCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
