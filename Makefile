# overrideable by env vars
CLI_VERSION ?= 0.1.0

# internal env vars
GODELCLI_DIR = godelcli

# simple constants
MAKE := make
CLI_VERSION_ENV_VAR := CLI_VERSION

.PHONY: install

default: install

install:
	${CLI_VERSION_ENV_VAR}=${CLI_VERSION} $(MAKE) -C $(GODELCLI_DIR)
