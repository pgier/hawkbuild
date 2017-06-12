package cmd

import "testing"

const testLicensesFile = "testdata/licenses.yaml"
const testBuildConfigFile = "testdata/build-config.yaml"
const testProjectName = "my-test-project1"
const apacheLicenseShortName = "ASL 2.0"

func TestReadLicenses(t *testing.T) {
	licenses := ReadLicenses(testLicensesFile)
	const expectedLicenseCount = 2
	assertEqual(t, expectedLicenseCount, len(licenses))
	const expectedLicenseShortName = "ASL 2.0"
	assertEqual(t, expectedLicenseShortName, licenses["Apache Software License 2.0"].ShortName)
}

func TestReadLicenseFileMap(t *testing.T) {
	licenses := ReadLicenses(testLicensesFile)
	const expectedLicenseCount = 2
	assertEqual(t, expectedLicenseCount, len(licenses))
	const expectedLicenseName = "EPL"
	assertEqual(t, expectedLicenseName, licenses["Eclipse Public License, Version 1.0"].ShortName)
}

const expectedProjectVersion = "1.0"

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

const licenseSummaryXMLOutputFile = "testoutput/license-report.xml"

const expectedArtifact1 = "test-project1-artifact1"

func TestWriteProjectLicenseSummaryXML(t *testing.T) {
	licenses := ReadLicenses(testLicensesFile)
	buildConfig := ReadBuildConfig(testBuildConfigFile)
	assertEqual(t, expectedProjectVersion, buildConfig.Projects[testProjectName].Version)
	WriteLicenseReportFile(licenses, buildConfig, licenseSummaryXMLOutputFile)
	licenseSummaryXML := ReadLicenseReportFile(licenseSummaryXMLOutputFile)

	if found, artifact := FindArtifact(expectedArtifact1, licenseSummaryXML.Artifacts); found {
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

const licenseSummaryTestInputFile = "testdata/license-report.xml"

const buildConfigOutputFile = "testoutput/build-config-test.yaml"

func TestReadLicenseReportFile(t *testing.T) {
	licenseSummary := ReadLicenseReportFile(licenseSummaryTestInputFile)
	assertEqual(t, expectedProjectVersion, licenseSummary.Artifacts[0].Version)
}

func TestConvertLicenseReportToBuildConfig(t *testing.T) {
	licenseSummary := ReadLicenseReportFile(licenseSummaryTestInputFile)
	config := ConvertLicenseSummaryToBuildConfig(licenseSummary)
	assertEqual(t, expectedProjectVersion, config.Projects["org.test.project1"].Version)
}
