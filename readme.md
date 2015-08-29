## Dog, console logging in go with style.

### Basic Usage
Import this library and create the log object:
```go
import "github.com/JonahBraun/dog"
var log = dog.NewDog(dog.DEBUG)
```

Then somewhere in a function:
```go
log.Warn("some warning", someVar)
```

### Full Example
```go
package main

import "github.com/JonahBraun/dog"

var log = dog.NewDog(dog.DEBUG)

func main() {
	number := 8
	msg := "error message from some library"

	log.Debug("needed for advanced debugging")

	// variable(s) can be appended to any of the log calls
	log.Info("Your number is", number)

	log.Warn("Possibly an issue")

	log.Err("This should't be happening…", msg)

	// Fatal returns os.Exit(int) so that you can easily exit
	log.Fatal("Fatal, must exit")(1)
}
```

<img width="517" alt="dog_simple" src="https://cloud.githubusercontent.com/assets/611339/9560595/a9af8b2e-4dd1-11e5-8d41-17e51a3b9d8f.png">
Note that the final line is `go run` indicating a non 0 exit status and not something dog is printing.

### Logging Levels and Customization
Five logging levels are defined: debug, info, warn, err, fatal. A log level is only displayed if it is at or above the level given when creating the dog object. For example, in response to a quiet flag you could only print err and fatal messages:
```go
	dog.NewDog(dog.ERR)
```

You can customize the appearance of a log level by assigning a new log function with CreateLog(color string, prefix string):
```go
	log.Warn = dog.CreateLog(dog.FgRed, "> ")
	log.Warn("Typical customization")
```

Color ANSI codes are defined to make this convenient. Because CreateLog just concatenates the color and prefix strings, you can do anything you want:
```go
	log.Warn = dog.CreateLog("", dog.FgRed+">"+dog.TR+" ")
	log.Warn("Even more minimal")

	log.Warn = dog.CreateLog(dog.FgYellow+dog.Reverse, "🐺  ")
	log.Warn("Dawg, this custom fatal log format is rockin the CLI!!1")
```
<img width="464" alt="dog_custom" src="https://cloud.githubusercontent.com/assets/611339/9560594/a99a5466-4dd1-11e5-953d-dc273dd9e23a.png">

### Why?
While many applications don't require five levels of color differentiated and prefixed console logging, some however do. This library was originally written for [Wago, a file watcher / build tool](https://github.com/JonahBraun/wago), which required log messages to be easily distinguishable from the application output being monitored.
