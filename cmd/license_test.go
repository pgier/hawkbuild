package cmd

import (
	"testing"

	"github.com/pgier/hawkbuild/config"
	"github.com/pgier/hawkbuild/util"
)

func TestLicenseCmd(t *testing.T) {
	args := []string{"--config", "../config/testdata/build-config.yaml",
		"-l", "../config/testdata/licenses.yaml",
		"-r", "testoutput/licence-cmd-report.xml"}
	LicenseCmd(args)
}

func TestLicenseCmdReverse(t *testing.T) {
	const buildConfigOutputFile = "testoutput/build-config-cmd.yaml"
	args := []string{"-v",
		"--config", buildConfigOutputFile,
		"-l", "../config/testdata/licenses.yaml",
		"-r", "../config/testdata/license-report.xml"}
	LicenseCmd(args)
	buildConfig := config.ReadBuildConfig(buildConfigOutputFile)
	util.AssertEqual(t, 2, len(buildConfig.Projects))
	util.AssertEqual(t, 1, len(buildConfig.Projects["org.test.project1"].MavenArtifacts))
	util.AssertEqual(t, 2, len(buildConfig.Projects["org.test.project2"].MavenArtifacts))
	util.AssertEqual(t, 2, len(buildConfig.Projects["org.test.project1"].Licenses))
}
