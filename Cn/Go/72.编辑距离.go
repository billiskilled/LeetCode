package main

/*
 * @lc app=leetcode.cn id=72 lang=golang
 *
 * [72] 编辑距离
 */
func minDistance(word1 string, word2 string) int {
	var len1 = len(word1)
	var len2 = len(word2)

	var dp = make([][]int, len1+1)
	for i := 0; i <= len1; i++ {
		dp[i] = make([]int, len2+1)
	}

	for i := 0; i <= len1; i++ {
		dp[i][0] = i
	}
	for j := 0; j <= len2; j++ {
		dp[0][j] = j
	}

	for i := 1; i <= len1; i++ {
		for j := 1; j <= len2; j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = dp[i-1][j-1]
			} else {
				dp[i][j] = getMin(dp[i-1][j], dp[i][j-1], dp[i-1][j-1]) + 1
			}
		}
	}
	// fmt.Printf("%v", dp)
	return dp[len1][len2]
}

func getMin(x int, y int, z int) int {
	if y < x {
		x = y
	}
	if z < x {
		x = z
	}
	return x
}

// func main() {
// 	minDistance("hoose", "ros")
// }
