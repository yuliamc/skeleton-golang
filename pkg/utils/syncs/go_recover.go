// Synchronized package
package syncs

import (
	"modalrakyat/skeleton-golang/config"
	"modalrakyat/skeleton-golang/pkg/utils/errors"
	"modalrakyat/skeleton-golang/pkg/utils/logs"
	"reflect"
)

// Go Routine Recovery. Execute Go Routine with panic handling to prevent system down due to panic
// and push panic log.
func GoRecover(f func()) {
	go func() {
		defer recoverPanic()
		f()
	}()
}

func recoverPanic() {
	if err := recover(); err != nil {
		fields := logs.Fields{
			"type_str":     "ERR-GORECOVER-PANIC",
			"mode":         config.Config.System.Mode,
			"error_type":   reflect.ValueOf(err).Type().String(),
			"error_string": errors.ToString(err),
			"error_stack":  errors.GetStack(err),
		}

		cl := logs.Log.WithFields(fields)
		logs.Log.Warnf("[START-GORECOVER-PANIC]\n%s\n[END-GORECOVER-PANIC]", errors.GetStack(err), cl)
	}
}
