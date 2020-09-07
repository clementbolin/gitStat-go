package scan

import (
	"os"
	"io/ioutil"

	error "github.com/ClementBolin/gitStat-go/pkg/err"
)

// scanFolder : scan recursive folder for find .git Path is folder to scan
func scanFolder(path string, email string, arrayPath []string, line *int) {
	// Read File
	file, err := ioutil.ReadDir(path)
	error.MangeErrExit(err)

	var newPath string

	for _, f := range file {
		if (f.Name() == ".git") {
			arrayPath[*line] = path + "/" + f.Name()
			*line++
			// continue
		}
		newPath = path + "/" + f.Name()
		if fNewPath, err := os.Stat(newPath); err == nil && fNewPath.IsDir() {
			scanFolder(newPath, email, arrayPath, line)
		}
	}
}

func counterPathFolder(path string, email string, counter *int) {
	// Read folder
	file, err := ioutil.ReadDir(path)
	error.MangeErrExit(err)

	var newPath string

	for _, f := range file {
		if (f.Name() == ".git") {
			*counter++
		}
		newPath = path + "/" + f.Name()
		if fNewPath, err := os.Stat(newPath); err == nil && fNewPath.IsDir() {
			counterPathFolder(newPath, email, counter)
		}
	}
}

// CreatePathFolder : create an array where we save all .git path
func CreatePathFolder(path string, email string) []string {
	// Init array Path
	var counter int = 0
	counterPathFolder(path, email, &counter)

	arrayPath := make([]string, counter)
	counter = 0
	scanFolder(path, email, arrayPath, &counter)
	return arrayPath
}

// ScanUniqueFolder : Scan folder not recursive
func ScanUniqueFolder(path string, email string) string {
	// Read File
	file, err := ioutil.ReadDir(path)
	error.MangeErrExit(err)

	// String path
	var pathReturn string

	for _, f := range file {
		if (f.Name() == ".git") {
			pathReturn = path + "/" + f.Name()
		}
	}
	return pathReturn
}
