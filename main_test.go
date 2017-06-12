package main

import "testing"

func TestMain_Invalid(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("The code did not panic")
		}
	}()
	args := []string{"hawkbuild", "foo"}
	Main(args)
}

func TestMain_Help(t *testing.T) {
	args := []string{"hawkbuild", "--help"}
	Main(args)
	args = []string{"hawkbuild"}
	Main(args)
}

func TestMain_LicenseCmd(t *testing.T) {
	args := []string{"hawkbuild", "license",
		"--config", "cmd/testdata/build-config.yaml",
		"-l", "cmd/testdata/licenses.yaml",
		"-r", "cmd/testoutput/licence-cmd-report.xml"}
	Main(args)
}
