package cmd

import (
	"testing"

	"github.com/pgier/hawkbuild/config"
	"github.com/pgier/hawkbuild/util"
)

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
	buildConfig := config.ReadBuildConfig("testoutput/build-config-cmd.yaml")
	util.AssertEqual(t, 2, len(buildConfig.Projects))
	util.AssertEqual(t, 1, len(buildConfig.Projects["org.test.project1"].MavenArtifacts))
	util.AssertEqual(t, 2, len(buildConfig.Projects["org.test.project2"].MavenArtifacts))
	util.AssertEqual(t, 2, len(buildConfig.Projects["org.test.project1"].Licenses))
}
