package main

/*
 * @lc app=leetcode.cn id=70 lang=golang
 *
 * [70] 爬楼梯
 */
func climbStairs(n int) int {
	// 递归 + 记忆化
	if n < 0 {
		return 1
	}
	if n <= 2 {
		return n
	}
	// var dp []int
	// dp = append(dp, 1)
	// dp = append(dp, 2)

	// var result = 0
	// temp := make([]int, 10000)
	// result += core(n, temp)
	// climbStairsCore(n, &result, temp)
	// return result

	// 动态规划
	var a int
	var b int
	var temp int
	a = 1
	b = 2

	for i := 2; i < n; i++ {
		temp = b
		b += a
		a = temp
	}
	return b
}

/**
// 使用一种指针存储result，很奇怪的写法
func climbStairsCore(n int, result *int, temp []int) {
	if n <= 1 {
		*result += 1
		return
	}
	if n == 2 {
		*result += 2
		return
	}

	if temp[n-1] != 0 {
		*result += temp[n-1]
	} else {
		climbStairsCore(n-1, result, temp)
	}
	if temp[n-2] != 0 {
		*result += temp[n-2]
	} else {
		climbStairsCore(n-2, result, temp)
	}
	temp[n] = *result
}
*/

/**
// 看起来正常一点的递归的写法
func core(n int, temp []int) int {
	if n <= 1 {
		return 1
	}
	if n == 2 {
		return 2
	}

	var cur int
	if temp[n-1] != 0 {
		cur = temp[n-1]
	} else {
		cur = core(n-1, temp)
	}
	if temp[n-2] != 0 {
		cur += temp[n-2]
	} else {
		cur += core(n-2, temp)
	}
	temp[n] = cur
	return cur
}
*/

// func main() {
// 	fmt.Println(climbStairs(200))
// }
