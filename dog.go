// Console logging with style (prefixed & colored).
//
// Includes terminal color labels for convenience.
// An experiment with unicode symbols and functional programming.
package dog

import (
	"fmt"
	"os"
)

// Logging levels, pass one to CreateDog.
const (
	DEBUG = iota
	INFO  = iota
	WARN  = iota
	ERR   = iota
	FATAL = iota
)

// ANSI color codes, use these as parameters for CreateLog and CreateFatal.
const (
	TR         = "\x1b[0m" // terminal reset
	Bright     = "\x1b[1m"
	Dim        = "\x1b[2m"
	Underscore = "\x1b[4m"
	Blink      = "\x1b[5m" // many terminals treat this like Bright
	Reverse    = "\x1b[7m"
	Hidden     = "\x1b[8m"

	FgBlack   = "\x1b[30m"
	FgRed     = "\x1b[31m"
	FgGreen   = "\x1b[32m"
	FgYellow  = "\x1b[33m"
	FgBlue    = "\x1b[34m"
	FgMagenta = "\x1b[35m"
	FgCyan    = "\x1b[36m"
	FgWhite   = "\x1b[37m"

	BgBlack   = "\x1b[40m"
	BgRed     = "\x1b[41m"
	BgGreen   = "\x1b[42m"
	BgYellow  = "\x1b[43m"
	BgBlue    = "\x1b[44m"
	BgMagenta = "\x1b[45m"
	BgCyan    = "\x1b[46m"
	BgWhite   = "\x1b[47m"
)

// Dog functions are created at runtime with NewDog. Use CreateLog and CreateFatal customize.
// Fatal returns os.Exit(int) to facilitate chaining an exit call.
type Dog struct {
	Debug, Info, Warn, Err func(...interface{})

	Fatal func(...interface{}) func(int)
}

// Creates a logging function that is ignored.
// Used for calls below the configured logging level.
func Ignore() func(...interface{}) {
	return func(v ...interface{}) {}
}

// Creates a logging function. Color can be from the color constants. Prefix can be an
// appropriate icon followed by a space, or two spaces for doublewide unicode symbols.
// The two parameters are merely concatanted in front of the log message, so they can be any string.
//
// To facilitate functional programming, it would be nice to return the passed variables
// however, this is nearly useless until generics are added to go.
func CreateLog(color string, prefix string) func(...interface{}) {
	return func(v ...interface{}) {
		v[0] = fmt.Sprintf("%v%v%v%v", color, prefix, v[0], TR)
		fmt.Println(v...)
	}
}

// Like CreateLog but returns os.Exit so that an exit call can easily be chained
func CreateFatal(color string, prefix string) func(...interface{}) func(int) {
	return func(v ...interface{}) func(int) {
		v[0] = fmt.Sprintf("%v%v%v%v", color, prefix, v[0], TR)
		fmt.Println(v...)
		return os.Exit
	}
}

// Creates a *Dog configured to print at and above the passed logging level.
func NewDog(level int) *Dog {
	switch level {
	case DEBUG:
		return &Dog{
			Debug: CreateLog(FgCyan, "→ "),
			Info:  CreateLog(FgGreen, "✏  "),
			Warn:  CreateLog(FgYellow, "⚠  "),
			Err:   CreateLog(FgRed, "✖  "),
			Fatal: CreateFatal(FgRed+Bright, "☠  "),
		}
	case INFO:
		return &Dog{
			Debug: Ignore(),
			Info:  CreateLog(FgGreen, "✏  "),
			Warn:  CreateLog(FgYellow, "⚠  "),
			Err:   CreateLog(FgRed, "✖  "),
			Fatal: CreateFatal(FgRed+Bright, "☠  "),
		}
	case WARN:
		return &Dog{
			Debug: Ignore(),
			Info:  Ignore(),
			Warn:  CreateLog(FgYellow, "⚠  "),
			Err:   CreateLog(FgRed, "✖  "),
			Fatal: CreateFatal(FgRed+Bright, "☠  "),
		}
	case ERR:
		return &Dog{
			Debug: Ignore(),
			Info:  Ignore(),
			Warn:  Ignore(),
			Err:   CreateLog(FgRed, "✖  "),
			Fatal: CreateFatal(FgRed+Bright, "☠  "),
		}
	case FATAL:
		return &Dog{
			Debug: Ignore(),
			Info:  Ignore(),
			Warn:  Ignore(),
			Err:   Ignore(),
			Fatal: CreateFatal(FgRed+Bright, "☠  "),
		}
	}

	return &Dog{}
}
