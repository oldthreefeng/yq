/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

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

	"github.com/spf13/cobra"
)

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		VersionStr()
	},
}

func init() {
	rootCmd.AddCommand(versionCmd)
}

var (
	Buildstamp = ""
	Githash    = ""
	Version    = "latest"
	Author     = ""
)

func VersionStr() {
	fmt.Printf(`yq is cli tools to trans yaml to json ; or json to yaml
run "yq -h" get more help, more see https://github.com/oldthreefeng/yq
`)
	fmt.Printf("yq version :       %s\n", Version)
	fmt.Printf("Git Commit Hash:     %s\n", Githash)
	fmt.Printf("Build Time :         %s\n", Buildstamp)
	fmt.Printf("BuildBy :            %s\n", Author)
}
