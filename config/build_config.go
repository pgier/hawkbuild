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
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pgier/hawkbuild/util"
	"gopkg.in/yaml.v2"
)

// License represents a software license for a package/project
type License struct {
	Name        string   `xml:"name"`
	ShortName   string   `xml:"-" yaml:"short-name"`
	AltNames    []string `xml:"-" yaml:"alt-names"`
	UpstreamURL string   `xml:"url" yaml:"upstream-url"`
}

// Licenses map of license name to license metadata
type Licenses map[string]License

// MavenArtifact represents a jar file or other Maven build artifact
type MavenArtifact struct {
	XMLName    xml.Name  `xml:"dependency" yaml:"-"`
	GroupID    string    `xml:"groupId" yaml:"groupId,omitempty"`
	ArtifactID string    `xml:"artifactId" yaml:"artifactId"`
	Version    string    `xml:"version" yaml:"-"`
	Type       string    `xml:"-" yaml:",omitempty"`
	Licenses   []License `xml:"licenses>license" yaml:"-"`
}

// Project contains project metadata and build configuration
type Project struct {
	Owner            string
	Version          string
	UpstreamVersion  string          `yaml:"upstream-version,omitempty"`
	UpstreamScmURL   string          `yaml:"upstream-scm-url"`
	MavenGroupID     string          `yaml:"maven-groupid,omitempty"`
	PackageName      string          `yaml:"package-name"`
	MavenPackageName string          `yaml:"maven-package-name,omitempty"`
	Licenses         []string        `yaml:"licenses"`
	MavenArtifacts   []MavenArtifact `yaml:"maven-artifacts" xml:"dependencies"`
	BuildType        string          `yaml:"build-type"`
	ScmURL           string          `yaml:"scm-url"`
	CommitID         string
	BuildOptions     string `yaml:"build-options,omitempty"`
}

// ProductConfig represents a list of project build configs and related info
type ProductConfig struct {
	defaultTarget string
	Projects      map[string]Project
}

// ReadLicenseConfig reads license information from a yaml file
// will panic if the file is not found
func ReadLicenseConfig(licenseFile string) Licenses {
	var licenses = Licenses{}
	err := yaml.Unmarshal([]byte(DefaultLicenseConfig), &licenses)
	util.Check(err)

	fileMode, err := os.Stat(licenseFile)
	if err == nil && !fileMode.IsDir() {
		licenseDat, err := ioutil.ReadFile(licenseFile)
		util.Check(err)
		err = yaml.Unmarshal(licenseDat, &licenses)
		util.Check(err)
	}

	return licenses
}

// ReadProductConfig reads build configuration information from a yaml file
func ReadProductConfig(productConfigFile string) ProductConfig {
	util.CheckFileExists(productConfigFile)
	configDat, err := ioutil.ReadFile(productConfigFile)
	util.Check(err)
	config := ProductConfig{}
	err = yaml.Unmarshal(configDat, &config)
	util.Check(err)
	return config
}

// WriteProductConfig writes a build config to a file
func WriteProductConfig(config ProductConfig, filename string) {
	configData, err := yaml.Marshal(config)
	util.Check(err)
	f, err := os.Create(filename)
	util.Check(err)
	defer f.Close()
	f.Write(configData)
	fmt.Printf("Wrote build config file: %s\n", filename)
}

// getLicenses gets a list of license structs with the given names from the map
func getMatchingLicenses(licenses map[string]License, licNames []string) []License {
	licList := []License{}
	for _, name := range licNames {
		licList = append(licList, licenses[name])
	}
	return licList
}

func findArtifactByArtifactID(artifacts []MavenArtifact, artifactID string) (bool, MavenArtifact) {
	for _, artifact := range artifacts {
		if artifact.ArtifactID == artifactID {
			return true, artifact
		}
	}
	return false, MavenArtifact{}
}

// LicenseReport list of artifacts and associated licenses
type LicenseReport struct {
	XMLName   xml.Name        `xml:"licenseSummary"`
	Artifacts []MavenArtifact `xml:"dependencies>dependency"`
}

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
func WriteLicenseReportFile(licenses Licenses, config ProductConfig, outputFile string) {
	licenseSummary := ConvertProductConfigToLicenseSummary(licenses, config)
	xmlBytes, err := xml.MarshalIndent(licenseSummary, "", "  ")
	util.Check(err)

	f, err := os.Create(outputFile)
	util.Check(err)
	defer f.Close()
	f.WriteString(xml.Header + string(xmlBytes))
	fmt.Printf("Wrote license report file: %s\n", outputFile)
}

// ConvertProductConfigToLicenseSummary converts from build config to license summary
func ConvertProductConfigToLicenseSummary(licenses Licenses, config ProductConfig) LicenseReport {
	licShortNameMap := createLicenseShortNameMap(licenses)
	licenseReport := LicenseReport{}
	for _, project := range config.Projects {
		for _, artifact := range project.MavenArtifacts {
			artifact.Licenses = getMatchingLicenses(licShortNameMap, project.Licenses)
			if util.IsEmptyString(artifact.Version) {
				artifact.Version = project.Version
			}
			licenseReport.Artifacts = append(licenseReport.Artifacts, artifact)
		}
	}
	return licenseReport
}

// createLicenseShortNameMap creates a map keyed by the short name instead of the full license name
func createLicenseShortNameMap(licenses map[string]License) map[string]License {
	licensesByShortName := make(map[string]License)
	for name, license := range licenses {
		if license.Name == "" {
			license.Name = name
		}
		licensesByShortName[license.ShortName] = license
	}
	return licensesByShortName
}

// LicenseReportToProductConfig converts from license summary to build config
func LicenseReportToProductConfig(licenses Licenses, licReport LicenseReport) ProductConfig {
	config := ProductConfig{Projects: make(map[string]Project)}
	for _, artifact := range licReport.Artifacts {
		if _, exists := config.Projects[artifact.GroupID]; !exists {
			config.Projects[artifact.GroupID] = Project{Version: artifact.Version, Licenses: []string{}}
		}
		project := config.Projects[artifact.GroupID]
		for _, nextLicense := range artifact.Licenses {
			if license, found := findLicenseByName(licenses, nextLicense.Name); found {
				if !util.StringInSlice(license.ShortName, project.Licenses) {
					project.Licenses = append(project.Licenses, license.ShortName)
				}
			} else {
				fmt.Printf("License '%s' not found in license config.\n", nextLicense.Name)
			}
		}
		if util.IsEmptyString(project.Version) {
			project.Version = artifact.Version
		}
		if util.IsEmptyString(project.MavenGroupID) {
			project.MavenGroupID = artifact.GroupID
		}
		if project.MavenGroupID == artifact.GroupID {
			artifact.GroupID = ""
		}
		project.MavenArtifacts = append(project.MavenArtifacts, artifact)
		config.Projects[project.MavenGroupID] = project
	}
	return config
}

func findLicenseByName(licenses Licenses, name string) (License, bool) {
	if license, found := licenses[name]; found {
		return license, true
	}
	for _, license := range licenses {
		if util.StringInSlice(name, license.AltNames) {
			return license, true
		}
	}
	return License{}, false
}
