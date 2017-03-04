package cmd

import (
	"fmt"
)

var getCmd = &Command{
	Name:    "get",
	Handler: get,
	Flags:   nil,
	UsageHelp: `get <thing>

Retrieve various aspects fo the Gödel system, including
state and configuration.`,
	Description: "Retrieve various aspects of the Gödel system",
	Examples: `
# Retrieve all the thingys
get`,
}

func get(args Args) {
	fmt.Println("get: voila")
	return
}

func init() {
	RootCmd.AddSubCommand(getCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
