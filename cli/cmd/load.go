package cmd

import (
    "errors"
    "fmt"

    "github.com/spf13/cobra"

    "godel/operations"
    "godel/types"
)

var loadCmd = &cobra.Command{
	Use:   "load <filename> [<filename> ...]",
	Short: "Load data into the Gödel system",
	Long: `Load various kinds of data into the Gödel system.`,
    Example: `# Load Insteon device data
load insteon-devices.json
`,
	RunE: load,
}

func load(cmd *cobra.Command, args []string) error {
    var devicesByAddress map[string]types.Device
    var err error

    if len(args) < 1 {
        return errors.New("At least one filename is required")
    }

    devicesByAddress, err = operations.Load(args)
    if err != nil {
        return err
    }

    for _, device := range devicesByAddress {
        fmt.Printf("Address: %s\n", device.Address)
        types.PrintDevice(device, "  ")
    }

    return nil
}

func init() {
	RootCmd.AddCommand(loadCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
