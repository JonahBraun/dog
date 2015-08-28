package dog

import (
	"fmt"
	"testing"
)

func broadcast(dog *Dog, msg string) {
	dog.Debug(msg + ": debug")
	dog.Info(msg + ": info")
	dog.Warn(msg + ": warn")
	dog.Err(msg + ": err")
	dog.Fatal(msg + ": fatal")

	fmt.Println()
}

func TestStuff(t *testing.T) {
	broadcast(NewDog(DEBUG), "debug level")

	broadcast(NewDog(INFO), "info level")

	broadcast(NewDog(WARN), "warn level")

	broadcast(NewDog(ERR), "err level")

	broadcast(NewDog(FATAL), "fatal level")

	dog := NewDog(DEBUG)
	dog.Debug("debug level")
	theAnswer := 42
	someString := "some string"
	dog.Info("info level with some vars:", someString, theAnswer)
	dog.Warn("warn level")
	dog.Err("error level")
	dog.Fatal("fatal level")
	fmt.Println()

	dog.Fatal = CreateFatal(FgYellow+Reverse, "🐺  ")
	dog.Fatal("Dawg, this custom fatal log format is rockin the CLI!!1")
	fmt.Println()

	dog.Fatal = CreateFatal("", FgRed+Reverse+">"+TR+" ")
	dog.Fatal("A minamilist log format")
	fmt.Println()

	dog.Fatal = CreateFatal("", "")
	dog.Fatal("Keeping the green screen long beards happy.")
	fmt.Println()

	NewDog(DEBUG).Fatal("this fatal log message is followed by a chained call to exit(0)")(0)
}
