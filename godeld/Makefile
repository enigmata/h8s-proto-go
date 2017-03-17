# overrideable by env vars
DAEMON_VERSION ?= NO_VERSION

# internal env vars
GIT_HASH = `git log -n 1 --pretty=format:%H`
BUILD_DATE = `date +%Y.%m.%d.%H%M%S`

# simple constants
DAEMON_EXE := godeld
GO := go
REMOVE := rm -f

.PHONY: install

default: install

install:
	$(REMOVE) $(GOBIN)/$(DAEMON_EXE)
	go install -ldflags "-X godel/version.buildTime=${BUILD_DATE} -X godel/version.gitCommit=${GIT_HASH} -X godel/version.versionNum=${DAEMON_VERSION}"
