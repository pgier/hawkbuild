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

package test

import (
	"fmt"
	"testing"
)

// AssertEqual assert that the two values are equal
func AssertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		message := fmt.Sprintf("Expected: %v  Actual: %v", expected, actual)
		t.Fatal(message)
	}
}

// AssertTrue assert that the given value is true
func AssertTrue(t *testing.T, value bool) {
	if !value {
		message := fmt.Sprintf("Expected: true  Actual: %v", value)
		t.Fatal(message)
	}
}
