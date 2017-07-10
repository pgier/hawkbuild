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

package config

import (
	"testing"

	"github.com/pgier/hawkbuild/test"
)

const (
	testBuildConfigFile        = "examples/build-config.yaml"
	testProjectName            = "my-test-project1"
	apacheLicenseName          = "Apache Software License 2.0"
	expectedProjectVersion     = "1.0"
	licenseReportXMLFile       = "testoutput/license-report.xml"
	expectedArtifact1          = "test-project1-artifact1"
	licenseReportTestInputFile = "examples/license-report.xml"
	BuildConfigOutputFile      = "testoutput/build-config-test.yaml"
)

func TestReadBuildConfigFile(t *testing.T) {
	config := ReadBuildConfigFile(testBuildConfigFile)
	test.AssertEqual(t, 2, len(config.Projects))
	project1 := config.Projects[testProjectName]
	test.AssertTrue(t, (len(project1.Licenses) > 0))
	test.AssertEqual(t, apacheLicenseName, project1.Licenses[0])
	test.AssertEqual(t, expectedProjectVersion, project1.Version)
	test.AssertEqual(t, expectedArtifact1, project1.MavenArtifacts[0].ArtifactID)
}

func TestReadBuildConfigFileFail(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("The code did not panic")
		}
	}()
	ReadBuildConfigFile("bad_filename")
}

func TestGenerateBuildConfig(t *testing.T) {
	outputFile := "testoutput/generated-config.yaml"
	GenerateBasicBuildConfig(outputFile)
	config := ReadBuildConfigFile(outputFile)
	test.AssertEqual(t, 2, len(config.Projects))
	test.AssertEqual(t, "2.1.0.Beta1", config.Projects["my-test-project2"].Version)
	test.AssertEqual(t, "test-project1-artifact1", config.Projects["my-test-project1"].MavenArtifacts[0].ArtifactID)
}
