package statistical

import (
	hlog "github.com/s3dteam/go-toolkit/log"
	"github.com/s3dteam/go-toolkit/log/logruslogger"
)

var (
	log hlog.Logger
)

// SetModuleLogger set module logger
func SetModuleLogger(name string, options *logruslogger.Options) {
	log = logruslogger.GetLoggerWithOptions(name, options)
}
