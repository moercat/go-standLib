package main

import (
	"fmt"
)

func coinChange(nums []int) int {
	dp := make(map[int]int)

	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])
	dp[2] = max(nums[0]+nums[2], nums[1])
	for i := 3; i < len(nums); i++ {
		dp[i] = max(dp[i-1], dp[i-2]+nums[i])
	}

	return dp[len(nums)-1]
}

func str(s string) {
	dp := make([][]int, len(s))

	for i := range dp {
		dp[i] = make([]int, len(s))
	}

	for i := 0; i < len(s); i++ {
		dp[i][i] = 1
	}

	for i := 1; i < len(s); i++ {
		for j := i; j < len(s); j++ {

		}
	}

}

// 测试
func main() {
	fmt.Println(coinChange([]int{1, 2, 5, 4, 5, 2, 3, 4, 5}))
}
