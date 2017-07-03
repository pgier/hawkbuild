package cmd

import (
	"flag"
	"fmt"

	"github.com/pgier/hawkbuild/config"
	"github.com/pgier/hawkbuild/util"
)

const (
	// LicenseCmdDescription provides help info for the license command
	LicenseCmdDescription    = "Generate a license report from a build configuration yaml file"
	defaultLicenseConfigFile = "licenses.yaml"
	defaultConfigFile        = "build-config.yaml"
	defaultLicenseReportFile = "license-summary.xml"
)

// LicenseCmd runs the "license" command
func LicenseCmd(args []string) {
	licenseFlags := flag.NewFlagSet("license", flag.ExitOnError)
	configFile := licenseFlags.String("config", defaultConfigFile, "Build config file")
	licenseFlags.StringVar(configFile, "c", defaultConfigFile, "Same as -config")
	licenseConfigFile := licenseFlags.String("license-config", defaultLicenseConfigFile, "Path to licenses file")
	licenseFlags.StringVar(licenseConfigFile, "l", defaultLicenseConfigFile, "Same as -license-config")
	reportFile := licenseFlags.String("report", defaultLicenseReportFile, "Path to license report file")
	licenseFlags.StringVar(reportFile, "t", defaultLicenseReportFile, "Same as -report")
	reverseFlag := licenseFlags.Bool("reverse", false, "Reverse process so that build config is generated from license report")
	licenseFlags.BoolVar(reverseFlag, "r", false, "Same as -reverse")
	err := licenseFlags.Parse(args)
	util.Check(err)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Problem accessing input file ", r)
		}
	}()

	if *licenseConfigFile != defaultLicenseConfigFile {
		util.CheckFileExists(*licenseConfigFile)
	}
	licenses := config.ReadLicenseConfig(*licenseConfigFile)

	if *reverseFlag {
		licReport := config.ReadLicenseReportFile(*reportFile)
		buildConfig := config.LicenseReportToBuildConfig(licenses, licReport)
		config.WriteBuildConfig(buildConfig, *configFile)
	} else {
		buildConfig := config.ReadBuildConfig(*configFile)
		config.WriteLicenseReportFile(licenses, buildConfig, *reportFile)
	}
}
