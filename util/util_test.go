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

func TestCheckFileExistsFail(t *testing.T) {
	badFile := "foo_bar.txt"
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The file %v was found, and it should not exist", badFile)
		}
	}()
	CheckFileExists(badFile)
}
