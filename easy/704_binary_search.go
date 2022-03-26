package main

import "fmt"

func main() {
	nums := []int{-1, 0, 3, 5, 9, 12, 14}
	target := -1
	fmt.Println(search(nums, target)) // -> 2
}

func search(nums []int, target int) int {
	low, mid, high := 0, len(nums)/2, len(nums)-1
	for low <= high {
		mid = (high + low) / 2
		if nums[mid] == target {
			return mid
		}
		if nums[mid] > target {
			high = mid - 1
		}
		if nums[mid] < target {
			low = mid + 1
		}
	}
	return -1
}
