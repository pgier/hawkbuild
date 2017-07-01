package config

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	os.Exit(retCode)
}

func setup() {
	fmt.Print("Setup for testing \n")
	os.Mkdir("testoutput", 0777)
}
