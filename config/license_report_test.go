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

func TestWriteLicenseReportFile(t *testing.T) {
	buildConfig := ReadBuildConfigFile(testBuildConfigFile)
	test.AssertEqual(t, expectedProjectVersion, buildConfig.Projects[testProjectName].Version)
	buildConfigs := []BuildConfig{buildConfig}
	WriteLicenseReportFile(buildConfigs, licenseReportXMLFile)
	licenseReportXML := ReadLicenseReportFile(licenseReportXMLFile)

	found, artifact := findArtifactByArtifactID(licenseReportXML.Artifacts, expectedArtifact1)
	test.AssertTrue(t, found)
	test.AssertEqual(t, "1.0", artifact.Version)
}

func TestGenerateLicenseReport(t *testing.T) {
	config := ReadBuildConfigFile(testBuildConfigFile)
	test.AssertEqual(t, expectedProjectVersion, config.Projects[testProjectName].Version)
	test.AssertEqual(t, expectedArtifact1, config.Projects[testProjectName].MavenArtifacts[0].ArtifactID)
	configs := []BuildConfig{config}
	licenseReport := generateLicenseReport(configs)
	ok, artifact := findArtifactByArtifactID(licenseReport.Artifacts, expectedArtifact1)
	test.AssertTrue(t, ok)
	test.AssertEqual(t, expectedArtifact1, artifact.ArtifactID)
	test.AssertEqual(t, 1, len(artifact.Licenses))
	test.AssertEqual(t, "Apache Software License 2.0", artifact.Licenses[0].Name)
}

func TestReadLicenseReportFile(t *testing.T) {
	licenseSummary := ReadLicenseReportFile(licenseReportTestInputFile)
	test.AssertEqual(t, 3, len(licenseSummary.Artifacts))
	found, artifact := findArtifactByArtifactID(licenseSummary.Artifacts, "project1-artifact1")
	test.AssertTrue(t, found)
	test.AssertEqual(t, expectedProjectVersion, artifact.Version)
	test.AssertEqual(t, 2, len(artifact.Licenses))
	test.AssertEqual(t, "Apache Software License 2.0", artifact.Licenses[0].Name)
}

func TestLicenseReportToBuildConfig(t *testing.T) {
	licenseReport := ReadLicenseReportFile(licenseReportTestInputFile)
	exampleConfig := ReadBuildConfigFile("examples/build-config.yaml")
	configs := []BuildConfig{exampleConfig, DefaultLicenseConfig}
	generatedConfig := licenseReportToBuildConfig(configs, licenseReport)
	test.AssertEqual(t, 2, len(generatedConfig.Projects))
	test.AssertEqual(t, "1.0", generatedConfig.Projects["org.test.project1"].Version)
	test.AssertEqual(t, 1, len(generatedConfig.Projects["org.test.project1"].MavenArtifacts))
}
