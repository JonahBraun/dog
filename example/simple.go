package main

import "github.com/JonahBraun/dog"

var log = dog.NewDog(dog.DEBUG)

func main() {
	number := 8
	msg := "error message form some library"

	// any variables can be appended to any of the log calls
	log.Debug("Might be needed for advanced debugging")
	log.Info("Your number is", number)
	log.Warn("Possibly an issue")
	log.Err("This should't be happening…", msg)

	// Fatal returns a call to os.Exit(int) so that you can easily exit
	log.Fatal("We need to exit")(1)
}
