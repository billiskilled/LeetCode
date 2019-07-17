package main

/*
 * @lc app=leetcode.cn id=120 lang=golang
 *
 * [120] 三角形最小路径和
 */
func minimumTotal(triangle [][]int) int {
	/**
	// 自己写的解法，逻辑过于复杂；在于从下到上的dp记录没有处理好
	var dp []int
	level := len(triangle)

	var start int
	var end int
	var preStart int
	for i := 0; i < level; i++ {
		preStart = start
		start = end
		col := len(triangle[i])
		for j := 0; j < col; j++ {
			end++
			if i == 0 {
				dp = append(dp, triangle[i][j])
			} else {
				if j == 0 {
					dp = append(dp, triangle[i][j]+dp[preStart])
				} else if j == col-1 {
					dp = append(dp, triangle[i][j]+dp[start-1])
				} else {
					min := getMinTwo(dp[preStart+j-1], dp[preStart+j])
					dp = append(dp, triangle[i][j]+min)
				}
			}
		}
		if i == level-1 {
			return getMin(dp, start, end)
		}
	}
	return 0
	*/

	/**
	// 使用两维数组的解法
	var dp [1000][1000]int

	level := len(triangle)

	levelLen := len(triangle[level-1])

	for i := 0; i < levelLen; i++ {
		dp[level-1][i] = triangle[level-1][i]
	}

	for i := level - 2; i >= 0; i-- {
		curLen := len(triangle[i])
		for j := 0; j < curLen; j++ {
			dp[i][j] = getMinTwo(dp[i+1][j], dp[i+1][j+1]) + triangle[i][j]
		}
	}

	return dp[0][0]
	*/

	/**
	var dp []int

	level := len(triangle)

	levelLen := len(triangle[level-1])

	for i := 0; i < levelLen; i++ {
		dp = append(dp, triangle[level-1][i])
	}

	for i := level - 2; i >= 0; i-- {
		curLen := len(triangle[i])
		for j := 0; j < curLen; j++ {
			dp[j] = getMinTwo(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
	*/

	// 回溯法，不能使用记忆化的方法，所以数量多了的时候会跑死
	level := len(triangle) - 1
	result := int(^uint(0) >> 1)

	core(triangle, level, 0, 0, 0, &result)
	return result
}

func core(a [][]int, level int, i int, j int, sum int, result *int) {
	if i == level {
		if a[i][j]+sum < *result {
			*result = a[i][j] + sum
		}
		return
	}

	sum += a[i][j]
	core(a, level, i+1, j, sum, result)
	core(a, level, i+1, j+1, sum, result)
}

func getMin(a []int, start int, end int) int {
	result := int(^uint(0) >> 1)
	for i := start; i < end; i++ {
		if a[i] < result {
			result = a[i]
		}
	}
	return result
}

func getMinTwo(x int, y int) int {
	if x < y {
		return x
	}
	return y
}

// func main() {
// 	temp := [][]int{{2}, {3, 4}, {6, 5, 7}, {4, 1, 8, 3}}
// 	fmt.Println(minimumTotal(temp))
// }
