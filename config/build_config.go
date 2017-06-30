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
	checkLicenseFile(licenseFile)
	licenseDat, err := ioutil.ReadFile(licenseFile)
	util.Check(err)
	var licenses = Licenses{}
	err = yaml.Unmarshal(licenseDat, &licenses)
	util.Check(err)
	return licenses
}

// ReadBuildConfig reads build configuration information from a yaml file
func ReadBuildConfig(buildConfigFile string) BuildConfig {
	checkBuildConfigFile(buildConfigFile)
	configDat, err := ioutil.ReadFile(buildConfigFile)
	util.Check(err)
	config := BuildConfig{}
	err = yaml.Unmarshal(configDat, &config)
	util.Check(err)
	return config
}

// WriteBuildConfig writes a build config to a file
func WriteBuildConfig(config BuildConfig, filename string) {
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

func checkLicenseFile(file string) {
	fmt.Printf("Reading license file %s\n", file)
	util.CheckFileExists(file)
}

func checkBuildConfigFile(file string) {
	fmt.Printf("Reading build config file %s\n", file)
	util.CheckFileExists(file)
}

func checkLicenseReportFile(file string) {
	fmt.Printf("Reading license report file %s\n", file)
	util.CheckFileExists(file)
}

// LicenseReport list of artifacts and associated licenses
type LicenseReport struct {
	XMLName   xml.Name        `xml:"licenseSummary"`
	Artifacts []MavenArtifact `xml:"dependencies>dependency"`
}

// ReadLicenseReportFile reads license report info from XML file
func ReadLicenseReportFile(licenseReportFile string) LicenseReport {
	checkLicenseReportFile(licenseReportFile)
	fileDat, err := ioutil.ReadFile(licenseReportFile)
	util.Check(err)
	licenseReport := LicenseReport{}
	err = xml.Unmarshal(fileDat, &licenseReport)
	util.Check(err)
	return licenseReport
}

// WriteLicenseReportFile writes license information to an XML file
func WriteLicenseReportFile(licenses Licenses, config BuildConfig, outputFile string) {
	licenseSummary := ConvertBuildConfigToLicenseSummary(licenses, config)
	xmlBytes, err := xml.MarshalIndent(licenseSummary, "", "  ")
	util.Check(err)

	f, err := os.Create(outputFile)
	util.Check(err)
	defer f.Close()
	f.WriteString(xml.Header + string(xmlBytes))
	fmt.Printf("Wrote license report file: %s\n", outputFile)
}

// ConvertBuildConfigToLicenseSummary converts from build config to license summary
func ConvertBuildConfigToLicenseSummary(licenses Licenses, config BuildConfig) LicenseReport {
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

// LicenseReportToBuildConfig converts from license summary to build config
func LicenseReportToBuildConfig(licenses Licenses, licReport LicenseReport) BuildConfig {
	config := BuildConfig{Projects: make(map[string]Project)}
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
