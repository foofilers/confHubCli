// Copyright © 2017 Igor Maculan <n3wtron@gmail.com>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"os"
	"github.com/Sirupsen/logrus"
	"github.com/foofilers/confHubCli/confWriters"
	"strings"
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.UsageString())
	},
}

var formats = []string{"json", "structJson", "structXml","xml", "properties"}

var getCmd = &cobra.Command{
	Use: "get [appName]",
	Short:"Get application configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println(cmd.UsageString())
			os.Exit(-1)
		}
		cl := GetClient(cmd)
		version := cmd.Flag("version").Value.String()
		format := cmd.Flag("format").Value.String()
		configs, err := cl.GetConfigs(args[0], version)
		if err != nil {
			logrus.Fatal(err)
		}
		switch format {

		case "json":
			jsonOut, err := confWriters.ConfToJson(configs, cmd.Flag("pretty").Value.String() == "true")
			if err != nil {
				logrus.Fatal(err)
			}
			fmt.Println(string(jsonOut))
		case "structJson":
			jsonOut, err := confWriters.ConfToStructuredJson(configs, cmd.Flag("pretty").Value.String() == "true")
			if err != nil {
				logrus.Fatal(err)
			}
			fmt.Println(string(jsonOut))
		case "xml":
			jsonOut, err := confWriters.ConfToStructuredXml(configs, cmd.Flag("pretty").Value.String() == "true")
			if err != nil {
				logrus.Fatal(err)
			}
			fmt.Println(string(jsonOut))
		case "properties":
			for p, v := range configs {
				fmt.Printf("%s=%s\n", p, v)
			}
		default:
			logrus.Fatal("invalid format")
		}
	},
}
var putCmd = &cobra.Command{
	Use: "put [appName] [key] [value]",
	Short:"Put application configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			fmt.Println(cmd.UsageString())
			os.Exit(-1)
		}
		version := cmd.Flag("version").Value.String()
		cl := GetClient(cmd)
		err := cl.SetValue(args[0], version, args[1], args[2])
		if err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	getCmd.Flags().String("version", "", "Application version (default: currentVersion) ")
	getCmd.Flags().StringP("format", "f", "json", "Output format " + strings.Join(formats, ", "))
	getCmd.Flags().Bool("pretty", false, "Pretty format")
	configCmd.AddCommand(getCmd)
	putCmd.Flags().String("version", "", "Application version (default: currentVersion) ")
	configCmd.AddCommand(putCmd)
	RootCmd.AddCommand(configCmd)

}
