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

	"crypto/sha1"
	"encoding/hex"
	"github.com/spf13/cobra"
)

var (
	sha1text string
)

// sha1Cmd represents the hash command
var sha1Cmd = &cobra.Command{
	Use:   "sha1",
	Short: "Generate sha1 hash",
	Long: `Generate sha1 Hash from text
	
	For example:
	crypto-cli sha1 text="password"
	`,
	Run: func(cmd *cobra.Command, args []string) {

		hash := sha1.New()
		hash.Write([]byte(sha1text))
		sha1_hash := hex.EncodeToString(hash.Sum(nil))

		fmt.Println(sha1_hash)
	},
}

func init() {
	rootCmd.AddCommand(sha1Cmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//sha1Cmd.PersistentFlags().String("word", "", "Word to hash")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	//sha1Cmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	sha1Cmd.Flags().StringVar(&sha1text, "text", "", "text to hash")
	sha1Cmd.MarkFlagRequired("text")
}
