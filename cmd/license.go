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

package cmd

import (
	"fmt"

	"github.com/pgier/hawkbuild/config"
	"github.com/pgier/hawkbuild/util"
	"github.com/spf13/cobra"
)

const (
	// LicenseCmdDescription provides help info for the license command
	LicenseCmdDescription    = "Generate a license report from a build configuration yaml file"
	defaultLicenseConfigFile = "licenses.yaml"
	defaultConfigFile        = "build-config.yaml"
	defaultLicenseReportFile = "license-summary.xml"
)

var licenseCmd = &cobra.Command{
	Use:   "license",
	Short: "Management and reporting for project licenses",
	Long:  "Management and reporting for project licenses",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("license called")
		runLicenseCmd(cmd)
	},
}

func init() {
	RootCmd.AddCommand(licenseCmd)

	licenseCmd.Flags().StringP("license-config", "l",
		defaultLicenseConfigFile, "License report XML file")
	licenseCmd.Flags().StringP("report", "t",
		defaultLicenseReportFile, "License report XML file")
	licenseCmd.Flags().BoolP("reverse", "r", false,
		"Reverse process so that build config is generated from license")
}

func runLicenseCmd(cmd *cobra.Command) {
	configFile, err := cmd.Flags().GetString("config")
	util.Check(err)
	licenseConfigFile, err := cmd.Flags().GetString("license-config")
	util.Check(err)
	licenseReportFile, err := cmd.Flags().GetString("report")
	util.Check(err)
	reverse, err := cmd.Flags().GetBool("reverse")
	util.Check(err)

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Problem accessing input file ", r)
		}
	}()

	if licenseConfigFile != defaultLicenseConfigFile {
		util.CheckFileExists(licenseConfigFile)
	}
	licenses := config.ReadLicenseConfig(licenseConfigFile)

	if reverse {
		licReport := config.ReadLicenseReportFile(licenseReportFile)
		buildConfig := config.LicenseReportToBuildConfig(licenses, licReport)
		config.WriteBuildConfig(buildConfig, configFile)
	} else {
		buildConfig := config.ReadBuildConfig(configFile)
		config.WriteLicenseReportFile(licenses, buildConfig, licenseReportFile)
	}
}
