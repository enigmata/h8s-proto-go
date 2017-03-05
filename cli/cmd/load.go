package cmd

import (
	"fmt"

	"godel/device"
	"godel/operations"
)

var loadCmd = &Command{
	Name:    "load",
	Handler: load,
	Flags:   nil,
	UsageHelp: `load <filename> [<filename> ...]

Load various kinds of data into the Gödel system.`,
	Description: "Load data into the Gödel system",
	Examples: `# Load Insteon device data
load insteon-devices.json`,
}

func load(cmd *Command, args Args) {
	var devicesByAddress map[string]device.Device
	var err error

	if len(args) < 1 {
		fmt.Println("Error: At least one filename is required\n")
		cmd.PrintUsage()
		return
	}

	devicesByAddress, err = operations.Load(args)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	for _, dev := range devicesByAddress {
		fmt.Printf("Address: %s\n", dev.Address)
		dev.Print("  ")
	}

	return
}

func init() {
	RootCmd.AddSubCommand(loadCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
