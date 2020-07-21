package exception

import (
	"fmt"
)

/***************************************
 * Example:
 *   TryCatch(func() {
 *       // do sth.
 *       Throw("msg")
 *   }, func(e interface{}) {
 *       // Exception caught
 *   })
 **************************************/

func Try(try func()) interface{} {
	var ret interface{} = nil
	TryCatch(try, func(e interface{}) {
		ret = e
	})
	return ret
}

func TryAndCatchError(try func()) error {
	e := Try(try)
	return ToError(e)
}

func TryCatch(try func(), catch func(e interface{})) {
	defer func() {
		if e := recover(); e != nil {
			catch(e)
		}
	}()
	try()
}

func Throw(e interface{}) {
	panic(e)
}

func ThrowErrorf(msg string, args ...interface{}) {
	err := fmt.Errorf(msg, args...)
	Throw(err)
}

func ThrowIfErrorNotNil(err error) {
	if err != nil {
		Throw(err)
	}
}

func ToError(e interface{}) error {
	if e == nil {
		return nil
	}
	switch e.(type) {
	case error:
		return e.(error)
	}
	return fmt.Errorf("%s", e)
}
