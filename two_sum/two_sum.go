package main

import "fmt"

func twoSum(nums []int, target int) []int {
	hashNums := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if position, ok := hashNums[target-nums[i]]; ok {
			return []int{position, i}
		} else {
			hashNums[nums[i]] = i
		}
	}

	return nil
}

func main() {
	nums := []int{3, 3}
	target := 6

	fmt.Println(twoSum(nums, target))
}
