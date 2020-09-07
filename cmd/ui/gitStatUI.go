package ui

import (
	"fmt"
	"sort"
	"time"

	"github.com/ClementBolin/gitStat-go/pkg/git"
)

const (
	daySixMonths = 183
	weekSixMonths = 26
)

type column []int

func createTable(arrayPath []string, email string) map[int]int {
	commitsTable := make(map[int]int, daySixMonths)

	for i := daySixMonths; i > 0; i-- {
		commitsTable[i] = 0
	}
	for _, e := range arrayPath {
		git.CountCommit(e, commitsTable, email)
	}
	return commitsTable
}

func printCommitsStat(commits map[int]int) {
	key := sortMapIntoSlice(commits)
	cols := buildCol(key, commits)
	printCells(cols)
}

func buildCol(keys []int, commits map[int]int) map[int]column {
	cols := make(map[int]column)
	col := column{}

	for _, k := range keys {
		week := int(k / 7) //26,25...1
		dayinweek := k % 7 // 0,1,2,3,4,5,6

		if dayinweek == 0 { //reset
			col = column{}
		}

		col = append(col, commits[k])

		if dayinweek == 6 {
			cols[week] = col
		}
	}

	return cols
}

func sortMapIntoSlice(m map[int]int) []int {
	var key []int

	for k := range m {
		key = append(key, k)
	}
	sort.Ints(key)
	return key
}

// Display UI

// printCells prints the cells of the graph
func printCells(cols map[int]column) {
	printMonths()
	for j := 6; j >= 0; j-- {
		for i := weekSixMonths + 1; i >= 0; i-- {
			if i == weekSixMonths+1 {
				printDayCol(j)
			}
			if col, ok := cols[i]; ok {
				//special case today
				if i == 0 && j == git.CalcOffset()-1 {
					printCell(col[j], true)
					continue
				} else {
					if len(col) > j {
						printCell(col[j], false)
						continue
					}
				}
			}
			printCell(0, false)
		}
		fmt.Printf("\n")
	}
}

// printMonths prints the month names in the first line, determining when the month
// changed between switching weeks
func printMonths() {
	week := git.GetActualDate(time.Now()).Add(-(daySixMonths * time.Hour * 24))
	month := week.Month()
	fmt.Printf("         ")
	for {
		if week.Month() != month {
			fmt.Printf("%s ", week.Month().String()[:3])
			month = week.Month()
		} else {
			fmt.Printf("    ")
		}

		week = week.Add(7 * time.Hour * 24)
		if week.After(time.Now()) {
			break
		}
	}
	fmt.Printf("\n")
}

// printDayCol given the day number (0 is Sunday) prints the day name,
// alternating the rows (prints just 2,4,6)
func printDayCol(day int) {
	out := "     "
	switch day {
	case 1:
		out = " Mon "
	case 3:
		out = " Wed "
	case 5:
		out = " Fri "
	}

	fmt.Printf(out)
}

func printCell(val int, today bool) {
	escape := "\033[0;37;30m"
	switch {
	case val > 0 && val < 5:
		escape = "\033[1;30;47m"
	case val >= 5 && val < 10:
		escape = "\033[1;30;43m"
	case val >= 10:
		escape = "\033[1;30;42m"
	}

	if today {
		escape = "\033[1;37;45m"
	}

	if val == 0 {
		fmt.Printf(escape + "  - " + "\033[0m")
		return
	}

	str := "  %d "
	switch {
	case val >= 10:
		str = " %d "
	case val >= 100:
		str = "%d "
	}

	fmt.Printf(escape+str+"\033[0m", val)
}

// DisplayUI : display commit in your terminal
func DisplayUI(arrayPath []string, email string) {
	commits := createTable(arrayPath, email)
	printCommitsStat(commits)
}
