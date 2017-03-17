package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	v1api "godel/api/http/v1"
	"godel/cmd"
	"godel/discovery"
	"godel/version"
)

var flags struct {
	version, help bool
}

var rootCmd = cmd.Command{
	Name:    "godeld",
	Handler: startDaemon,
	FlagSet: flag.NewFlagSet("godeld", flag.ContinueOnError),
	UsageHelp: `Daemon of the Gödel system.

The Gödel system is a mesh network of devices comprising an automation
environment for the betterment of lifestyle. The mesh is formed by the
interconnection of Gödel daemons.`,
	Description: "Daemon of the Gödel system",
}

func init() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lmicroseconds | log.Lshortfile)
	log.SetOutput(os.Stdout)

	rootCmd.FlagSet.BoolVar(&flags.version, "version", false, "version information")
	rootCmd.FlagSet.BoolVar(&flags.help, "help", false, "command usage")
}

func startDaemon(cmd *cmd.Command, args cmd.Args) {
	if flags.help {
		cmd.PrintUsage()
		return
	}
	if flags.version {
		version.GetVersion().Print("")
		return
	}

	daemonListenAddr := discovery.GetDaemonAddress()

	log.Printf("INFO: Daemon is starting, and will listen @ %s ...", daemonListenAddr)

	mux := http.NewServeMux()

	v1api.InstallHandlers(mux)

	s := &http.Server{
		Addr:           daemonListenAddr,
		Handler:        mux,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	log.Fatal(s.ListenAndServe())

	return
}

func main() {
	rootCmd.Run()
}
