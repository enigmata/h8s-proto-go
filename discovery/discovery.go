package discovery

const (
	defaultDaemonIP         string = "0.0.0.0"
	defaultDaemonPort       string = "19999"
	defaultDaemonListenAddr string = defaultDaemonIP + ":" + defaultDaemonPort
)

func GetDaemonAddress() string {
	return defaultDaemonListenAddr
}
