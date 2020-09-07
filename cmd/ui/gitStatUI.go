package ui

import (
	"fmt"

	"github.com/ClementBolin/gitStat-go/pkg/git"
)

const (
	daySixMonths = 183
	weekSixMonths = 26
)

func createTable(arrayPath []string, email string) {
	commitsTable := make(map[int]int, daySixMonths)

	for i := daySixMonths; i > 0; i-- {
		commitsTable[i] = 0
	}
	for _, e := range arrayPath {
		git.CountCommit(e, commitsTable, email)
	}
	for e := range commitsTable {
		fmt.Println("day :", e, "commits number :", commitsTable[e])
	}
}

// DisplayUI : display commit in your terminal
func DisplayUI(arrayPath []string, email string) {
	createTable(arrayPath, email)
}
