package dog

import (
	"fmt"
	"testing"
)

func broadcast(dog *Dog, msg string) {
	dog.Debug(msg)
	dog.Info(msg)
	dog.Warn(msg)
	dog.Err(msg)
	dog.Fatal(msg)

	fmt.Println()
}

func TestNormal(t *testing.T) {
	broadcast(NewDog(LevelDebug), "debug level")

	broadcast(NewDog(LevelInfo), "info level")

	broadcast(NewDog(LevelWarn), "warn level")

	broadcast(NewDog(LevelErr), "err level")

	broadcast(NewDog(LevelFatal), "fatal level")

	NewDog(LevelDebug).Exit(0)
}
