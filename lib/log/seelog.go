package log

import (
	"fmt"
	"github.com/cihub/seelog"
)

const (
	LOG_CONFIG_FILE   = "etc/log.xml"
	LOG_CONFIG_FORMAT = "etc/log-%s.xml"
)

var Log seelog.LoggerInterface

func InitLog() {
	var err error
	Log, err = seelog.LoggerFromConfigAsString(loggerConfig)

	if err != nil {
		seelog.Critical("err parsing log config", err)
		return
	}
}

func InitLogWithConfig(conf ...string) error {
	file := LOG_CONFIG_FILE
	if len(conf) > 0 {
		file = fmt.Sprintf(LOG_CONFIG_FORMAT, conf[0])
	}
	logger, err := seelog.LoggerFromConfigAsFile(file)
	if err != nil {
		seelog.Criticalf("parsing log config %s failed: %s", file, err)
		return err
	}
	seelog.ReplaceLogger(logger)
	seelog.Infof("init log from %s success", file)
	return nil
}
