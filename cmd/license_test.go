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

func clearLicenseCmdFlags() {
	clearRootCmdFlags()
}

func TestLicenseCmd(t *testing.T) {
	clearLicenseCmdFlags()
}

func TestLicenseReportCmd(t *testing.T) {
	clearLicenseCmdFlags()
	licenseReportFile := "testoutput/test-license-report.xml"
	args := []string{"license", "report", "--config",
		"../config/examples/build-config.yaml",
		"-o", licenseReportFile}
	RootCmd.SetArgs(args)
	err := RootCmd.Execute()
	if err != nil {
		t.Fatalf("Unable to run basic command: %v", err)
	}
	licenseReport := config.ReadLicenseReportFile(licenseReportFile)
	test.AssertTrue(t, len(licenseReport.Artifacts) > 1)
}
