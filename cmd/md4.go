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

	"encoding/hex"
	"github.com/spf13/cobra"
	"golang.org/x/crypto/md4"
)

var (
	md4text string
)

// md4Cmd represents the hash command
var md4Cmd = &cobra.Command{
	Use:   "md4",
	Short: "Generate md4 hash",
	Long: `Generate md4 Hash from text
	
	For example:
	crypto-cli md4 text="password"
	`,
	Run: func(cmd *cobra.Command, args []string) {

		hash := md4.New()
		hash.Write([]byte(md4text))
		md4_hash := hex.EncodeToString(hash.Sum(nil))

		fmt.Println(md4_hash)
	},
}

func init() {
	rootCmd.AddCommand(md4Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//md4Cmd.PersistentFlags().String("word", "", "Word to hash")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//md4Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	md4Cmd.Flags().StringVar(&md4text, "text", "", "text to hash")
	md4Cmd.MarkFlagRequired("text")
}
