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
	"io/ioutil"
	"os"
	"testing"

	"github.com/pgier/hawkbuild/test"
)

func TestIsEmptyString(t *testing.T) {
	test.AssertTrue(t, IsEmptyString(""))
	test.AssertTrue(t, !IsEmptyString("foo"))
}

func TestStringInSlice(t *testing.T) {
	testSlice := []string{"a", "b", "c"}
	test.AssertTrue(t, StringInSlice("b", testSlice))
	test.AssertTrue(t, !StringInSlice("d", testSlice))
}

func TestCheckFileExists(t *testing.T) {
	tempFile, err := ioutil.TempFile(os.TempDir(), "hawkbuild-test")
	test.AssertTrue(t, err == nil)
	CheckFileExists(tempFile.Name())
}

func TestIsFile(t *testing.T) {
	tempFile, err := ioutil.TempFile(os.TempDir(), "hawkbuild-test-isfile")
	test.AssertTrue(t, err == nil)
	test.AssertTrue(t, IsFile(tempFile.Name()))
	test.AssertTrue(t, !IsFile("file/does/not/exist"))
}

func TestCheckFileExistsFail(t *testing.T) {
	badFile := "foo_bar.txt"
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The file %v was found, and it should not exist", badFile)
		}
	}()
	CheckFileExists(badFile)
}
