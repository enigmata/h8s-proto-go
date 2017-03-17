package version

import (
	"encoding/json"
	"fmt"
	"runtime"
)

var (
	// these variables are overwritten during buildtime
	versionNum string = "build-import-versionnum"
	gitCommit  string = "build-import-gitcommit"
	buildTime  string = "build-import-buildtime"
)

type VersionInfo struct {
	VersionNum string
	GitCommit  string
	BuildTime  string
	Arch       string
	OS         string
}

func GetVersion() *VersionInfo {
	return &VersionInfo{
		VersionNum: versionNum,
		GitCommit:  gitCommit,
		BuildTime:  buildTime,
		Arch:       runtime.GOARCH,
		OS:         runtime.GOOS,
	}
}

func (v VersionInfo) Print(prefix string) {
	fmt.Println()
	fmt.Printf("%sVersion Number:   %s\n", prefix, v.VersionNum)
	fmt.Printf("%sGit Commit:       %s\n", prefix, v.GitCommit)
	fmt.Printf("%sBuild Time:       %s\n", prefix, v.BuildTime)
	fmt.Printf("%sArchitecture:     %s\n", prefix, v.Arch)
	fmt.Printf("%sOperating System: %s\n", prefix, v.OS)
	fmt.Println()
	return
}

func (v VersionInfo) Marshal() string {
	b, _ := json.Marshal(v)
	return string(b)
}
