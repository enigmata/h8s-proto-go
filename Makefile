# overrideable by env vars
CLI_VERSION ?= 0.1.0
DAEMON_VERSION ?= 0.1.0

# internal env vars
GODELCLI_DIR = godelcli
GODELD_DIR = godeld

# simple constants
MAKE := make
CLI_VERSION_ENV_VAR := CLI_VERSION
DAEMON_VERSION_ENV_VAR := DAEMON_VERSION

.PHONY: install

default: install

install:
	${CLI_VERSION_ENV_VAR}=${CLI_VERSION} $(MAKE) -C $(GODELCLI_DIR) install
	${DAEMON_VERSION_ENV_VAR}=${DAEMON_VERSION} $(MAKE) -C $(GODELD_DIR) install
