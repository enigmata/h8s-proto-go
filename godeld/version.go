package main

import (
	"fmt"
	"runtime"
)

var (
	daemonVersion string = "build-import-daemonversion"
	gitCommit     string = "build-import-gitcommit"
	buildTime     string = "build-import-buildtime"
)

func printVersion() {
	fmt.Println()
	fmt.Printf("Version:          %s\n", daemonVersion)
	fmt.Printf("Git Commit:       %s\n", gitCommit)
	fmt.Printf("Build Time:       %s\n", buildTime)
	fmt.Printf("Architecture:     %s\n", runtime.GOARCH)
	fmt.Printf("Operating System: %s\n", runtime.GOOS)
	fmt.Println()
	return
}
