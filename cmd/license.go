package cmd

import (
	"encoding/xml"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

const (
	// LicenseCmdDescription provides help info for the license command
	LicenseCmdDescription    = "Generate a license report from a build configuration yaml file"
	defaultLicenseFile       = "licenses.yaml"
	defaultConfigFile        = "build-config.yaml"
	defaultLicenseReportFile = "license-summary.xml"
)

// LicenseCmd runs the "license" command
func LicenseCmd(args []string) {
	licenseFlags := flag.NewFlagSet("license", flag.ExitOnError)
	configFile := licenseFlags.String("config", defaultConfigFile, "Build config file")
	licenseFlags.StringVar(configFile, "c", defaultConfigFile, "Same as -config")
	licenseFile := licenseFlags.String("license-list", defaultLicenseFile, "Path to licenses file")
	licenseFlags.StringVar(licenseFile, "l", defaultLicenseFile, "Same as -license-list")
	reportFile := licenseFlags.String("report", defaultLicenseReportFile, "Path to license report file")
	licenseFlags.StringVar(reportFile, "r", defaultLicenseReportFile, "Same as -report")
	reverseFlag := licenseFlags.Bool("reverse", false, "Reverse process so that build config is generated from license report")
	licenseFlags.BoolVar(reverseFlag, "v", false, "Same as -reverse")
	err := licenseFlags.Parse(args)
	check(err)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Problem accessing input file ", r)
		}
	}()
	checkLicenseFile(*licenseFile)
	licenses := ReadLicenses(*licenseFile)

	if *reverseFlag {
		checkLicenseReportFile(*reportFile)
		licReport := ReadLicenseReportFile(*reportFile)
		buildConfig := LicenseReportToBuildConfig(licenses, licReport)
		WriteBuildConfig(buildConfig, *configFile)
	} else {
		checkBuildConfigFile(*configFile)
		buildConfig := ReadBuildConfig(*configFile)
		WriteLicenseReportFile(licenses, buildConfig, *reportFile)
	}
}

func checkLicenseFile(file string) {
	if file == defaultLicenseFile {
		fmt.Printf("Reading default license file %s\n", defaultLicenseFile)
	} else {
		fmt.Printf("Reading license file %s\n", file)
	}
	checkFile(file)
}

func checkBuildConfigFile(file string) {
	if file == defaultConfigFile {
		fmt.Printf("Reading default build config file %s\n", defaultConfigFile)
	} else {
		fmt.Printf("Reading build config file %s\n", file)
	}
	checkFile(file)
}

func checkLicenseReportFile(file string) {
	if file == defaultLicenseReportFile {
		fmt.Printf("Reading default license report file %s\n", defaultLicenseFile)
	} else {
		fmt.Printf("Reading license report file %s\n", file)
	}
	checkFile(file)
}

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
	fmt.Printf("Wrote license report file: %s\n", outputFile)
}

// ConvertBuildConfigToLicenseSummary converts from build config to license summary
func ConvertBuildConfigToLicenseSummary(licenses Licenses, config BuildConfig) LicenseReport {
	licShortNameMap := createLicenseShortNameMap(licenses)
	licenseReport := LicenseReport{}
	for _, project := range config.Projects {
		for _, artifact := range project.MavenArtifacts {
			artifact.Licenses = getMatchingLicenses(licShortNameMap, project.Licenses)
			if IsEmptyString(artifact.Version) {
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
				if !StringInSlice(license.ShortName, project.Licenses) {
					project.Licenses = append(project.Licenses, license.ShortName)
				}
			} else {
				fmt.Printf("License '%s' not found in license config.\n", nextLicense.Name)
			}
		}
		if IsEmptyString(project.Version) {
			project.Version = artifact.Version
		}
		if IsEmptyString(project.MavenGroupID) {
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
		if StringInSlice(name, license.AltNames) {
			return license, true
		}
	}
	return License{}, false
}
