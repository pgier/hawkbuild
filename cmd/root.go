// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"os"

	"github.com/pgier/hawkbuild/util"
	"github.com/spf13/cobra"
)

const (
	defaultConfigFileName = "hawkbuild.yaml"
)

var configFiles []string
var verbose bool

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "hawkbuild",
	Short: "CLI app for managing build configuration",
	//Long:  "CLI app for managing build configuration",
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}

func init() {
	RootCmd.PersistentFlags().StringArrayVarP(&configFiles, "config", "c",
		[]string{}, "Config file(s) containing build config and license info. "+
			"By default will look for \"hawkbuild.yaml\" in the current directory")
	RootCmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false,
		"Enable verbose output")
}

func checkConfigFile(configFiles []string) {
	if len(configFiles) == 0 {
		util.CheckFileExists(defaultConfigFile)
	} else {
		for _, cfgFile := range configFiles {
			util.CheckFileExists(cfgFile)
		}
	}
}
