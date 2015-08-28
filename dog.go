// Dog: Logging with style.
//
// Includes terminal color labels for convenience
// An experiment with CLI emoji

package dog

import (
	"fmt"
	"os"
)

// logging levels
const (
	DEBUG = iota
	INFO  = iota
	WARN  = iota
	ERR   = iota
	FATAL = iota
)

// color reference for convenience
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

type Dog struct {
	Debug, Info, Warn, Err func(...interface{})

	// Fatal returns a os.Exit so that an exit call can easily be chained
	Fatal func(...interface{}) func(int)
}

// creates a logging function that is ignored
// used for calls below the logging level
func Ignore() func(...interface{}) {
	return func(v ...interface{}) {}
}

// creates a logging function
// to facilitate functional programming, it would be nice to return the passed variables
// however, this is nearly useless until generics are added to go
func CreateLog(color string, prefix string) func(...interface{}) {
	return func(v ...interface{}) {
		v[0] = fmt.Sprintf("%v%v%v%v", color, prefix, v[0], TR)
		fmt.Println(v...)
	}
}

// Fatal returns a os.Exit so that an exit call can easily be chained
func CreateFatal(color string, prefix string) func(...interface{}) func(int) {
	return func(v ...interface{}) func(int) {
		v[0] = fmt.Sprintf("%v%v%v%v", color, prefix, v[0], TR)
		fmt.Println(v...)
		return os.Exit
	}
}

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
