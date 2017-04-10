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
)

var configCmd = &cobra.Command{
	Use:   "config",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.UsageString())
	},
}

var getCmd = &cobra.Command{
	Use: "get [appName] [appVersion]",
	Short:"Get application configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println(cmd.UsageString())
			os.Exit(-1)
		}
		cl := GetClient(cmd)
		res, err := cl.GetConfigs(args[0], args[1], cmd.Flag("format").Value.String())
		if err != nil {
			logrus.Fatal(err)
		}

		fmt.Println(res)
	},
}
var putCmd = &cobra.Command{
	Use: "put [appName] [appVersion] [key] [value]",
	Short:"Put application configuration",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 4 {
			fmt.Println(cmd.UsageString())
			os.Exit(-1)
		}
		cl := GetClient(cmd)
		err := cl.SetValue(args[0], args[1], args[2], args[3])
		if err != nil {
			logrus.Fatal(err)
		}
	},
}

func init() {
	getCmd.Flags().StringP("format", "f", "json", "Output format [json,xml,properties]")
	configCmd.AddCommand(getCmd)
	configCmd.AddCommand(putCmd)
	RootCmd.AddCommand(configCmd)

}
