package manager

import (
	"errors"
	"fmt"
	. "pure-init/lib/log"
)

func Error(msg string, args ...interface{}) error {
	errMsg := fmt.Sprintf(msg, args...)
	Log.Error(errMsg)
	return errors.New(errMsg)
}
