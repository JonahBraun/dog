package main

import "github.com/JonahBraun/dog"

var log *dog.Dog

func main() {
	// in response to some conditional
	log = dog.NewDog(dog.WARN)

	log.Warn = dog.CreateLog(dog.FgRed, "> ")
	log.Warn("Typical customization")

	log.Warn = dog.CreateLog("", dog.FgRed+">"+dog.TR+" ")
	log.Warn("Even more minimal")

	log.Warn = dog.CreateLog(dog.FgYellow+dog.Reverse, "🐺  ")
	log.Warn("Dawg, this custom fatal log format is rockin the CLI!!1")
}
