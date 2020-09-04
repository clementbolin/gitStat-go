package scan

import (
	"fmt"
	"os"
	"io/ioutil"

	error "github.com/ClementBolin/gitStat-go/pkg/err"
)

// ScanFolder : scan folder for find .git Path is folder to scan
func ScanFolder(path string, email string) {
	// Read File
	file, err := ioutil.ReadDir(path)
	error.MangeErrExit(err)

	var gitScan GitScan
	gitScan.Init(email)

	for _, f := range file {
		if (f.Name() == ".git") {
			gitScan.GitScanCommit(path + "/" + f.Name())
		}
		fmt.Println(f.Name())
	}
	fmt.Println("commit number : ", gitScan.GetCounter())
}

// ScanRecursiveFolder : Scan recursive folder
// func ScanRecursiveFolder(path string, email string) {
// 	var gitScan GitScan
// 	gitScan.Init(email)

// 	_, err := os.Stat(path)
// 	error.MangeErrExit(err)
// }

// CreateGitStatFile : create .gitStat-go
func CreateGitStatFile() {
	file, err := os.Create("~/.gitStat-go")
	defer file.Close()
	error.MangeErr(err)
}
