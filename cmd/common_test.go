package cmd

import (
	"fmt"
	"testing"
)

func assertEqual(t *testing.T, expected interface{}, actual interface{}) {
	if expected != actual {
		message := fmt.Sprintf("Expected: %v  Actual: %v", expected, actual)
		t.Fatal(message)
	}
}

func assertTrue(t *testing.T, value bool) {
	if !value {
		message := fmt.Sprintf("Expected: true  Actual: %v", value)
		t.Fatal(message)
	}
}
