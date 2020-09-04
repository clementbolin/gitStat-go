package main

import (
	"fmt"
	"flag"
	"os"

	"github.com/ClementBolin/gitStat-go/cmd/scan"
)

// HELP : array of help message
var HELP = [3]string {
	"Add folder path to scan for Git repository",
	"Email to scan in Git commit",
	"scans all folders recursively and adds up all commits",
}

func help() {
	fmt.Println(`gitStat-go help:

   -add	path folder to scan for Git repository, if not specified, scan current repository.
   -email	email to scan in Git commit, required parameter
   -recursive	scans all folders recursively and adds up all commits`)
}

func main() {
	var folderFlag string
	var emailFlag string
	var recursive *bool

	// Check flag
	recursive = flag.Bool("recursive", false, HELP[2]);
	flag.StringVar(&folderFlag, "add", ".", HELP[0]);
	flag.StringVar(&emailFlag, "email", "example@email.com", HELP[1])
	flag.Parse()

	if (emailFlag == "example@email.com") {
		help()
		os.Exit(0)
	}
	var gitScan scan.GitScan
	gitScan.Init(emailFlag)
	if (*recursive) {
		scan.ScanFolder(folderFlag, emailFlag, &gitScan)
		fmt.Println("commit number : ", gitScan.GetCounter())
		if (gitScan.GetCounter() <= 0) {
			fmt.Printf("none commit find for %s\n\n", emailFlag)
			help()
		}
		os.Exit(0)
	} else {
		scan.ScanUniqueFolder(folderFlag, emailFlag, &gitScan)
		fmt.Println("commit number : ", gitScan.GetCounter())
		if (gitScan.GetCounter() <= 0) {
			fmt.Printf("none commit find for %s\n\n", emailFlag)
			help()
		}
		os.Exit(0)
	}
}
	