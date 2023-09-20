package errors

import "fmt"

type Error struct {
	messageTemplate string
	messageArgs     []string
}

func (e *Error) Error() string {
	args := make([]interface{}, len(e.messageArgs))
	for i, v := range e.messageArgs {
		args[i] = v
	}
	return fmt.Sprintf(e.messageTemplate, args...)
}
