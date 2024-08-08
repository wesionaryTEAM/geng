package pkg

import (
	"github.com/charmbracelet/log"
	"os"
)

type GengLogger struct {
	*log.Logger
}

var _logger *GengLogger

func GetLogger() *GengLogger {
	if _logger == nil {
		l := log.New(os.Stderr)
		_logger = &GengLogger{l}
	}
	return _logger
}
