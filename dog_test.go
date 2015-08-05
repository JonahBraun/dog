package dog

import (
	"testing"
)

func TestNormal(t *testing.T) {
	dog := NewDog()
	dog.Debug("should not show this debug")
	dog.Info("info")
	dog.Warn("warn")
	dog.Err("error")

	dog = NewDog(ConfigLevel("debug"))
	dog.Debug("debug")
	dog.Info("info")
	dog.Warn("warn")
	dog.Err("error")

	dog = NewDog(ConfigLevel("err"))
	dog.Err("just this err")
}

func TestFatal(t *testing.T) {
	dog := NewDog()
	dog.Fatal("fatal")
}
