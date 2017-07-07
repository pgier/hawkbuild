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
	"github.com/pgier/hawkbuild/util"
)

func clearGenerateCmdFlags() {
	clearRootCmdFlags()
}

func TestGenerate(t *testing.T) {
	clearGenerateCmdFlags()
	generatedConfigFile := "testoutput/hawkbuild.yaml"
	args := []string{"generate", "-o", generatedConfigFile}
	RootCmd.SetArgs(args)
	err := RootCmd.Execute()
	if err != nil {
		t.Fatalf("Unable to run command: %v", err)
	}
	buildConfig := config.ReadBuildConfigFile(generatedConfigFile)
	test.AssertTrue(t, len(buildConfig.Projects) == 2)
}

func TestGenerateFromLicenseFile(t *testing.T) {
	clearGenerateCmdFlags()
	generatedConfigFile := "testoutput/hawkbuild-from-license.yaml"
	licenseReportFile := "../config/examples/license-report.xml"
	existingConfigFile := "../config/examples/build-config.yaml"
	args := []string{"generate", "--config", existingConfigFile,
		"-l", licenseReportFile, "-o", generatedConfigFile}
	RootCmd.SetArgs(args)
	err := RootCmd.Execute()
	if err != nil {
		t.Fatalf("Unable to run command: %v", err)
	}
	test.AssertTrue(t, util.IsFile(generatedConfigFile))
	buildConfig := config.ReadBuildConfigFile(generatedConfigFile)
	test.AssertTrue(t, len(buildConfig.Projects) == 2)
}
