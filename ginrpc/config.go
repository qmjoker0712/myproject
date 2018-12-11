package ginrpc

type Config struct {
	GinPort string `toml:"GinPort" json:"GinPort"`
}

//==============================
func DefaultConfig() *Config {
	return &Config{
		GinPort: "8630",
	}
}
