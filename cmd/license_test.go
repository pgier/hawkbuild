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
	licenseReportFile := "testoutput/test-license-report.xml"
	args := []string{"license", "--config",
		"../config/testdata/build-config.yaml",
		"-l", "../config/testdata/licenses.yaml",
		"-r", licenseReportFile}
	RootCmd.SetArgs(args)
	err := RootCmd.Execute()
	if err != nil {
		t.Fatalf("Unable to run basic command: %v", err)
	}
	licenseReport := config.ReadLicenseReportFile(licenseReportFile)
	test.AssertTrue(t, len(licenseReport.Artifacts) > 1)
}

func TestLicenseCmdGenerateConfig(t *testing.T) {
	const productConfigOutputFile = "testoutput/test-reverse-build-config.yaml"
	args := []string{"license", "-g",
		"--config", productConfigOutputFile,
		"-l", "../config/testdata/licenses.yaml",
		"-r", "../config/testdata/license-report.xml"}
	RootCmd.SetArgs(args)
	err := RootCmd.Execute()
	if err != nil {
		t.Fatalf("Unable to run basic command: %v", err)
	}
	productConfig := config.ReadProductConfig(productConfigOutputFile)
	test.AssertEqual(t, 2, len(productConfig.Projects))
	test.AssertEqual(t, 1, len(productConfig.Projects["org.test.project1"].MavenArtifacts))
	test.AssertEqual(t, 2, len(productConfig.Projects["org.test.project2"].MavenArtifacts))
	test.AssertEqual(t, 2, len(productConfig.Projects["org.test.project1"].Licenses))
}
