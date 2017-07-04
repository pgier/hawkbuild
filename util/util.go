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

package util

import (
	"fmt"
	"os"
	"strings"
)

// Check if error is nil or panic
func Check(err error) {
	if err != nil {
		panic(err)
	}
}

// CheckFileExists if file exists and is not a directory or panic
func CheckFileExists(file string) {
	fileMode, err := os.Stat(file)
	if os.IsNotExist(err) {
		panic(err)
	} else if fileMode.IsDir() {
		panic(fmt.Sprintf("File is a directory: %s", file))
	}
}

// StringInSlice checks if a string is contained in the slice
func StringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

// IsEmptyString checks if a string contains only whitespace
func IsEmptyString(aStr string) bool {
	return len(strings.TrimSpace(aStr)) == 0
}
