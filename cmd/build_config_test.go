package cmd

import "testing"

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
	assertEqual(t, expectedLicenseCount, len(licenses))
	assertEqual(t, apacheLicenseShortName, licenses[apacheLicenseFullName].ShortName)
}

func TestCreateLicenseShortNameMap(t *testing.T) {
	licenses := ReadLicenses(testLicensesFile)
	licShortNameMap := createLicenseShortNameMap(licenses)
	assertEqual(t, apacheLicenseFullName, licShortNameMap[apacheLicenseShortName].Name)
}

func TestReadBuildConfig(t *testing.T) {
	config := ReadBuildConfig(testBuildConfigFile)
	const expectedPackageCount = 2
	assertEqual(t, expectedPackageCount, len(config.Projects))
	const expectedLicenseName = apacheLicenseShortName
	project1 := config.Projects[testProjectName]
	assertTrue(t, (len(project1.Licenses) > 0))
	assertEqual(t, expectedLicenseName, project1.Licenses[0])
	assertEqual(t, expectedProjectVersion, project1.Version)
	assertEqual(t, expectedArtifact1, project1.MavenArtifacts[0].ArtifactID)
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
	assertEqual(t, expectedProjectVersion, buildConfig.Projects[testProjectName].Version)
	WriteLicenseReportFile(licenses, buildConfig, licenseReportXMLFile)
	licenseReportXML := ReadLicenseReportFile(licenseReportXMLFile)

	found, artifact := findArtifactByArtifactID(licenseReportXML.Artifacts, expectedArtifact1)
	assertTrue(t, found)
	assertEqual(t, "1.0", artifact.Version)
}

func TestConvertBuildConfigToLicenseSummary(t *testing.T) {
	licenses := ReadLicenses(testLicensesFile)
	config := ReadBuildConfig(testBuildConfigFile)
	assertEqual(t, expectedProjectVersion, config.Projects[testProjectName].Version)
	assertEqual(t, expectedArtifact1, config.Projects[testProjectName].MavenArtifacts[0].ArtifactID)
	licenseSummary := ConvertBuildConfigToLicenseSummary(licenses, config)
	ok, artifact := findArtifactByArtifactID(licenseSummary.Artifacts, expectedArtifact1)
	assertTrue(t, ok)
	assertEqual(t, expectedArtifact1, artifact.ArtifactID)
}

func TestReadLicenseReportFile(t *testing.T) {
	licenseSummary := ReadLicenseReportFile(licenseReportTestInputFile)
	assertEqual(t, 3, len(licenseSummary.Artifacts))
	found, artifact := findArtifactByArtifactID(licenseSummary.Artifacts, "project1-artifact1")
	assertTrue(t, found)
	assertEqual(t, expectedProjectVersion, artifact.Version)
	assertEqual(t, "Apache Software License 2.0", artifact.Licenses[0].Name)
	assertEqual(t, 2, len(artifact.Licenses))
}

func TestLicenseReportToBuildConfig(t *testing.T) {
	licenseReport := ReadLicenseReportFile(licenseReportTestInputFile)
	licenses := ReadLicenses(testLicensesFile)
	config := LicenseReportToBuildConfig(licenses, licenseReport)
	assertEqual(t, 2, len(config.Projects))
	assertEqual(t, expectedProjectVersion, config.Projects["org.test.project1"].Version)
	assertEqual(t, 1, len(config.Projects["org.test.project1"].MavenArtifacts))
}
