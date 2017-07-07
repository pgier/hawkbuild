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
	"encoding/xml"
	"io/ioutil"
	"os"

	"github.com/pgier/hawkbuild/util"
)

// ReadLicenseReportFile reads license report info from XML file
func ReadLicenseReportFile(licenseReportFile string) LicenseReport {
	util.CheckFileExists(licenseReportFile)
	fileDat, err := ioutil.ReadFile(licenseReportFile)
	util.Check(err)
	licenseReport := LicenseReport{}
	err = xml.Unmarshal(fileDat, &licenseReport)
	util.Check(err)
	return licenseReport
}

// WriteLicenseReportFile writes license information to an XML file
func WriteLicenseReportFile(configs []BuildConfig, outputFile string) {
	licenseReport := generateLicenseReport(configs)
	xmlBytes, err := xml.MarshalIndent(licenseReport, "", "  ")
	util.Check(err)

	f, err := os.Create(outputFile)
	util.Check(err)
	defer f.Close()
	f.WriteString(xml.Header + string(xmlBytes))
}

// GenerateLicenseReport generates a license report from a build config
func GenerateLicenseReport(configFiles []string, outputFile string) {
	configs := ReadBuildConfigFiles(configFiles)
	configs = append(configs, DefaultLicenseConfig)
	WriteLicenseReportFile(configs, outputFile)
}

// generateLicenseReport generates a license report from a build config
func generateLicenseReport(configs []BuildConfig) LicenseReport {
	licenseReport := LicenseReport{}
	for _, config := range configs {
		for _, project := range config.Projects {
			for _, artifact := range project.MavenArtifacts {
				artifact.Licenses = getMatchingLicenses(configs, project.Licenses)
				if util.IsEmptyString(artifact.Version) {
					artifact.Version = project.Version
				}
				if util.IsEmptyString(artifact.GroupID) {
					artifact.GroupID = project.MavenGroupID
				}
				licenseReport.Artifacts = append(licenseReport.Artifacts, artifact)
			}
		}
	}
	return licenseReport
}

// GenerateBuildConfigFromLicenseReportFile generate a build config file from
// an existing license report file
func GenerateBuildConfigFromLicenseReportFile(existingConfigFiles []string,
	licReportFile string, outputConfigFile string) {
	configs := ReadBuildConfigFiles(existingConfigFiles)
	configs = append(configs, DefaultLicenseConfig)
	licReport := ReadLicenseReportFile(licReportFile)
	generatedConfig := licenseReportToBuildConfig(configs, licReport)
	writeBuildConfig(generatedConfig, outputConfigFile)
}
