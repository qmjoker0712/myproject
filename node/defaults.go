package node

// Default HTTP and WS config
const (
	DefaultHTTPHost = "localhost" // Default host interface for the HTTP RPC server
	DefaultHTTPPort = 8545        // Default TCP port for the HTTP RPC server
	DefaultWSHost   = "localhost" // Default host interface for the websocket RPC server
	DefaultWSPort   = 8546        // Default TCP port for the websocket RPC server
)

// DefaultConfig contains reasonable default settings.
var DefaultConfig = Config{
	HTTPPort:         DefaultHTTPPort,
	HTTPModules:      []string{""},
	HTTPVirtualHosts: []string{"localhost"},
	WSPort:           DefaultWSPort,
	WSModules:        []string{""},
}
