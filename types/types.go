package types

import (
	"crypto/tls"
)

// MysqlOptions mysql options
type MysqlOptions struct {
	Hostname           string
	Port               string
	User               string
	Password           string
	DBName             string
	TablePrefix        string
	MaxOpenConnections int
	MaxIdleConnections int
	ConnMaxLifetime    int // unit second
	Debug              bool
}

// RedisOptions redis options
type RedisOptions struct {
	Host        string
	Port        string
	Password    string
	IdleTimeout int
	MaxIdle     int
	MaxActive   int
}

// CommonOptions common options
type CommonOptions struct {
	DisableRedis bool
	DisableMySQL bool
}

// SMSOptions sms options
type SMSOptions struct {
	Format  string // default json
	Key     string
	Pwd     string
	SMSSign string `toml:"smsSign"`
}

// EmailOptions email options
type EmailOptions struct {
	Address      string          `toml:"address"`
	Count        int             `toml:"count"`
	CC           string          `toml:"cc"`
	BCC          string          `toml:"bcc"`
	OptTLSConfig []tls.Config    `toml:"optTLSConfig"`
	AuthConfig   EmailAuthConfig `toml:"authConfig"`
	Timeout      int             `toml:"timeout"`
}

// EmailAuthConfig email auth config
type EmailAuthConfig struct {
	Identity string `toml:"identity"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Host     string `toml:"host"`
}
type GinPort struct {
	GinPort string `toml:"ginPort"`
}
