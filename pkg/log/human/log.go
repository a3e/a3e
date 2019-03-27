// Package human is a convenience wrapper for logging human-readable messages of various
// levels to the terminal. Much of this code has been shamelessly stolen from
// https://github.com/helm/helm/blob/master/log/log.go.
//
// Thank you Helm team!
package human

import (
	"fmt"
	"io"
	"os"

	"github.com/labstack/gommon/color"
)

// Stdout is the logging destination for normal messages.
var Stdout io.Writer = os.Stdout

// Stderr is the logging destination for error messages.
var Stderr io.Writer = os.Stderr

// IsDebugging toggles whether or not to enable debug output and behavior.
var IsDebugging = false

// Msg passes through the formatter, but otherwise prints exactly as-is.
//
// No prettification.
func Msg(format string, v ...interface{}) {
	fmt.Fprintf(Stdout, appendNewLine(format), v...)
}

// Check checks err, prints an error message using Err(), and returns err
// all if err != nil
// Otherwise, prints the success format string, formatted with successArgs
// using Msg, and returns nil.
func Check(err error, successFmt string, successArgs ...interface{}) error {
	if err != nil {
		Err(err.Error())
		return err
	}
	Msg(successFmt, successArgs...)
	return nil
}

// Die prints an error and then call os.Exit(1).
func Die(format string, v ...interface{}) {
	Err(format, v...)
	if IsDebugging {
		panic(fmt.Sprintf(format, v...))
	}
	os.Exit(1)
}

// CleanExit prints a message and then exits with 0.
func CleanExit(format string, v ...interface{}) {
	Info(format, v...)
	os.Exit(0)
}

// Err prints an error message. It does not cause an exit.
func Err(format string, v ...interface{}) {
	fmt.Fprint(Stderr, color.Red("[ERROR] "))
	fmt.Fprintf(Stderr, appendNewLine(format), v...)
}

// ErrRet does the same thing as Err(format, v...), except returns an
// error with the given format string and arguments
func ErrRet(format string, v ...interface{}) error {
	Err(format, v...)
	return fmt.Errorf(format, v...)
}

// Info prints a green-tinted message.
func Info(format string, v ...interface{}) {
	fmt.Fprint(Stderr, "---> ")
	fmt.Fprintf(Stderr, appendNewLine(format), v...)
}

// Debug prints a cyan-tinted message if IsDebugging is true.
func Debug(msg string, v ...interface{}) {
	if IsDebugging {
		fmt.Fprint(Stderr, color.Cyan("[DEBUG] "))
		Msg(msg, v...)
	}
}

// Warn prints a yellow-tinted warning message.
func Warn(format string, v ...interface{}) {
	fmt.Fprint(Stderr, color.Yellow("[WARN] "))
	Msg(format, v...)
}

func appendNewLine(format string) string {
	return format + "\n"
}
