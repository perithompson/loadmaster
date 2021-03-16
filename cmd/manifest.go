/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

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
	"archive/tar"
	"errors"
	"fmt"
	"io"
	"os"

	"io/ioutil"

	"github.com/containers/image/v5/manifest"
	"github.com/spf13/cobra"
)

// manifestIdCmd represents the manifestId command
var manifestIdCmd = &cobra.Command{
	Use:   "manifest-id",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		manifestPath := args[0]
		err := getManifestId(manifestPath)
		if err != nil {
			fmt.Printf("Error getting manifest id from tar: %e", err)
		}

	},
}

func getManifestId(tarPath string) error {
	tarFile, err := os.Open(tarPath)
	if err != nil {
		return err
	}
	defer tarFile.Close()
	tarReader := tar.NewReader(tarFile)
	for {
		header, err := tarReader.Next()

		if err == io.EOF {
			break
		}

		if err != nil {
			return errors.New("error getting files from tar")
		}

		name := header.Name
		if name == "manifest.json" {
			man, _ := ioutil.ReadAll(tarReader)
			digest, err := manifest.Digest(man)
			if err != nil {
				fmt.Printf("Error computing digest: %v", err)
			}
			fmt.Printf("%s\n", digest)
		}

	}
	return nil
}

func init() {
	rootCmd.AddCommand(manifestIdCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// manifestIdCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// manifestIdCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
