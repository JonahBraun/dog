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
	LevelDebug = iota
	LevelInfo  = iota
	LevelWarn  = iota
	LevelErr   = iota
	LevelFatal = iota
)

// color reference for convenience
const (
	TR         = "\x1b[0m" // terminal reset
	Bright     = "\x1b[1m"
	Dim        = "\x1b[2m"
	Underscore = "\x1b[4m"
	Blink      = "\x1b[5m"
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
	Debug, Info, Warn, Err, Fatal func(v ...interface{}) interface{}
}

func (dog *Dog) Exit(code int) {
	os.Exit(code)
}

// creates a logging function that is ignored
// used for calls below the logging level
func ignore() func(...interface{}) interface{} {
	return func(v ...interface{}) interface{} {
		return v
	}
}

// creates a logging function
func log(color string, icon string) func(...interface{}) interface{} {
	return func(v ...interface{}) interface{} {
		v[0] = fmt.Sprintf("%v%v%v%v", color, icon, v[0], TR)
		fmt.Println(v...)
		return v
	}
}

func NewDog(level int) *Dog {
	switch level {
	case LevelDebug:
		return &Dog{
			Debug: log(FgCyan, "→ "),
			Info:  log(FgGreen, "✏  "),
			Warn:  log(FgYellow, "⚠  "),
			Err:   log(FgRed, "✖  "),
			Fatal: log(FgRed+Bright, "☠  "),
		}
	case LevelInfo:
		return &Dog{
			Debug: ignore(),
			Info:  log(FgGreen, "✏  "),
			Warn:  log(FgYellow, "⚠  "),
			Err:   log(FgRed, "✖  "),
			Fatal: log(FgRed+Bright, "☠  "),
		}
	case LevelWarn:
		return &Dog{
			Debug: ignore(),
			Info:  ignore(),
			Warn:  log(FgYellow, "⚠  "),
			Err:   log(FgRed, "✖  "),
			Fatal: log(FgRed+Bright, "☠  "),
		}
	case LevelErr:
		return &Dog{
			Debug: ignore(),
			Info:  ignore(),
			Warn:  ignore(),
			Err:   log(FgRed, "✖  "),
			Fatal: log(FgRed+Bright, "☠  "),
		}
	case LevelFatal:
		return &Dog{
			Debug: ignore(),
			Info:  ignore(),
			Warn:  ignore(),
			Err:   ignore(),
			Fatal: log(FgRed+Bright, "☠  "),
		}
	}

	return &Dog{}
}
