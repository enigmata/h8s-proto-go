package cmd

import (
	"flag"
	"fmt"
	"os"
	"text/tabwriter"
)

type CommandHandler func(Args)

type Args []string

type CommandFlags struct {
	FlagSet *flag.FlagSet
	Flags   interface{}
}

type Commands map[string]*Command

type Command struct {
	Name        string
	Handler     CommandHandler
	UsageHelp   string
	Description string
	Examples    string
	Flags       *CommandFlags

	// pkg-internal vars
	subCommands Commands
}

func (c Command) PrintUsage() {
	fmt.Println(c.UsageHelp)
	if len(c.subCommands) > 0 {
		fmt.Println("Commands:")
		w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
		for name, subCmd := range c.subCommands {
			fmt.Fprintf(w, "%s\t%s\n", name, subCmd.Description)
		}
		w.Flush()
	}
	if c.Examples != "" {
		fmt.Println("\nExamples:")
		fmt.Println(c.Examples)
		fmt.Println()
	}
}

func (c *Command) Run() {
	var err error
	var cmd, subCmd *Command
	var ok bool
	var i int = 1

	cmd = c

	if len(os.Args) <= 1 {
		cmd.PrintUsage()
		return
	}

	for {
		if cmd.Handler != nil {
			if cmd.Flags != nil && cmd.Flags.FlagSet != nil {
				err = cmd.Flags.FlagSet.Parse(os.Args[i:])
				if err != nil {
					fmt.Println("Error: Could not parse command-line args: " + err.Error())
					return
				}
			}
			cmd.Handler(os.Args[i:])
			return
		} else if cmd.subCommands != nil {
			if i < len(os.Args) {
				subCmd, ok = cmd.subCommands[os.Args[i]]
				if ok {
					cmd = subCmd
					i += 1
					continue
				} else {
					fmt.Printf("\nError: Invalid command: \"%s\"\n\nUsage:\n", os.Args[i])
					cmd.PrintUsage()
					return
				}
			} else {
				fmt.Printf("\nError: Missing subcommand of command \"%s\"\n\nUsage:\n", cmd.Name)
				cmd.PrintUsage()
				return
			}
		} else {
			// should not arrive here because to do so
			// means the commands are not defined properly
			fmt.Println("Internal Error: Commands broken")
			return
		}
	}
}
