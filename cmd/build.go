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
	"github.com/pgier/hawkbuild/util"
	"github.com/spf13/cobra"
)

var projectName string

// mavenBuildCmd represents the mavenBuild command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Integration with brew maven-build command",
	Long: `Provides integration with the 'brew maven-build' command.

Can generate a command from a hawkbuild config.
Note that this command requires a valid config file, either
the default value or one set using --config <config-file>.`,
	PreRun: func(cmd *cobra.Command, args []string) {
		fileMode, err := os.Stat(configFiles[0])
		if os.IsNotExist(err) {
			fmt.Printf("Error: Config file does not exist: %v\n\n", configFiles[0])
			cmd.Usage()
			os.Exit(1)
		} else if fileMode.IsDir() {
			fmt.Printf("Error: Given config file is a directory: %s\n\n", configFiles[0])
			cmd.Usage()
			os.Exit(1)
		}
		listProjects, err := cmd.Flags().GetBool("list")
		util.Check(err)
		if len(args) < 1 && !listProjects {
			fmt.Println("Error: Please provide project name")
			cmd.Usage()
			os.Exit(1)
		}
	},
	Run: func(cmd *cobra.Command, args []string) {
		runBuildCmd(cmd, args)
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
	buildCmd.Flags().Bool("list", false, "List all projects in the current configuration")
}

func runBuildCmd(cmd *cobra.Command, args []string) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Problem accessing input file ", r)
		}
	}()

	listProjects, err := cmd.Flags().GetBool("list")
	util.Check(err)

	util.CheckFileExists(configFiles[0])
	buildConfig := config.ReadBuildConfigFile(configFiles[0])
	if listProjects {
		for projName := range buildConfig.Projects {
			fmt.Println(projName)
		}
	} else {
		projName := args[0]
		if project, ok := buildConfig.Projects[projName]; ok {
			owner := project.Owner
			fmt.Printf("owner: %v\n", owner)
		} else {
			fmt.Printf("project: %v does not exist in the current config %v", projName, configFiles[0])
		}
	}
}
