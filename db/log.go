package db

import (
	hlog "github.com/s3dteam/go-toolkit/log"
	"github.com/s3dteam/go-toolkit/log/logruslogger"
)

var (
	Log hlog.Logger
)

// SetModuleLogger set module logger
func SetModuleLogger(name string, options *logruslogger.Options) {
	Log = logruslogger.GetLoggerWithOptions(name, options)
}
