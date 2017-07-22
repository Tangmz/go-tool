package util

import (
	"errors"
	"fmt"
)

func Error(format string, args... interface{}) error {
	return errors.New(fmt.Sprintf(format, args...))
}
