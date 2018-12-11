package utils

import (
	"fmt"
	"io"
	"myproject/statistical"
	"os"
	"runtime"

	"github.com/s3dteam/go-toolkit/db/mysqldao"
	"github.com/s3dteam/go-toolkit/log/logruslogger"

	"myproject/node"
	"myproject/types"
)

// RegisterStatisticalService register Statistical service
func RegisterStatisticalService(stack *node.Node) {
	var err error
	err = stack.Register(func(ctx *node.ServiceContext) (node.Service, error) {
		return statistical.New(ctx)
	})
	if err != nil {
		panic(fmt.Sprintf("Failed to register the statistical service: %v", err))
	}
}
func defaultLogConfig() logruslogger.Options {
	return logruslogger.Options{
		Level:          "info",
		Depth:          8,
		WithCallerHook: true,
		Formatter:      "text", // only support json and text
		DisableConsole: false,
		Write:          true,
		FileName:       "node",
		Path:           "/tmp/myproject",
		Debug:          false,
	}
}

// SetLogConfig set log config
func SetLogConfig(cfg *logruslogger.Options) {
	if cfg == nil {
		cfg = new(logruslogger.Options)
		*cfg = defaultLogConfig()
		return
	}
	if cfg.Level == "" {
		cfg.Level = "info"
	}
	if cfg.Depth == 0 {
		cfg.Depth = 8
	}
	if cfg.Formatter == "" {
		cfg.Formatter = "text"
	}
	if cfg.FileName == "" {
		cfg.FileName = "node"
	}
	if cfg.Path == "" {
		cfg.Path = "./logs"
	}
}

// SetMySQLConfig set mysql config
func SetMySQLConfig(cfg *mysqldao.MysqlConifg) {
}

// SetRedisConfig set redis config
func SetRedisConfig(cfg *types.RedisOptions) {

}

// SetNodeConfig applies node-related command line flags to the config.
func SetNodeConfig(cfg *node.Config) {
	setHTTP(cfg)
	setWS(cfg)
}

// setHTTP creates the HTTP RPC listener interface string from the set
// command line flags, returning empty if the HTTP endpoint is disabled.
func setHTTP(cfg *node.Config) {
	if cfg.HTTPHost == "" {
		cfg.HTTPHost = "127.0.0.1"
	}

	if cfg.HTTPPort == 0 {
		cfg.HTTPPort = node.DefaultHTTPPort
	}
	if len(cfg.HTTPCors) == 0 {
		cfg.HTTPCors = node.DefaultConfig.HTTPCors
	}

	if len(cfg.HTTPModules) == 0 {
		cfg.HTTPModules = node.DefaultConfig.HTTPModules
	}

	if len(cfg.HTTPVirtualHosts) == 0 {
		cfg.HTTPVirtualHosts = node.DefaultConfig.HTTPVirtualHosts
	}
}

// setWS creates the WebSocket RPC listener interface string from the set
// command line flags, returning empty if the HTTP endpoint is disabled.
func setWS(cfg *node.Config) {
	if cfg.WSHost == "" {
		cfg.WSHost = "127.0.0.1"
	}
	if cfg.WSPort == 0 {
		cfg.WSPort = node.DefaultHTTPPort
	}
	if len(cfg.WSOrigins) == 0 {
		cfg.WSOrigins = node.DefaultConfig.WSOrigins
	}

	if len(cfg.WSModules) == 0 {
		cfg.WSModules = node.DefaultConfig.WSModules
	}
}

// Fatalf formats a message to standard error and exits the program.
// The message is also printed to standard output if standard error
// is redirected to a different file.
func Fatalf(format string, args ...interface{}) {
	w := io.MultiWriter(os.Stdout, os.Stderr)
	if runtime.GOOS == "windows" {
		// The SameFile check below doesn't work on Windows.
		// stdout is unlikely to get redirected though, so just print there.
		w = os.Stdout
	} else {
		outf, _ := os.Stdout.Stat()
		errf, _ := os.Stderr.Stat()
		if outf != nil && errf != nil && os.SameFile(outf, errf) {
			w = os.Stderr
		}
	}
	fmt.Fprintf(w, "Fatal: "+format+"\n", args...)
	os.Exit(1)
}
