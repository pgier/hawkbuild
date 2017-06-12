package cmd

import "testing"

func TestLicenseCmd(t *testing.T) {
	args := []string{"--config", "testdata/build-config.yaml",
		"-l", "testdata/licenses.yaml",
		"-r", "testoutput/licence-cmd-report.xml"}
	LicenseCmd(args)
}

func TestLicenseCmdReverse(t *testing.T) {
	args := []string{"-v",
		"--config", "testoutput/build-config.yaml",
		"-l", "testdata/licenses.yaml",
		"-r", "testdata/license-report.xml"}
	LicenseCmd(args)
}
