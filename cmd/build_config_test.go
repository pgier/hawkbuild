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

	if found, artifact := FindArtifact(expectedArtifact1, licenseReportXML.Artifacts); found {
		assertEqual(t, "1.0", artifact.Version)
	} else {
		t.Fatal("No matching artifact in license summary")
	}

}

// FindArtifact searches the list to find an artifact with the given id
func FindArtifact(artifactID string, artifactList []MavenArtifact) (bool, MavenArtifact) {
	for _, artifact := range artifactList {
		if artifact.ArtifactID == artifactID {
			return true, artifact
		}
	}
	return false, MavenArtifact{}
}

func TestConvertBuildConfigToLicenseSummary(t *testing.T) {
	licenses := ReadLicenses(testLicensesFile)
	config := ReadBuildConfig(testBuildConfigFile)
	assertEqual(t, expectedProjectVersion, config.Projects[testProjectName].Version)
	assertEqual(t, expectedArtifact1, config.Projects[testProjectName].MavenArtifacts[0].ArtifactID)
	licenseSummary := ConvertBuildConfigToLicenseSummary(licenses, config)
	ok, artifact := FindArtifact(expectedArtifact1, licenseSummary.Artifacts)
	assertTrue(t, ok)
	assertEqual(t, expectedArtifact1, artifact.ArtifactID)
}

func TestReadLicenseReportFile(t *testing.T) {
	licenseSummary := ReadLicenseReportFile(licenseReportTestInputFile)
	assertEqual(t, expectedProjectVersion, licenseSummary.Artifacts[0].Version)
}

func TestConvertLicenseReportToBuildConfig(t *testing.T) {
	licenseSummary := ReadLicenseReportFile(licenseReportTestInputFile)
	config := ConvertLicenseSummaryToBuildConfig(licenseSummary)
	assertEqual(t, expectedProjectVersion, config.Projects["org.test.project1"].Version)
}
