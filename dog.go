// Dog: Logging with style.
//
// Includes terminal color labels
// An experiment with CLI emoji

package dog

import (
	"fmt"
	"os"
)

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

func Talk(v ...interface{}) {
	Log(FgCyan, "→ ", v...)
}

func Note(v ...interface{}) {
	Log(FgGreen, "✏  ", v...)
}

func Warn(v ...interface{}) {
	Log(FgYellow, "⚠  ", v...)
}

func Err(v ...interface{}) {
	Log(FgRed, "✖  ", v...)
}

func Fatal(v ...interface{}) {
	Log(FgRed+Blink, "☠  "+FgRed, v...)
	os.Exit(1)
}

func Log(color string, icon string, v ...interface{}) {
	v[0] = fmt.Sprintf("%v%v%v%v", color, icon, v[0], TR)

	fmt.Println(v...)
}

// TODO write a simple but useful func to inspect an object
func Inspect(v ...interface{}) {
	//spew.Dump(v)
	fmt.Println(v...)
}

type Dog struct {
	Debug, Info, Warn, Err, Fatal func(v ...interface{}) interface{}
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
func ConfigLevel(level string) func(Dog) Dog {
	return func(dog Dog) Dog {
		switch level {
		case "debug":
			dog.Debug = log(FgCyan, "→ ")
		case "info":
			// this is the default
		case "warn":
			dog.Info = ignore()
		case "err":
			dog.Info = ignore()
			dog.Warn = ignore()
		case "fatal":
			dog.Info = ignore()
			dog.Warn = ignore()
			dog.Err = ignore()
		}
		return dog
	}
}

func NewDog(config ...func(Dog) Dog) Dog {

	// defaults
	dog := Dog{
		Debug: ignore(),
		Info:  log(FgGreen, "✏  "),
		Warn:  log(FgYellow, "⚠  "),
		Err:   log(FgRed, "✖  "),
		Fatal: func() func(...interface{}) interface{} {
			log := log(FgRed+Bright, "☠  ")
			return func(v ...interface{}) interface{} {
				log(v...)
				os.Exit(1)
				return v // this is never executed, but required per the function signature
			}
		}(),
	}

	// configure
	for _, c := range config {
		dog = c(dog)
	}

	return dog
}
