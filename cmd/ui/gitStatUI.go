package ui

import (
	"fmt"
	"time"
	"github.com/ClementBolin/gitStat-go/cmd/scan"
)

const (
	daySixMonths = 183
	weekSixMonths = 26
)

func getActualDate(t time.Time) time.Time {
	year, month, day := t.Date();
	actualDay := time.Date(year, month, day, 0, 0, 0, 0, t.Location())
	fmt.Println("Location :", t.Location())
	return actualDay
}

func createTable() {
	commitsTable := make(map[int]int, daySixMonths)
	for i := daySixMonths; i > 0; i-- {
		commitsTable[i] = 0
	}
}

// DisplayUI : display commit in your terminal
func DisplayUI(gitScan scan.GitScan) {
	fmt.Println("total Commit :", gitScan.GetCounter() , "actual day :", getActualDate(time.Now()))
}
