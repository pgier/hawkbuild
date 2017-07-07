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
	"github.com/pgier/hawkbuild/config"
	"github.com/spf13/cobra"
)

const (
	defaultConfigOutputFile string = "hawkbuild-new.yaml"
)

var (
	licenseReportFile string
	configOutputFile  string
)

// generateCmd represents the generate command
var generateCmd = &cobra.Command{
	Use:   "generate",
	Short: "Generate a config file",
	Long: `Generate a config file.  By default will generate a basic config
file template.  Can also be used to generate a config file from
another format.  For example, to generate a config from an existing
license report:

	hawkbuild generate -l license-report.xml -o my-new-config.yaml`,
	Run: func(cmd *cobra.Command, args []string) {

		if licenseReportFile != "" {
			config.GenerateBuildConfigFromLicenseReportFile(configFiles,
				licenseReportFile, configOutputFile)
		} else {
			config.GenerateBasicBuildConfig(configOutputFile)
		}
	},
}

func init() {
	RootCmd.AddCommand(generateCmd)

	generateCmd.Flags().StringVarP(&licenseReportFile, "license-report", "l", "",
		"Generate the build config from an existing license report file.")
	generateCmd.Flags().StringVarP(&configOutputFile, "output", "o",
		defaultConfigOutputFile, "File path to output the generated build config "+
			"yaml file.  Defaults to "+defaultConfigOutputFile)
}
