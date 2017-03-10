package main

import (
	"log"
	"os"

	"godel/cmd"
)

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Llongfile)
	log.SetOutput(os.Stdout)
}

func main() {
	rootCmd.Run()
}

var rootCmd = cmd.Command{
	Name:    "godeld",
	Handler: startDaemon,
	Flags:   nil,
	UsageHelp: `Daemon of the Gödel system.

The Gödel system is a mesh network of devices comprising an automation
environment for the betterment of lifestyle. The mesh is formed by the
interconnection of Gödel daemons.`,
	Description: "Daemon of the Gödel system",
}

func startDaemon(cmd *cmd.Command, args cmd.Args) {
	log.Println("INFO: Daemon is starting ...")
	return
}
