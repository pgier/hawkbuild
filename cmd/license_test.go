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
	"os/exec"
	"testing"

	"github.com/pgier/hawkbuild/config"
	"github.com/pgier/hawkbuild/test"
)

func TestLicenseCmd(t *testing.T) {
	licenseReportFile := "testoutput/test-license-report.xml"
	cmd := exec.Command("hawkbuild", "license", "--config",
		"../config/testdata/build-config.yaml",
		"-l", "../config/testdata/licenses.yaml",
		"-t", licenseReportFile)
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Unable to run basic command: %v", err)
	}
	licenseReport := config.ReadLicenseReportFile(licenseReportFile)
	test.AssertTrue(t, len(licenseReport.Artifacts) > 1)
}

func TestLicenseCmdReverse(t *testing.T) {
	const buildConfigOutputFile = "testoutput/test-reverse-build-config.yaml"
	cmd := exec.Command("hawkbuild", "license", "-r",
		"--config", buildConfigOutputFile,
		"-l", "../config/testdata/licenses.yaml",
		"-t", "../config/testdata/license-report.xml")
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Unable to run basic command: %v", err)
	}
	buildConfig := config.ReadBuildConfig(buildConfigOutputFile)
	test.AssertEqual(t, 2, len(buildConfig.Projects))
	test.AssertEqual(t, 1, len(buildConfig.Projects["org.test.project1"].MavenArtifacts))
	test.AssertEqual(t, 2, len(buildConfig.Projects["org.test.project2"].MavenArtifacts))
	test.AssertEqual(t, 2, len(buildConfig.Projects["org.test.project1"].Licenses))
}
