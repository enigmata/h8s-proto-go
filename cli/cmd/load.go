package cmd

import (
	"fmt"

	"godel/operations"
	"godel/types"
)

var loadCmd = Command{
	Name:    "load",
	Handler: load,
	Flags:   nil,
	UsageHelp: `load <filename> [<filename> ...]

Load various kinds of data into the Gödel system.`,
	Description: "Load data into the Gödel system",
	Examples: `# Load Insteon device data
load insteon-devices.json`,
}

func load(args Args) {
	fmt.Println("load: voila")
	var devicesByAddress map[string]types.Device
	var err error

	if len(args) < 1 {
		fmt.Println("Error: At least one filename is required")
		return
	}

	devicesByAddress, err = operations.Load(args)
	if err != nil {
		fmt.Printf("Error: %s\n", err.Error())
		return
	}

	for _, device := range devicesByAddress {
		fmt.Printf("Address: %s\n", device.Address)
		types.PrintDevice(device, "  ")
	}

	return
}

func init() {
	//RootCmd.AddCommand(loadCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
