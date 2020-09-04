package scan

import (
	"os"
	"io/ioutil"

	error "github.com/ClementBolin/gitStat-go/pkg/err"
)

// ScanFolder : scan folder for find .git Path is folder to scan
func ScanFolder(path string, email string, gitScan *GitScan) {
	// Read File
	file, err := ioutil.ReadDir(path)
	error.MangeErrExit(err)

	var newPath string

	for _, f := range file {
		if (f.Name() == ".git") {
			gitScan.GitScanCommit(path + "/" + f.Name())
		}
		newPath = path + "/" + f.Name()
		if fNewPath, err := os.Stat(newPath); err == nil && fNewPath.IsDir() {
			ScanFolder(newPath, email, gitScan)
		}
	}
}
