package cmd

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"

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

// Project represents a project build configuration
type Project struct {
	Version string
	// MavenName represents the top level groupId:artifactId of the project
	MavenGroupID string `yaml:"maven-groupid,omitempty"`
	// MavenTopArtifact the top level/parent maven artifact of the project
	MavenTopArtifactID string          `yaml:"maven-top-artifactid,omitempty"`
	Licenses           []string        `yaml:"licenses"`
	MavenArtifacts     []MavenArtifact `yaml:"maven-artifacts" xml:"dependencies"`
}

// BuildConfig represents a list of project build configs
type BuildConfig struct {
	defaultTarget string
	Projects      map[string]Project
}

// ReadLicenses reads license information from a yaml file
func ReadLicenses(licenseFile string) Licenses {
	licenseDat, err := ioutil.ReadFile(licenseFile)
	check(err)
	var licenses = Licenses{}
	err = yaml.Unmarshal(licenseDat, &licenses)
	check(err)
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
