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
		"--config", "testoutput/build-config-cmd.yaml",
		"-l", "testdata/licenses.yaml",
		"-r", "testdata/license-report.xml"}
	LicenseCmd(args)
	buildConfig := ReadBuildConfig("testoutput/build-config-cmd.yaml")
	assertEqual(t, 2, len(buildConfig.Projects))
	assertEqual(t, 1, len(buildConfig.Projects["org.test.project1"].MavenArtifacts))
	assertEqual(t, 2, len(buildConfig.Projects["org.test.project2"].MavenArtifacts))
	assertEqual(t, 2, len(buildConfig.Projects["org.test.project1"].Licenses))
}
