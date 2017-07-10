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
	Long: "Management and reporting for project licenses." +
		"by default will print all licenses in the current config.",
	Run: func(cmd *cobra.Command, args []string) {
		runLicenseCmd()
	},
}

var licenseReportCmd = &cobra.Command{
	Use:   "report",
	Short: "Generarte license report",
	Long:  "Print a license report to a file",
	Run: func(cmd *cobra.Command, args []string) {
		runLicenseReportCmd(reportOutputFile)
	},
}

var (
	spdxID           bool
	reportOutputFile string
)

func init() {
	RootCmd.AddCommand(licenseCmd)
	licenseCmd.Flags().BoolVar(&spdxID, "spdxid", false, "Print the SPDX identifier")
	licenseCmd.AddCommand(licenseReportCmd)
	licenseReportCmd.Flags().StringVarP(&reportOutputFile, "output", "o",
		defaultLicenseReportFile, "License report output file")
}

func runLicenseCmd() {
	for name, license := range config.DefaultLicenseConfig.Licenses {
		if spdxID {
			fmt.Println(license.Name)
		} else {
			fmt.Println(name)
		}
	}
}

func runLicenseReportCmd(outputFile string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Problem accessing input file ", r)
		}
	}()

	if len(configFiles) > 0 {
		config.GenerateLicenseReport(configFiles, outputFile)
	}
}
