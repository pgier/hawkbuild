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

	"github.com/pgier/hawkbuild/config"
	"github.com/spf13/cobra"
)

const (
	rhBuildCmd     = "rhpkg"
	fedoraBuildCmd = "fedpkg"
)

var (
	projectNameArg string
	listFlag       bool
	scratchFlag    bool
	printFlag      bool
)

// mavenBuildCmd represents the mavenBuild command
var buildCmd = &cobra.Command{
	Use:   "build project",
	Short: "Integration with brew maven-build command",
	Long: `Provides integration with the 'brew maven-build' command.

Can generate a command from a hawkbuild config.
Note that this command requires a valid config file, either
the default value or one set using --config <config-file>.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		if len(configFiles) < 1 {
			fmt.Println("You must specify at least one config file")
			cmd.Usage()
			os.Exit(1)
		}
		for i, configFile := range configFiles {
			fileMode, err := os.Stat(configFile)
			if os.IsNotExist(err) {
				fmt.Printf("Error: Config file does not exist: %v\n\n", configFiles[i])
				cmd.Usage()
				os.Exit(1)
			} else if fileMode.IsDir() {
				fmt.Printf("Error: Config file is a directory: %s\n\n", configFiles[i])
				cmd.Usage()
				os.Exit(1)
			}
		}

		if listFlag {
			return
		}
		if len(args) < 1 {
			fmt.Println("Error: Please provide project name")
			cmd.Usage()
			os.Exit(1)
		}
		projectNameArg = args[0]
	},
	Run: func(cmd *cobra.Command, args []string) {
		if listFlag {
			runBuildListCmd()
		} else {
			runBuildCmd(cmd, args)
		}
	},
}

func init() {
	RootCmd.AddCommand(buildCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// mavenBuildCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// mavenBuildCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	buildCmd.Flags().BoolVar(&listFlag, "list", false, "List all projects in the"+
		" current configuration")
	buildCmd.Flags().BoolVar(&scratchFlag, "scratch", true, "Run a scratch"+
		" (throwaway) build")
	buildCmd.Flags().BoolVar(&printFlag, "print", false, "Print the command to run"+
		" without actually running it")
}

func runBuildListCmd() {
	for _, configFile := range configFiles {
		nextConfig := config.ReadBuildConfigFile(configFile)
		for name := range nextConfig.Projects {
			fmt.Println(name)
		}
	}
}

func runBuildCmd(cmd *cobra.Command, args []string) {
	buildConfig := config.ReadBuildConfigFile(configFiles[0])
	if project, ok := buildConfig.Projects[projectNameArg]; ok {
		config.BuildProject(project, buildConfig.DefaultTarget, scratchFlag, printFlag)
	} else {
		fmt.Printf("project: %v does not exist in the current config %v", projectNameArg, configFiles[0])
	}
}
