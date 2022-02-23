package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	s := strconv.Itoa(x)

	i, j := 0, len(s)-1

	for i < j {
		if s[i] != s[j] {
			return false
		}
		i++
		j--
	}
	return true
}

func main() {
	fmt.Println(isPalindrome(-121))
}
