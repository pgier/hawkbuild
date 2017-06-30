package config

import (
	"testing"

	"github.com/pgier/hawkbuild/util"
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
	licenses := ReadLicenses(testLicensesFile)
	const expectedLicenseCount = 2
	util.AssertEqual(t, expectedLicenseCount, len(licenses))
	util.AssertEqual(t, apacheLicenseShortName, licenses[apacheLicenseFullName].ShortName)

}

func TestCreateLicenseShortNameMap(t *testing.T) {
	licenses := ReadLicenses(testLicensesFile)
	licShortNameMap := createLicenseShortNameMap(licenses)
	util.AssertEqual(t, apacheLicenseFullName, licShortNameMap[apacheLicenseShortName].Name)
}

func TestReadBuildConfig(t *testing.T) {
	config := ReadBuildConfig(testBuildConfigFile)
	const expectedPackageCount = 2
	util.AssertEqual(t, expectedPackageCount, len(config.Projects))
	const expectedLicenseName = apacheLicenseShortName
	project1 := config.Projects[testProjectName]
	util.AssertTrue(t, (len(project1.Licenses) > 0))
	util.AssertEqual(t, expectedLicenseName, project1.Licenses[0])
	util.AssertEqual(t, expectedProjectVersion, project1.Version)
	util.AssertEqual(t, expectedArtifact1, project1.MavenArtifacts[0].ArtifactID)
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
	licenses := ReadLicenses(testLicensesFile)
	buildConfig := ReadBuildConfig(testBuildConfigFile)
	util.AssertEqual(t, expectedProjectVersion, buildConfig.Projects[testProjectName].Version)
	WriteLicenseReportFile(licenses, buildConfig, licenseReportXMLFile)
	licenseReportXML := ReadLicenseReportFile(licenseReportXMLFile)

	found, artifact := findArtifactByArtifactID(licenseReportXML.Artifacts, expectedArtifact1)
	util.AssertTrue(t, found)
	util.AssertEqual(t, "1.0", artifact.Version)
}

func TestConvertBuildConfigToLicenseSummary(t *testing.T) {
	licenses := ReadLicenses(testLicensesFile)
	config := ReadBuildConfig(testBuildConfigFile)
	util.AssertEqual(t, expectedProjectVersion, config.Projects[testProjectName].Version)
	util.AssertEqual(t, expectedArtifact1, config.Projects[testProjectName].MavenArtifacts[0].ArtifactID)
	licenseSummary := ConvertBuildConfigToLicenseSummary(licenses, config)
	ok, artifact := findArtifactByArtifactID(licenseSummary.Artifacts, expectedArtifact1)
	util.AssertTrue(t, ok)
	util.AssertEqual(t, expectedArtifact1, artifact.ArtifactID)
}

func TestReadLicenseReportFile(t *testing.T) {
	licenseSummary := ReadLicenseReportFile(licenseReportTestInputFile)
	util.AssertEqual(t, 3, len(licenseSummary.Artifacts))
	found, artifact := findArtifactByArtifactID(licenseSummary.Artifacts, "project1-artifact1")
	util.AssertTrue(t, found)
	util.AssertEqual(t, expectedProjectVersion, artifact.Version)
	util.AssertEqual(t, "Apache Software License 2.0", artifact.Licenses[0].Name)
	util.AssertEqual(t, 2, len(artifact.Licenses))
}

func TestLicenseReportToBuildConfig(t *testing.T) {
	licenseReport := ReadLicenseReportFile(licenseReportTestInputFile)
	licenses := ReadLicenses(testLicensesFile)
	config := LicenseReportToBuildConfig(licenses, licenseReport)
	util.AssertEqual(t, 2, len(config.Projects))
	util.AssertEqual(t, expectedProjectVersion, config.Projects["org.test.project1"].Version)
	util.AssertEqual(t, 1, len(config.Projects["org.test.project1"].MavenArtifacts))
}
