package util

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
