package interact

import (
	"errors"
	"fmt"
	"reflect"
)

var ErrNotANumber = errors.New("not a number")
var ErrNotBoolean = errors.New("not y, n, yes, or no")

type NotAssignableError struct {
	Destination reflect.Type
	Value       reflect.Type
}

func (err NotAssignableError) Error() string {
	return fmt.Sprintf("chosen value (%T) is not assignable to %T", err.Value, err.Destination)
}
