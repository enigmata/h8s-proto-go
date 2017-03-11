# overrideable by env vars
CLI_VERSION ?= NO_VERSION

# internal env vars
GIT_HASH = `git log -n 1 --pretty=format:%H`
BUILD_DATE = `date +%Y.%m.%d.%H%M%S`

# simple constants
CLI_EXE := godelcli
GO := go
REMOVE := rm -f

.PHONY: install

default: install

install:
	$(REMOVE) $(GOBIN)/$(CLI_EXE)
	go install -ldflags "-X main.buildTime=${BUILD_DATE} -X main.gitCommit=${GIT_HASH} -X main.cliVersion=${CLI_VERSION}"
