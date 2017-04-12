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

// versionCmd represents the version command
var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.UsageString())
	},
}
var setVersionCmd = &cobra.Command{
	Use:   "set [appName] [appVersion]",
	Short: "Set a default application version",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println(cmd.UsageString())
			os.Exit(-1)
		}
		cl := GetClient(cmd)
		if err := cl.SetDefaultVersion(args[0], args[1]); err != nil {
			logrus.Fatal(err)
		}
	},
}

var createVersionCmd = &cobra.Command{
	Use:   "create [appName] [appVersion]",
	Short: "Create a new app version",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Println(cmd.UsageString())
			os.Exit(-1)
		}
		cl := GetClient(cmd)
		if err := cl.AddVersion(args[0], args[1]); err != nil {
			logrus.Fatal(err)
		}
	},
}

var copyVersionCmd = &cobra.Command{
	Use:   "copy [appName] [srcVersion] [dstVersion]",
	Short: "Create an application version",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			fmt.Println(cmd.UsageString())
			os.Exit(-1)
		}
		cl := GetClient(cmd)
		if err := cl.CopyVersion(args[0], args[1], args[2]); err != nil {
			logrus.Fatal(err)
		}
	},
}

var listVersionCmd = &cobra.Command{
	Use:   "list [appName] ",
	Short: "List application versions",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Println(cmd.UsageString())
			os.Exit(-1)
		}
		cl := GetClient(cmd)
		versions, err := cl.GetVersions(args[0]);
		if err != nil {
			logrus.Fatal(err)
		}
		for _, v := range versions {
			fmt.Println(v)
		}
	},
}

func init() {
	versionCmd.AddCommand(createVersionCmd)
	versionCmd.AddCommand(setVersionCmd)
	versionCmd.AddCommand(copyVersionCmd)
	versionCmd.AddCommand(listVersionCmd)

	RootCmd.AddCommand(versionCmd)

}
