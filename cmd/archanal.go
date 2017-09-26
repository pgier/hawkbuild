// Copyright © 2017 NAME HERE <EMAIL ADDRESS>
//
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
	"archive/zip"
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/spf13/cobra"
)

var (
	archiveFileName string
)

// archanalCmd represents the archanal command
var archanalCmd = &cobra.Command{
	Use:   "archanal <filename>",
	Short: "Analyze an archive file",
	Long: `Analyze an archive file

Default behaviour is to print the tree of nested archives inside an 
archive (zip, jar, etc).`,
	PreRun: func(cmd *cobra.Command, args []string) {

		if len(args) < 1 {
			fmt.Println("Error: Please provide archive filename")
			cmd.Usage()
			os.Exit(1)
		}
		archiveFileName = args[0]
	},
	Run: func(cmd *cobra.Command, args []string) {
		tree := CreateArchiveTree(archiveFileName)
		PrintTree(tree, "  ")
	},
}

func init() {
	RootCmd.AddCommand(archanalCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// archanalCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// archanalCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// ArchiveNode represents a node/entry in the archive
type ArchiveNode struct {
	Name     string        // The filename in the archive
	children []ArchiveNode // Nested archives
}

const achiveFileTypesRegex = ".*(\\.(ear|jar|war|zip))"

var archiveFileTypes = regexp.MustCompile(achiveFileTypesRegex)

// CreateArchiveTree creates a tree of ArchiveNodes which
// contains any nested archives
func CreateArchiveTree(filename string) ArchiveNode {
	buffer, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("File: %s\n", filename)
	zipReader, err := zip.NewReader(bytes.NewReader(buffer), int64(len(buffer)))
	if err != nil {
		log.Fatal(err)
	}
	return readArchive(filename, zipReader)
}

func readArchive(filename string, reader *zip.Reader) ArchiveNode {
	node := ArchiveNode{
		Name:     filename,
		children: []ArchiveNode{},
	}
	for _, file := range reader.File {
		if archiveFileTypes.MatchString(file.Name) {
			filename := Filename(file.Name)

			readCloser, err := file.Open()
			if err != nil {
				log.Fatal(err)
			}
			buffer, err := ioutil.ReadAll(readCloser)
			readCloser.Close()
			if err != nil {
				log.Fatal(err)
			}
			zipReader, err := zip.NewReader(bytes.NewReader(buffer), int64(len(buffer)))
			if err != nil {
				log.Fatal(err)
			}
			childNode := readArchive(filename, zipReader)
			node.children = append(node.children, childNode)
		}
	}
	return node
}

const fileSep = "/"

// Filename returns just the filename portion from a path string
func Filename(filepath string) string {
	return filepath[(strings.LastIndex(filepath, fileSep) + 1):]
}

// PrintTree print the tree from the given node starting at
// the given level of indentation
func PrintTree(node ArchiveNode, indent string) {
	fmt.Printf("%s%s\n", indent, node.Name)
	for _, child := range node.children {
		PrintTree(child, indent+indent)
	}
}
