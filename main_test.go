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

package main

import (
	"os"
	"os/exec"
	"testing"

	"github.com/pgier/hawkbuild/cmd"
	"github.com/pgier/hawkbuild/config"
	"github.com/pgier/hawkbuild/test"
)

func TestMain(m *testing.M) {
	setup()
	retCode := m.Run()
	os.Exit(retCode)
}

func setup() {
	os.Mkdir("testoutput", 0777)
}

const execName = "hawkbuild"

func TestHelp(t *testing.T) {
	args := []string{"--help"}
	cmd.RootCmd.SetArgs(args)
	main()
}

func TestDefaultCmd(t *testing.T) {
	cmd := exec.Command(execName)
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Unable to run command: %v", err)
	}
}

func TestLicenseCmdExec(t *testing.T) {
	licenseReportFile := "testoutput/test-license-report.xml"
	cmd := exec.Command(execName, "license",
		"--config", "config/testdata/build-config.yaml",
		"-l", "config/testdata/licenses.yaml",
		"-r", licenseReportFile)
	err := cmd.Run()
	if err != nil {
		t.Fatalf("Unable to run command: %v", err)
	}
	licenseReport := config.ReadLicenseReportFile(licenseReportFile)
	test.AssertTrue(t, len(licenseReport.Artifacts) > 1)
}
