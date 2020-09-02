package scan

import (
	"fmt"
	"os"
	"io/ioutil"

	error "github.com/ClementBolin/gitStat-go/pkg/err"
)

// ScanFolder : scan folder for find .git Path is folder to scan
func ScanFolder(path string) {
	file, err := ioutil.ReadDir(path)
	error.MangeErrExit(err)

	for _, f := range file {
		if (f.Name() == ".git") {
			fmt.Println("It's .git file")
		}
		fmt.Println(f.Name())
	}
}

// CreateGitStatFile : create .gitStat-go
func CreateGitStatFile() {
	file, err := os.Create("~/.gitStat-go")
	defer file.Close()
	error.MangeErr(err)
}
