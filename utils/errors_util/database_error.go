package errors_util

import "fmt"

type DBExecuteError struct {
	Msg string
}

func (e *DBExecuteError) Error() string {
	return fmt.Sprintf("DBExecuteError: %s", e.Msg)
}
