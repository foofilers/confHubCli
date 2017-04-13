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
	"bufio"
	"strings"
	"encoding/json"
)

// applicationCmd represents the application command
var applicationCmd = &cobra.Command{
	Use:   "application",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(cmd.UsageString())
	},
}

var createCmd = &cobra.Command{
	Use:   "create [appName]",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			os.Exit(-1)
		}
		cli := GetClient(cmd)
		if err := cli.AddApplication(args[0]); err != nil {
			logrus.Fatalf("Error adding application :%s", err)
		}
	},
}

var renameCmd = &cobra.Command{
	Use:   "rename [appName] [newAppName]",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			cmd.Usage()
			os.Exit(-1)
		}
		cli := GetClient(cmd)
		if err := cli.RenameApplication(args[0], args[1]); err != nil {
			logrus.Fatalf("Error renaming application: %s", err)
		}
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Run: func(cmd *cobra.Command, args []string) {
		cli := GetClient(cmd)
		apps, err := cli.ListApplication()
		if err != nil {
			logrus.Fatalf("Error retrieving application: %s", err)
		}
		fmt.Println("[")
		for _, app := range apps {
			out, err := json.Marshal(app)
			if err != nil {
				logrus.Fatalf("Error retrieving application: %s", err)
			}
			fmt.Printf("\t%+s,\n", out)
		}
		fmt.Println("]")
	},
}

var deleteCmd = &cobra.Command{
	Use:   "delete [appName]",
	Short: "A brief description of your command",

	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			os.Exit(-1)
		}
		appName := args[0]
		cli := GetClient(cmd)
		var answer string
		var err error
		if cmd.Flag("yes").Value.String() != "true" {
			reader := bufio.NewReader(os.Stdin)
			fmt.Print("Are you sure to delete the application " + appName + "?[y/N]: ")
			answer, err = reader.ReadString('\n')
			if err != nil {
				logrus.Fatal(err)
			}
		} else {
			answer = "y"
		}
		if strings.TrimSpace(answer) == "y" || strings.TrimSpace(answer) == "Y" {
			if err := cli.DeleteApplication(appName); err != nil {
				logrus.Fatalf("Error deleting application :%s", err)
			}
			logrus.Info("Application deleted")
		} else {
			logrus.Info("Operation canceled")
		}
	},
}

var watchAppCmd = &cobra.Command{
	Use:   "watch [appName]",
	Short: "Watch application changes",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			cmd.Usage()
			os.Exit(-1)
		}
		cli := GetClient(cmd)
		watchCh, err := cli.WatchApp([]string{args[0]});
		if err != nil {
			logrus.Fatalf("Error adding application :%s", err)
		}
		for ch := range watchCh {
			fmt.Printf("%+v\n", ch)
		}
	},
}

func init() {
	applicationCmd.AddCommand(listCmd)
	applicationCmd.AddCommand(createCmd)
	applicationCmd.AddCommand(renameCmd)

	deleteCmd.Flags().BoolP("yes", "y", false, "Delete without ask")
	applicationCmd.AddCommand(deleteCmd)

	applicationCmd.AddCommand(watchAppCmd)


	RootCmd.AddCommand(applicationCmd)

}
