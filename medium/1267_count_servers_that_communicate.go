package main

import "fmt"

func main() {
	grid := [][]int{
		{1, 1, 0, 0, 0},
		{0, 0, 1, 1, 0},
		{0, 0, 0, 0, 0},
		{1, 0, 0, 1, 0},
	}

	//grid := [][]int{
	//	{1, 0},
	//	{0, 1},
	//}

	//grid := [][]int{
	//	{1, 0},
	//	{1, 1},
	//}

	fmt.Println(countServers(grid))
}

// faster than 71.43%
func countServers(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	computers, rowNums, columnNums := 0, len(grid), len(grid[0])
	rowCount, columnCount := make([]int, rowNums), make([]int, columnNums)

	for i := 0; i < rowNums; i++ {
		for j := 0; j < columnNums; j++ {
			if grid[i][j] == 1 {
				computers++
				rowCount[i]++
				columnCount[j]++
			}
		}
	}
	for i := 0; i < rowNums; i++ {
		for j := 0; j < columnNums; j++ {
			if grid[i][j] == 1 && rowCount[i] == 1 && columnCount[j] == 1 {
				computers--
			}
		}
	}
	return computers
}
