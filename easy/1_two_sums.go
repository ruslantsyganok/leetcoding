package main

import "fmt"

// Given an array of integers nums and an integer target, return indices of the two numbers such that they add up to target.
// You may assume that each input would have exactly one solution, and you may not use the same element twice.
// You can return the answer in any order.

// Example: input -> slice: []int64{2,7,11,15}, target -> 6      output -> [0,1]
// Example: [3, 3] -> [0, 1]

func main() {
	nums := []int{1, 1, 1, 1, 1, 4, 1, 1, 1, 1, 1, 7, 1, 1, 1, 1, 1}
	target := 11
	fmt.Println(twoSums(nums, target))
}

func twoSums(nums []int, target int) []int {
	numMap := make(map[int]int)
	for index, num := range nums {
		if idx, found := numMap[target-num]; found {
			return []int{index, idx}
		}
		numMap[num] = index
	}
	return []int{0, 0}
}

func slowTwoSums(nums []int, target int) []int {
	for i1, num1 := range nums {
		for i2, num2 := range nums {
			if num1+num2 == target && i1 != i2 {
				return []int{i1, i2}
			}
		}
	}
	return []int{0, 0}
}
