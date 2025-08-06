package desc

import (
	"fmt"
	"strings"
)

type str string

func (me str) String() string {
	return string(me)
}

func Str(s string) str {
	return str(s)
}

func Err(err error, context fmt.Stringer) error {
	if err == nil {
		return nil
	}
	var message = context.String()
	if !strings.Contains(message, "%w") {
		message = message + " ⁝ %w"
	}
	return fmt.Errorf(message, err)
}
