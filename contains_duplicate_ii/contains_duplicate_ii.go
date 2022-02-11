package main

import "fmt"

func myAbs(number int) int {
	if number < 0 {
		return -number
	}
	return number
}

func containsNearbyDuplicate(nums []int, k int) bool {
	var hash = map[int]int{}

	for i, num := range nums {
		if position, ok := hash[num]; ok && position != i && myAbs(i-position) <= k {
			return true
		} else {
			hash[num] = i
		}
	}

	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	k := 3

	fmt.Println(containsNearbyDuplicate(nums, k))
}
