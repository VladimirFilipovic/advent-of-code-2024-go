package main

import (
	"advent-of-code-2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func checkReportsForSafety(fileLocation string) {
	content := utils.ReadFromFile(fileLocation)

	reports := strings.Split(string(content), "\r\n")

	fmt.Println("Reports: \n", strings.Join(reports, "\n "))

	safeReports := 0
	for _, report := range reports {
		levels := strings.Fields(report)
		fmt.Println("Levels", levels)
		var isAscendingOrderOfLevels bool
		for j, level := range levels {
			fmt.Println("Index: ", j, "Level: ", level, "Len:", len(levels))
			if j == len(levels)-1 {
				safeReports++
				break
			}
			levelValue, _ := strconv.Atoi(level)
			nextLevelValue, _ := strconv.Atoi(levels[j+1])
			if levelValue == nextLevelValue {
				break
			}
			if levelValue > nextLevelValue {
				if !isAscendingOrderOfLevels && j > 0 {
					break
				}
				isAscendingOrderOfLevels = true
				difference := levelValue - nextLevelValue
				if difference > 3 {
					break
				}
			}
			if levelValue < nextLevelValue {
				if isAscendingOrderOfLevels {
					break
				}
				isAscendingOrderOfLevels = false
				difference := nextLevelValue - levelValue
				if difference > 3 {
					break
				}
			}
		}

	}

	fmt.Println("Safe reports: ", safeReports)
}

func checkReportsForSafetyWithProblemDampener(fileLocation string) {
	content := utils.ReadFromFile(fileLocation)

	reports := strings.Split(string(content), "\r\n")

	fmt.Println("Reports: \n", strings.Join(reports, "\n "))

	safeReports := 0
	for _, report := range reports {
		levels := strings.Fields(report)
		var isLevelSafe bool
		var issueIndex int

		// 5 7 8 9 23
		// 0 1 2 3 4
		isLevelSafe, issueIndex = checkLevelsForSafety(levels)
		fmt.Println("Issue index: ", issueIndex)

		if !isLevelSafe {
			for {

				levelsDeleteProblematicOne := slices.Delete(slices.Clone(levels), issueIndex-1, issueIndex)
				isLevelSafeAfterProblematic, _ := checkLevelsForSafety(levelsDeleteProblematicOne)

				if issueIndex == len(levels) && !isLevelSafeAfterProblematic {
					break
				}

				if isLevelSafeAfterProblematic {
					isLevelSafe = true
					break
				}

				levelsDeleteNextToProlematic := slices.Delete(slices.Clone(levels), issueIndex, issueIndex+1)
				levelsSafeNextToProlematic, _ := checkLevelsForSafety(levelsDeleteNextToProlematic)

				if levelsSafeNextToProlematic {
					isLevelSafe = true
					break
				}

				issueIndex++
			}
		}

		if isLevelSafe {
			safeReports++
		}

	}

	fmt.Println("Safe reports with dampener: ", safeReports)
}

func checkLevelsForSafety(levels []string) (bool, int) {
	fmt.Println("Levels", levels)
	var isAscendingOrderOfLevels bool
	for j, level := range levels {
		fmt.Println("Index: ", j, "Level: ", level, "Len:", len(levels))
		if j == len(levels)-1 {
			return true, j
		}
		levelValue, _ := strconv.Atoi(level)
		nextLevelValue, _ := strconv.Atoi(levels[j+1])
		if levelValue == nextLevelValue {
			if j > 0 {
				return false, j
			}
			return false, j + 1
		}
		if levelValue > nextLevelValue {
			if !isAscendingOrderOfLevels && j > 0 {
				if j > 0 {
					return false, j
				}
				return false, j + 1
			}
			isAscendingOrderOfLevels = true
			difference := levelValue - nextLevelValue
			if difference > 3 {
				if j > 0 {
					return false, j
				}
				return false, j + 1
			}
		}
		if levelValue < nextLevelValue {
			if isAscendingOrderOfLevels {
				if j > 0 {
					return false, j
				}
				return false, j + 1
			}
			isAscendingOrderOfLevels = false
			difference := nextLevelValue - levelValue
			if difference > 3 {
				if j > 0 {
					return false, j
				}
				return false, j + 1
			}
		}
	}
	return false, len(levels) - 1
}

func main() {
	checkReportsForSafety("/day-2/input.txt")
	checkReportsForSafetyWithProblemDampener("/day-2/input.txt")
}
