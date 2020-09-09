package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"path"

	v1api "github.com/enigmata/h8s-proto-go/api/http/v1"
	"github.com/enigmata/h8s-proto-go/cmd"
	"github.com/enigmata/h8s-proto-go/discovery"
	"github.com/enigmata/h8s-proto-go/version"
)

var versionCmd = &cmd.Command{
	Name:    "version",
	Handler: versionPrint,
	UsageHelp: `version

Print all version information.`,
	Description: "Print all version information",
}

func versionPrint(cmd *cmd.Command, args cmd.Args) {
	var daemonVer version.VersionInfo

	fmt.Println("Client:")
	version.GetVersion().Print("  ")

	res, err := http.Get("http://" + path.Join(discovery.GetDaemonAddress(), v1api.UriPathAdminVersion))
	if err != nil {
		fmt.Printf("Error: Cannot send request to daemon for its version: %s\n", err.Error())
		return
	}
	defer res.Body.Close()
	ver, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error: Cannot extract daemon's version info from response: %s\n", err.Error())
		return
	}

	fmt.Printf("Daemon (@ %s):\n", discovery.GetDaemonAddress())
	err = daemonVer.Unmarshal(ver)
	if err != nil {
		fmt.Printf("Error: Cannot extract daemon's version info: %s\n", err.Error())
		return
	}
	daemonVer.Print("  ")

	return
}

func init() {
	rootCmd.AddSubCommand(versionCmd)
}
