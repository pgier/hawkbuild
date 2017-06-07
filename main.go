package main

import "fmt"

func main() {
	licenses := ReadLicenses("testdata/licenses.yaml")
	fmt.Println(string(licenses.List[0].Name))
}
