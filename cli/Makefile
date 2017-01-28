# Command-line tools
GO := go
MOVE := mv

# Buildable executables
CLI_BUILD_NAME := cli
CLI_NAME := godel

all: install

install:
	$(GO) install
	$(MOVE) $(GOBIN)/$(CLI_BUILD_NAME) $(GOBIN)/$(CLI_NAME) 
