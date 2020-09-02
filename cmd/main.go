package main

import (
	"fmt"
	"flag"
	"os"

	"github.com/ClementBolin/gitStat-go/cmd/scan"
)

// HELP : array of help message
var HELP = [2]string {
	"<foler path>\tAdd folder path to scan for Git repository",
	"<email@email.com>\tEmail to scan in Git commit",
}

func help() {
	fmt.Println(`gitStat-go help:

   --add	path folder to scan for Git repository, if not specified, scan current repository.
   --email	email to scan in Git commit, required parameter`)
}

func main() {
	var folderFlag string
	var emailFlag string
	var debugFlag string

	// Check flag
	flag.StringVar(&debugFlag, "debug", "false", HELP[0]);
	flag.StringVar(&folderFlag, "add", ".", HELP[0]);
	flag.StringVar(&emailFlag, "email", "example@email.com", HELP[1])
	flag.Parse()

	if (debugFlag == "false") {
		if (emailFlag == "example@email.com") {
			help()
			os.Exit(0)
		}
	}
	scan.ScanFolder(folderFlag)
}
	