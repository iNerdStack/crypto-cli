/*
Copyright © 2021 NERD STACK <isaacajetunmobi@gmail.com>

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

	"crypto/md5"
	"encoding/hex"
	"github.com/spf13/cobra"
)

var (
	md5text string
)

// md5Cmd represents the hash command
var md5Cmd = &cobra.Command{
	Use:   "md5",
	Short: "Generate md5 hash",
	Long: `Generate md5 Hash from text
	
	For example:
	crypto-cli md5 text="password"
	`,
	Run: func(cmd *cobra.Command, args []string) {

		hash := md5.New()
		hash.Write([]byte(md5text))
		md5_hash := hex.EncodeToString(hash.Sum(nil))

		fmt.Println(md5_hash)
	},
}

func init() {
	rootCmd.AddCommand(md5Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//md5Cmd.PersistentFlags().String("word", "", "Word to hash")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//md5Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	md5Cmd.Flags().StringVar(&md5text, "text", "", "text to hash")
	md5Cmd.MarkFlagRequired("text")
}
