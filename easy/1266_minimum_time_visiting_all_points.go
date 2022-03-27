package main

import (
	"fmt"
	"math"
)

func main() {
	points := [][]int{{1, 1}, {3, 4}, {-1, 0}}
	//points := [][]int{{3, 2}, {-2, 2}}
	fmt.Println(minTimeToVisitAllPoints(points))
}

// faster than 93.75%
func minTimeToVisitAllPoints(points [][]int) int {
	var seconds int
	for i := 0; i <= len(points)-2; i++ {
		now, next := points[i], points[i+1]

		temp1 := math.Abs(float64(now[0] - next[0]))
		temp2 := math.Abs(float64(now[1] - next[1]))

		if temp1 > temp2 {
			seconds += int(temp1)
		} else {
			seconds += int(temp2)
		}
	}
	return seconds
}
