package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"
)

// LicenseReport list of artifacts and associated licenses
type LicenseReport struct {
	XMLName   xml.Name        `xml:"licenseSummary"`
	Artifacts []MavenArtifact `xml:"dependencies>dependency"`
}

// ReadLicenseReportFile reads license report info from XML file
func ReadLicenseReportFile(licenseSummaryFile string) LicenseReport {
	fileDat, err := ioutil.ReadFile(licenseSummaryFile)
	check(err)
	licenseReport := LicenseReport{}
	err = xml.Unmarshal(fileDat, &licenseReport)
	check(err)
	return licenseReport
}

// WriteLicenseReportFile writes license information to an XML file
func WriteLicenseReportFile(licenses Licenses, config BuildConfig, outputFile string) {
	licenseSummary := ConvertBuildConfigToLicenseSummary(licenses, config)
	xmlBytes, err := xml.MarshalIndent(licenseSummary, "", "  ")
	check(err)

	f, err := os.Create(outputFile)
	check(err)
	defer f.Close()
	f.WriteString(xml.Header + string(xmlBytes))
}

// ConvertBuildConfigToLicenseSummary converts from build config to license summary
func ConvertBuildConfigToLicenseSummary(licenses Licenses, config BuildConfig) LicenseReport {
	licMap := licenses.Map
	licenseReport := LicenseReport{}
	for _, project := range config.Projects {
		for _, artifact := range project.MavenArtifacts {
			artifact.Licenses = GetLicenses(licMap, project.Licenses)
			if IsEmptyString(artifact.Version) {
				artifact.Version = project.Version
			}
			licenseReport.Artifacts = append(licenseReport.Artifacts, artifact)
		}
	}
	return licenseReport
}

// ConvertLicenseSummaryToBuildConfig converts from license summary to build config
func ConvertLicenseSummaryToBuildConfig(licReport LicenseReport) BuildConfig {
	config := BuildConfig{Projects: make(map[string]Project)}
	for _, artifact := range licReport.Artifacts {
		if _, ok := config.Projects[artifact.GroupID]; !ok {
			config.Projects[artifact.GroupID] = Project{Version: artifact.Version}
		}
		project := config.Projects[artifact.GroupID]
		for _, license := range artifact.Licenses {
			if !StringInSlice(license.ShortName, project.Licenses) {
				project.Licenses = append(project.Licenses, license.ShortName)
			}
		}
		if IsEmptyString(project.Version) {
			project.Version = artifact.Version
		}
		config.Projects[artifact.GroupID] = project
	}
	return config
}
