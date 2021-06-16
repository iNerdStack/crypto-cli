package main

import (
	hashtypes "github.com/inerdstack/crypto-cli/hashtypes"

	"flag"
	"fmt"
	"os"
)

var (
	HelpMessage string = `expected a hashtype arg 'md4', 'md5', 'sha1','sha512':
	Example: crypto-cli md5 --text="word";
  `
)

func main() {

	md5Cmd := flag.NewFlagSet("md5", flag.ExitOnError)
	md5Name := md5Cmd.String("text", "", "a string")

	md4Cmd := flag.NewFlagSet("md4", flag.ExitOnError)
	md4Name := md4Cmd.String("text", "", "a string")

	sha1Cmd := flag.NewFlagSet("sha1", flag.ExitOnError)
	sha1Name := sha1Cmd.String("text", "", "a string")

	sha512Cmd := flag.NewFlagSet("sha512", flag.ExitOnError)
	sha512Name := sha512Cmd.String("text", "", "a string")

	if len(os.Args) < 2 {
		fmt.Println(HelpMessage)
		os.Exit(1)
	}

	switch os.Args[1] {
	case "md4":
		md4Cmd.Parse(os.Args[2:])
		fmt.Println(hashtypes.MD4Hash(*md4Name))
	case "md5":
		md5Cmd.Parse(os.Args[2:])
		fmt.Println(hashtypes.MD5Hash(*md5Name))
	case "sha1":
		sha1Cmd.Parse(os.Args[2:])
		fmt.Println(hashtypes.SHA1Hash(*sha1Name))
	case "sha512":
		sha512Cmd.Parse(os.Args[2:])
		fmt.Println(hashtypes.SHA512Hash(*sha512Name))
	default:
		fmt.Println(HelpMessage)
		os.Exit(1)
	}
}
