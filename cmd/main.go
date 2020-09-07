package main

import (
	"fmt"
	"flag"
	"os"

	"github.com/ClementBolin/gitStat-go/cmd/scan"
	"github.com/ClementBolin/gitStat-go/cmd/ui"
)

// HELP : array of help message
var HELP = [3]string {
	"Add folder path to scan for Git repository, if not specified, scan current repository",
	"Email to scan in Git commit",
	"scans all folders recursively and adds up all commits",
}

func help() {
	fmt.Println(`gitStat-go help:

   -add		path folder to scan for Git repository, if not specified, scan current repository.
   -email	email to scan in Git commit, required parameter
   -r	scans all folders recursively and adds up all commits`)
}

// Deleate char at indx
func delChar(s []rune, index int) []rune {
    return append(s[0:index], s[index+1:]...)
}

// clear path string
func clearPath(path string) string {
	for i := 0; i != len(path)-2; i++ {
		c := string(path[i])
		c1 := string(path[i+1])
		if (c == "/" && c1 == "/") {
			path = string(delChar([]rune(path), i))
		}
	}
	return path
}

func main() {
	var folderFlag string
	var emailFlag string
	var recursive *bool

	// Check flag
	recursive = flag.Bool("r", false, HELP[2]);
	flag.StringVar(&folderFlag, "add", ".", HELP[0]);
	flag.StringVar(&emailFlag, "email", "example@email.com", HELP[1])
	flag.Parse()

	if (emailFlag == "example@email.com") {
		help()
		os.Exit(0)
	}

	if (*recursive) {
		arrayPath := scan.CreatePathFolder(folderFlag, emailFlag)
		if (len(arrayPath) <= 0) {
			fmt.Printf("none commit find for %s\n\n", emailFlag)
			help()
			os.Exit(0)
		}
		for i, e := range arrayPath {
			arrayPath[i] = clearPath(e) 
		}
		ui.DisplayUI(arrayPath, emailFlag)
		os.Exit(0)
	} else {
		path := scan.ScanUniqueFolder(folderFlag, emailFlag)
		if (path[0] == "") {
			fmt.Printf("none commit find for %s\n\n", emailFlag)
			help()
			os.Exit(0)
		}
		path[0] = clearPath(path[0])
		ui.DisplayUI(path, emailFlag)
		os.Exit(0)
	}
}
	