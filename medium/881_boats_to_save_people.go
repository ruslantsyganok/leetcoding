package main

import (
	"fmt"
	"math"
	"sort"
)

// You are given an array people where people[i] is the weight
// of the ith person, and an infinite number of boats where
// each boat can carry a maximum weight of limit. Each boat
// carries at most two people at the same time, provided the
// sum of the weight of those people is at most limit.
//
// Return the minimum number of boats to carry every given person.

func main() {
	people := []int{8, 2, 3, 6, 2, 6} // -> 4
	limit := 8

	fmt.Println(numRescueBoats(people, limit))
}

// Faster than 80.08%
func numRescueBoats(people []int, limit int) int {
	var boats int
	sort.Ints(people)
	length := len(people)

	for {
		if length == 0 {
			break
		}
		first, last := people[0], people[len(people)-1]
		if first+last > limit || length == 1 {
			people = people[:len(people)-1]
			length -= 1
			boats++
		} else {
			people = people[1 : len(people)-1]
			length -= 2
			boats++
		}
	}

	return boats
}

// faster than 95.02%
func test(people []int, limit int) int {
	sort.Ints(people)
	first, last, boats := 0, len(people)-1, 0

	for first <= last {
		if first == last {
			boats++
			return boats
		}
		if people[first]+people[last] <= limit {
			first++
			last--
		} else {
			last--
		}
		boats++
	}
	return boats
}

// Very slow solution
func numRescueBoatsSlow(people []int, limit int) int {
	var boats int

	for {
		if len(people) == 0 {
			break
		}

		minWeight, idx1 := math.MaxInt, 0
		maxWeight, idx2 := 0, 0

		for idx, weight := range people {
			if weight < minWeight {
				minWeight = weight
				idx2 = idx
			}
			if weight > maxWeight {
				maxWeight = weight
				idx1 = idx
			}
		}

		if minWeight+maxWeight > limit || len(people) == 1 {
			boats++
			people = removeOne(people, idx1)
			continue
		}
		boats++
		people = removeTwo(people, idx1, idx2)
	}

	return boats
}

func removeOne(people []int, idx int) []int {
	if len(people) > 1 {
		people = append(people[:idx], people[idx+1:]...)
		return people
	}
	return nil
}

func removeTwo(people []int, idx1, idx2 int) []int {
	if idx1 > idx2 {
		idx1, idx2 = idx2, idx1
	}
	if len(people) >= 2 {
		people = append(people[:idx2], people[idx2+1:]...)
		people = append(people[:idx1], people[idx1+1:]...)
		return people
	}
	return nil
}
