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
	testLicensesFile           = "testdata/licenses.yaml"
	testBuildConfigFile        = "testdata/build-config.yaml"
	testProjectName            = "my-test-project1"
	apacheLicenseShortName     = "ASL 2.0"
	apacheLicenseFullName      = "Apache Software License 2.0"
	expectedProjectVersion     = "1.0"
	licenseReportXMLFile       = "testoutput/license-report.xml"
	expectedArtifact1          = "test-project1-artifact1"
	licenseReportTestInputFile = "testdata/license-report.xml"
	buildConfigOutputFile      = "testoutput/build-config-test.yaml"
)

func TestReadLicenses(t *testing.T) {
	licenses := ReadLicenseConfig(testLicensesFile)
	const minExpectedLicenseCount = 2
	test.AssertTrue(t, minExpectedLicenseCount < len(licenses))
	test.AssertEqual(t, apacheLicenseShortName, licenses[apacheLicenseFullName].ShortName)
}

func TestCreateLicenseShortNameMap(t *testing.T) {
	licenses := ReadLicenseConfig(testLicensesFile)
	licShortNameMap := createLicenseShortNameMap(licenses)
	test.AssertEqual(t, apacheLicenseFullName, licShortNameMap[apacheLicenseShortName].Name)
}

func TestReadBuildConfig(t *testing.T) {
	config := ReadBuildConfig(testBuildConfigFile)
	const expectedPackageCount = 2
	test.AssertEqual(t, expectedPackageCount, len(config.Projects))
	const expectedLicenseName = apacheLicenseShortName
	project1 := config.Projects[testProjectName]
	test.AssertTrue(t, (len(project1.Licenses) > 0))
	test.AssertEqual(t, expectedLicenseName, project1.Licenses[0])
	test.AssertEqual(t, expectedProjectVersion, project1.Version)
	test.AssertEqual(t, expectedArtifact1, project1.MavenArtifacts[0].ArtifactID)
}

func TestReadMissingFile(t *testing.T) {
	defer func() {
		if recover() == nil {
			t.Errorf("The code did not panic")
		}
	}()
	ReadBuildConfig("bad_filename")
}

func TestWriteProjectLicenseSummaryXML(t *testing.T) {
	licenses := ReadLicenseConfig(testLicensesFile)
	buildConfig := ReadBuildConfig(testBuildConfigFile)
	test.AssertEqual(t, expectedProjectVersion, buildConfig.Projects[testProjectName].Version)
	WriteLicenseReportFile(licenses, buildConfig, licenseReportXMLFile)
	licenseReportXML := ReadLicenseReportFile(licenseReportXMLFile)

	found, artifact := findArtifactByArtifactID(licenseReportXML.Artifacts, expectedArtifact1)
	test.AssertTrue(t, found)
	test.AssertEqual(t, "1.0", artifact.Version)
}

func TestConvertBuildConfigToLicenseSummary(t *testing.T) {
	licenses := ReadLicenseConfig(testLicensesFile)
	config := ReadBuildConfig(testBuildConfigFile)
	test.AssertEqual(t, expectedProjectVersion, config.Projects[testProjectName].Version)
	test.AssertEqual(t, expectedArtifact1, config.Projects[testProjectName].MavenArtifacts[0].ArtifactID)
	licenseSummary := ConvertBuildConfigToLicenseSummary(licenses, config)
	ok, artifact := findArtifactByArtifactID(licenseSummary.Artifacts, expectedArtifact1)
	test.AssertTrue(t, ok)
	test.AssertEqual(t, expectedArtifact1, artifact.ArtifactID)
}

func TestReadLicenseReportFile(t *testing.T) {
	licenseSummary := ReadLicenseReportFile(licenseReportTestInputFile)
	test.AssertEqual(t, 3, len(licenseSummary.Artifacts))
	found, artifact := findArtifactByArtifactID(licenseSummary.Artifacts, "project1-artifact1")
	test.AssertTrue(t, found)
	test.AssertEqual(t, expectedProjectVersion, artifact.Version)
	test.AssertEqual(t, "Apache Software License 2.0", artifact.Licenses[0].Name)
	test.AssertEqual(t, 2, len(artifact.Licenses))
}

func TestLicenseReportToBuildConfig(t *testing.T) {
	licenseReport := ReadLicenseReportFile(licenseReportTestInputFile)
	licenses := ReadLicenseConfig(testLicensesFile)
	config := LicenseReportToBuildConfig(licenses, licenseReport)
	test.AssertEqual(t, 2, len(config.Projects))
	test.AssertEqual(t, expectedProjectVersion, config.Projects["org.test.project1"].Version)
	test.AssertEqual(t, 1, len(config.Projects["org.test.project1"].MavenArtifacts))
}
