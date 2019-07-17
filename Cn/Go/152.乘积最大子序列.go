package main

/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子序列
 */
func maxProduct(nums []int) int {
	var max []int
	var min []int

	max = append(max, nums[0])
	min = append(min, nums[0])

	numsLen := len(nums)

	for i := 1; i < numsLen; i++ {
		if nums[i] >= 0 {
			if max[i-1]*nums[i] > nums[i] {
				max = append(max, max[i-1]*nums[i])
			} else {
				max = append(max, nums[i])
			}

			if min[i-1]*nums[i] < nums[i] {
				min = append(min, min[i-1]*nums[i])
			} else {
				min = append(min, nums[i])
			}
		} else {
			if min[i-1]*nums[i] > nums[i] {
				max = append(max, min[i-1]*nums[i])
			} else {
				max = append(max, nums[i])
			}

			if max[i-1]*nums[i] < nums[i] {
				min = append(min, max[i-1]*nums[i])
			} else {
				min = append(min, nums[i])
			}
		}
	}
	var result int
	const INT_MAX = int(^uint(0) >> 1)
	result = ^INT_MAX
	maxLen := len(max)
	for i := 0; i < maxLen; i++ {
		if max[i] > result {
			result = max[i]
		}
	}
	return result
}
