package cmd

import (
    "fmt"

    "github.com/spf13/cobra"
)

var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Retrieve various aspects of the Gödel system",
	Long: `Retrieve various aspects fo the Gödel system, including
state and configuration.`,
    Example: `# Retrieve all the thingys
get
`,
	Run: get,
}

func get(cmd *cobra.Command, args []string) {
    fmt.Println("voila")
    return
}

func init() {
	RootCmd.AddCommand(getCmd)

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
