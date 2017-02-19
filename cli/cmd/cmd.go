package cmd

import (
    "flag"
    "fmt"
    "text/tabwriter"
    "os"
)


type commandHandler func (*commandFlags)

type commandFlags struct {
    flagSet *flag.FlagSet
    flags interface{}
}

type commands map[string]*command

type command struct {
    name string
    handler commandHandler
    subCommands commands
    usageHelp string
    description string
    examples string
    flags *commandFlags
}

func (c command) printUsage() {
    fmt.Println(c.usageHelp)
    fmt.Println("Commands:")
    w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
    for name, subCmd := range c.subCommands {
        fmt.Fprintf(w, "%s\t%s\n", name, subCmd.description)
    }
    w.Flush()
    fmt.Println()
}

const (
    cliName = "godel"
    cliDescription string = "Command-line interface to the Gödel system"
    cliUsage string = `
Command-line interface to the Gödel system.

Provides a means to interact with the Gödel system from a terminal
session, or enables scripted automation of various aspects
and behaviours of the Gödel system.

The Gödel system is a mesh network of devices comprising an automation
environment for the betterment of lifestyle.
`
    cliExamples string = `
# Retrieve all the thingys
godel get ...
`

    getCmdName = "get"
    getCmdDescription string = "Retrieve various aspects of the Gödel system"
    getCmdUsage string = `get <thing>

Retrieve various aspects fo the Gödel system, including
state and configuration.
`
    getCmdExamples string = `
# Retrieve all the thingys
get
`

    loadCmdName = "load"
    loadCmdDescription string = "Load data into the Gödel system"
    loadCmdUsage string = `load <filename> [<filename> ...]

Load various kinds of data into the Gödel system.
`
    loadCmdExamples string = `# Load Insteon device data
load insteon-devices.json
`
)

var cmds = command{
    name: cliName,
    handler: nil,
    subCommands: commands{
        "get": &command{
            name: getCmdName,
            handler: nil,
            subCommands: nil,
            flags: nil,
            usageHelp: getCmdUsage,
            description: getCmdDescription,
            examples: getCmdExamples,
        },
        "load": &command{
            name: loadCmdName,
            handler: nil,
            subCommands: nil,
            flags: nil,
            usageHelp: loadCmdUsage,
            description: loadCmdDescription,
            examples: loadCmdExamples,
        },
    },
    flags: nil,
    usageHelp: cliUsage,
    description: cliDescription,
    examples: cliExamples,
}

func Run() {
    var err error
    var cmd, subCmd *command
    var ok bool
    var i int = 1

    cmd = &cmds

    if len(os.Args) <= 1 {
        cmd.printUsage()
        return
    }

    for {
        if cmd.handler != nil {
            if cmd.flags != nil && cmd.flags.flagSet != nil {
                err = cmd.flags.flagSet.Parse(os.Args[i:])
                if err != nil {
                    fmt.Println("Error: Could not parse command-line args: " + err.Error())
                    return
                }
            }
            cmd.handler(cmd.flags)
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
                    cmd.printUsage()
                    return
                }
            } else {
                fmt.Printf("\nError: Missing subcommand of command \"%s\"\n\nUsage:\n", cmd.name)
                cmd.printUsage()
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
