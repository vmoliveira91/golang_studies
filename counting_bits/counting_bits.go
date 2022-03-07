package main

import (
	"fmt"
	"strconv"
	"strings"
)

func countBits(n int) []int {
	var ans []int

	for i := 0; i <= n; i++ {
		numberInBinary := strconv.FormatInt(int64(i), 2)
		ans = append(ans, strings.Count(string(numberInBinary), "1"))
	}

	return ans
}

func main() {
	fmt.Println(countBits(5))
}
