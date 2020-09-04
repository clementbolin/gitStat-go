package scan

import (
	git "github.com/ClementBolin/gitStat-go/pkg/git"
)

// GitScan : Git Structure
type GitScan struct {
	email string
	counter int
}

// Init : init GitScan structure with an email
func (gitScan *GitScan) Init(email string) {
	gitScan.counter = 0
	gitScan.email = email
}

// GitScanCommit : scan and increment counter in your commit
func (gitScan *GitScan) GitScanCommit(path string) {
	git.CountCommit(path, &gitScan.counter, gitScan.email)
}

// GetCounter : Get counter variable
func (gitScan GitScan) GetCounter() int {
	return gitScan.counter;
}
