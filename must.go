package must

import (
	"fmt"
	"strings"
)

func Get[V any](v V, e error) V {
	if e != nil {
		panic(e)
	}
	return v
}

func Have[V any](v V, found bool) V {
	if !found {
		panic("does not have")
	}
	return v
}

func Desc(err error, context interface{ String() string }) error {
	if err == nil {
		return nil
	}
	var message = context.String()
	if !strings.Contains(message, "%w") {
		message = message + " ⁝ %w"
	}
	return fmt.Errorf(message, err)
}
