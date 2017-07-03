package main

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

const execName = "hawkbuild"

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	os.Exit(retCode)
}

func setup() {
	fmt.Print("Setup for testing \n")
	os.Mkdir("testoutput", 0777)
}

func TestMain_Invalid(t *testing.T) {
	if os.Getenv("FORKED_PROCESS") == "1" {
		args := []string{execName, "foo"}
		Main(args)
		return
	}
	cmd := exec.Command(os.Args[0], "-test.run=TestMain_Invalid")
	cmd.Env = append(os.Environ(), "FORKED_PROCESS=1")
	err := cmd.Run()
	fmt.Printf("The error: %q\n", err.(*exec.ExitError))
	if e, ok := err.(*exec.ExitError); ok && !e.Success() {
		return
	}
	t.Fatalf("Invalid command exited with err %v, want exit status 1", err)
}

func TestMain_Help(t *testing.T) {
	args := []string{execName, "--help"}
	Main(args)
	args = []string{execName}
	Main(args)
}

func TestMain_LicenseCmd(t *testing.T) {
	args := []string{execName, "license",
		"--config", "config/testdata/build-config.yaml",
		"-l", "config/testdata/licenses.yaml",
		"-t", "testoutput/licence-cmd-report.xml"}
	Main(args)
}
