// passdrill: typing drills for practicing passphrases

package main

import (
	"fmt"
	"os"
	"strings"
)

const hashAlgorithm = "sha512"
const passphraseHashFilename = "passdrill." + hashAlgorithm
const help = "Use -s to save passphrase hash for practice."

func input(msg string) string {
	response := ""
	fmt.Print(msg)
	fmt.Scan(&response)
	return response
}

func prompt() string {
	fmt.Println("WARNING: the passphrase will be shown so that you can check it!")
	confirmed := ""
	passwd := ""
	for confirmed != "y" {
		passwd = input("Type passphrase to hash (it will be echoed): ")
		if passwd == "" || passwd == "q" {
			fmt.Println("ERROR: the passphrase cannot be empty or 'q'.")
			continue
		}
		fmt.Println("Passphrase to be hashed ->", passwd)
		confirmed = strings.ToLower(input("Confirm (y/n): "))
	}
	return passwd
}

func save_hash(args []string) {
	prompt()
}

func practice() {
	fmt.Println("Practice...")
}

func main() {
	if len(os.Args) > 1 {
		save_hash(os.Args)
	} else {
		practice()
	}
}
