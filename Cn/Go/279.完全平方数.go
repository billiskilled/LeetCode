package main

import (
	"math"
)

/*
 * @lc app=leetcode.cn id=279 lang=golang
 *
 * [279] 完全平方数
 */

/**
状态定义      : dp[i]：到i时需要几步  re[i]
状态转移方程  : dp[i] = 1 + min(dp[n - 1^2], dp[n - 2^2], dp[n - 3^2]... dp[n - k^2]) k表示最大的满足k^2<n的整数
*/
/**
是否可用贪心算法 ?
答案：不可以。
举例：12  贪心为：9 + 1 + 1 + 1  最优为 4 + 4 + 4
*/
// func main() {
// 	var n = 12
// 	fmt.Println(numSquares(n))
// }

func numSquares(n int) int {
	dp := []int{1000000000, 1, 2, 3}
	if n < 3 {
		return n
	}

	for i := 4; i <= n; i++ {
		dp = append(dp, findMin(i, dp))
	}

	return dp[n]
}

func findMin(n int, dp []int) int {
	var maxSqrtV = int(math.Sqrt(float64(n)))

	if maxSqrtV*maxSqrtV == n {
		return 1
	}

	var tempArr []int
	for i := 1; i <= maxSqrtV; i++ {
		tempArr = append(tempArr, dp[n-i*i])
	}

	return 1 + getMin(tempArr)
}

func getMin(arr []int) int {
	var min = math.MaxInt32
	for _, v := range arr {
		if v < min {
			min = v
		}
	}

	return min
}
