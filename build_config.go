package main

import (
	"encoding/xml"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

// License represents a software license for a package/project
type License struct {
	XMLName     xml.Name `xml:"license"`
	Name        string   `xml:"name" yaml:"name"`
	ShortName   string   `xml:"-" yaml:"short-name"`
	UpstreamURL string   `xml:"url" yaml:"upstream-url"`
}

// Licenses lists the licenses
type Licenses struct {
	List []License `yaml:"licenses"`
	Map  map[string]License
}

// MavenArtifact represents a jar file or other Maven build artifact
type MavenArtifact struct {
	XMLName    xml.Name `xml:"dependency"`
	GroupID    string   `xml:"groupId" yaml:"groupId"`
	ArtifactID string   `xml:"artifactId" yaml:"artifactId"`
	Version    string   `xml:"version"`
	Type       string   `xml:"-"`
	Licenses   []License
}

// Project represents a project build configuration
type Project struct {
	Name    string
	Version string
	// MavenName represents the top level groupId:artifactId of the project
	MavenName      string          `yaml:"maven-name"`
	Licenses       []string        `yaml:"licenses"`
	MavenArtifacts []MavenArtifact `yaml:"maven-artifacts" xml:"dependencies"`
}

// BuildConfig represents a list of project build configs
type BuildConfig struct {
	defaultTarget string
	Projects      map[string]Project
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// ReadLicenses reads license information from a yaml file
func ReadLicenses(licenseFile string) Licenses {
	licenseDat, err := ioutil.ReadFile(licenseFile)
	check(err)
	var licenses = Licenses{}
	err = yaml.Unmarshal(licenseDat, &licenses)
	check(err)
	licenses.Map = make(map[string]License)
	for _, license := range licenses.List {
		licenses.Map[license.ShortName] = license
	}
	return licenses
}

// ReadBuildConfig reads build configuration information from a yaml file
func ReadBuildConfig(buildConfigFile string) BuildConfig {
	configDat, err := ioutil.ReadFile(buildConfigFile)
	check(err)
	config := BuildConfig{}
	err = yaml.Unmarshal(configDat, &config)
	check(err)
	return config
}

// WriteBuildConfig writes a build config to a file
func WriteBuildConfig(config BuildConfig, filename string) {
	configData, err := yaml.Marshal(config)
	check(err)
	f, err := os.Create(filename)
	check(err)
	defer f.Close()
	f.Write(configData)
}

// GetLicenses gets a list of licenses with the given shortnames from the map
func GetLicenses(licMap map[string]License, licShortNames []string) []License {
	licList := make([]License, len(licShortNames))
	for i, shortName := range licShortNames {
		licList[i] = licMap[shortName]
	}
	return licList
}
