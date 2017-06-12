package cmd

import (
	"fmt"
	"os"
	"strings"
)

func check(err error) {
	//fmt.Println("got an error")
	if err != nil {
		panic(err)
	}
}

func checkFile(file string) {
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
