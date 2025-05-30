// Given a list of strings that represent files, print the directory structure.

// Input files = [
//   “/app/components/header”,
//   “/app/services”,
//   “/app/tests/components/header”,
//   "/app/components/services",
//   “/images/image.png”,
//   "/tsconfig.json",
//   "/index.html",
// ];
// Output:
// -app
// --components
// ---header
// --services
// --tests
// ---components
// ----header
// -images
// --Image.png
// -tsconfig.json
// -index.html

package main

import (
	"fmt"
	"sort"
	"strings"
)

var files = []string{
	"/app/components/header",
	"/app/services",
	"/app/tests/components/header",
	"/app/components/services",
	"/images/image.png",
	"/tsconfig.json",
	"/index.html",
}

func main() {
	directories := make(map[string]interface{}, 0)

	for _, file := range files {
		parts := strings.Split(file, "/")[1:]
		addToDirectory(parts, directories)
	}

	fmt.Println("Directory Structure: ", directories)

	printDirectory(directories, 0)
}

func addToDirectory(parts []string, directories map[string]interface{}) {
	if len(parts) == 0 {
		return
	}

	if len(parts) == 1 {
		directories[parts[0]] = nil
		return
	}

	if _, ok := directories[parts[0]]; !ok {
		directories[parts[0]] = make(map[string]interface{}, 0)
	}

	addToDirectory(parts[1:], directories[parts[0]].(map[string]interface{}))
}

func printDirectory(directories map[string]interface{}, level int) {
	keys := make([]string, 0)
	for key := range directories {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	for _, key := range keys {
		fmt.Println(strings.Repeat("-", level+1), key)
		if directories[key] != nil {
			printDirectory(directories[key].(map[string]interface{}), level+1)
		}
	}
}
