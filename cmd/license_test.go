/*
Copyright 2017 Red Hat, Inc.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cmd

import (
	"testing"

	"github.com/pgier/hawkbuild/config"
	"github.com/pgier/hawkbuild/test"
)

func TestLicenseCmd(t *testing.T) {
	args := []string{"--config", "../config/testdata/build-config.yaml",
		"-l", "../config/testdata/licenses.yaml",
		"-t", "testoutput/licence-cmd-report.xml"}
	LicenseCmd(args)
}

func TestLicenseCmdReverse(t *testing.T) {
	const buildConfigOutputFile = "testoutput/build-config-cmd.yaml"
	args := []string{"-r",
		"--config", buildConfigOutputFile,
		"-l", "../config/testdata/licenses.yaml",
		"-t", "../config/testdata/license-report.xml"}
	LicenseCmd(args)
	buildConfig := config.ReadBuildConfig(buildConfigOutputFile)
	test.AssertEqual(t, 2, len(buildConfig.Projects))
	test.AssertEqual(t, 1, len(buildConfig.Projects["org.test.project1"].MavenArtifacts))
	test.AssertEqual(t, 2, len(buildConfig.Projects["org.test.project2"].MavenArtifacts))
	test.AssertEqual(t, 2, len(buildConfig.Projects["org.test.project1"].Licenses))
}
