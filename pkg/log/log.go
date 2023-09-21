package log

import (
	"fmt"
	"os"
)

// Info logs a message to stdout.
func Info(format string, a ...interface{}) {
	fmt.Fprintf(os.Stdout, format+"\n", a...)
}

// Warn logs a warning message to stdout.
func Warn(format string, a ...interface{}) {
	fmt.Fprintf(os.Stdout, "WARNING: "+format+"\n", a...)
}

// Error logs an error message to stderr.
func Error(format string, a ...interface{}) {
	fmt.Fprintf(os.Stderr, "ERROR: "+format+"\n", a...)
}
