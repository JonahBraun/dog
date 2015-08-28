## Dog, console logging in go with style.

### Basic Usage

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

### Logging Levels and Customization
Five logging levels are defined: debug, info, warn, err, fatal. A log level is only displayed if it is at or above the level given when creating the dog object. For example, in response to a quiet flag you could only print err and fatal messages:

	dog.NewDog(dog.ERR)

You can customize the appearance of a log level by assigning a new log function with CreateLog(color string, prefix string):

	log.Warn = dog.CreateLog(dog.FgRed, "> ")
	log.Warn("New warning format")

Color ANSI codes are defined to make this convenient. Because CreateLog just concatenates the color and prefix strings, you can do anything you want:

	log.Warn = dog.CreateLog("", dog.FgRed+">"+dog.TR+" ")
	log.Warn("Minimal")

	log.Warn = dog.CreateLog(dog.FgYellow+dog.Reverse, "🐺  ")
	log.Warn("Dawg, this custom fatal log format is rockin the CLI!!1")
