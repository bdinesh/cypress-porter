/*
Copyright © 2021 Dinesh Bogolu

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

package cmd

import (
	"fmt"
	"github.com/bdinesh/cypress-porter/porter"
	"github.com/yargevad/filepathx"
	"log"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
)

// portCmd represents the port command
var portCmd = &cobra.Command{
	Use:   "port",
	Short: "Ports a Protractor file to Cypress",
	Run: func(cmd *cobra.Command, args []string) {
		for _, arg := range args {
			filePaths := getGlobMatches(arg)

			for _, path := range filePaths {
				portFile(path)
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(portCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// portCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// portCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func readFile(path string) string {
	data, err := os.ReadFile(path)

	if err != nil {
		log.Fatal(err)
	}

	return string(data)
}

func writeFile(path string, data []byte) {
	err := os.WriteFile(path, data, 0666)

	if err != nil {
		log.Fatal(err)
	}
}

func isPageFile(path string) bool {
	matched, err := filepath.Match("*Page.ts", path)

	if err != nil {
		log.Fatal(err)
	}

	return matched
}

func isSpecFile(path string) bool {
	matched, err := filepath.Match("*.spec.ts", path)

	if err != nil {
		log.Fatal(err)
	}

	return matched
}

func getGlobMatches(path string) []string {
	matches, err := filepathx.Glob(path)

	if err != nil {
		log.Fatal(err)
	}

	return matches
}

func portFile(path string) {
	fmt.Printf("Porting %s\n", path)
	d := readFile(path)
	fileName := filepath.Base(path)

	if isPageFile(fileName) {
		s := porter.PortPage(d)
		writeFile(path, []byte(s))
	} else if isSpecFile(fileName) {
		fmt.Println("Porting not implemented for spec files")
	} else {
		fmt.Printf("Porting not implemented for %s\n", fileName)
	}

	fmt.Printf("Ported %s\n", path)
}
