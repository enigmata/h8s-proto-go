package main

import (
	"log"
	"net/http"
	"os"
	"time"

	v1api "godel/api/http/v1"
	"godel/cmd"
)

const (
	daemonIP         string = "0.0.0.0"
	daemonPort       string = "19999"
	daemonListenAddr string = daemonIP + ":" + daemonPort
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
	log.Printf("INFO: Daemon is starting, and will listen @ %s ...", daemonListenAddr)

	mux := http.NewServeMux()

	mux.HandleFunc(v1api.UriPathAdminVersion, v1api.AdminVersionHandler)

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
