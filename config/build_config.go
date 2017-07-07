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

// BuildConfig represents a list of project build configs and related info
type BuildConfig struct {
	defaultTarget string
	Projects      map[string]Project
	Licenses      map[string]License
}

// ReadBuildConfigFiles reads a list of build config files
func ReadBuildConfigFiles(configFiles []string) []BuildConfig {
	configs := []BuildConfig{}
	for _, configFile := range configFiles {
		configs = append(configs, ReadBuildConfigFile(configFile))
	}
	return configs
}

// ReadBuildConfigFile reads build configuration information from a yaml file
func ReadBuildConfigFile(buildConfigFile string) BuildConfig {
	util.CheckFileExists(buildConfigFile)
	configDat, err := ioutil.ReadFile(buildConfigFile)
	util.Check(err)
	return readBuildConfig(configDat)
}

// readBuildConfig reads build configuration information from a byte array
func readBuildConfig(buildConfig []byte) BuildConfig {
	config := BuildConfig{}
	err := yaml.Unmarshal(buildConfig, &config)
	util.Check(err)
	return config
}

// writeBuildConfig writes a build config to a file
func writeBuildConfig(config BuildConfig, filename string) {
	configData, err := yaml.Marshal(config)
	util.Check(err)
	f, err := os.Create(filename)
	util.Check(err)
	defer f.Close()
	f.Write(configData)
}

// getLicenses gets a list of license structs with the given names from the map
func getMatchingLicenses(configs []BuildConfig, licNames []string) []License {
	licList := []License{}
	configs = append(configs, DefaultLicenseConfig)
	for _, name := range licNames {
		for _, config := range configs {
			if license, ok := config.Licenses[name]; ok {
				licList = append(licList, license)
				break
			}
		}
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

// licenseReportToBuildConfig converts from license summary to build config
func licenseReportToBuildConfig(configs []BuildConfig, licReport LicenseReport) BuildConfig {
	config := BuildConfig{Projects: make(map[string]Project)}
	for _, artifact := range licReport.Artifacts {
		if _, exists := config.Projects[artifact.GroupID]; !exists {
			config.Projects[artifact.GroupID] = Project{Version: artifact.Version, Licenses: []string{}}
		}
		project := config.Projects[artifact.GroupID]
		for _, nextLicense := range artifact.Licenses {
			if license, found := findLicenseByName(configs, nextLicense.Name); found {
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

func findLicenseByName(configs []BuildConfig, name string) (License, bool) {
	for _, config := range configs {
		if license, found := config.Licenses[name]; found {
			return license, true
		}
	}
	return License{}, false
}

// GenerateBasicBuildConfig generates a basic config file from built-in template
func GenerateBasicBuildConfig(outputFile string) {
	buildConfig := readBuildConfig([]byte(BuildConfigTemplate))
	writeBuildConfig(buildConfig, outputFile)
}

// MergeConfigFiles merges a list of configs into a single one
func MergeConfigFiles(configs []BuildConfig) BuildConfig {
	newConfig := BuildConfig{}
	// TODO:implement this
	return newConfig
}
