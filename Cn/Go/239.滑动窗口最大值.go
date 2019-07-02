/*
 * @lc app=leetcode.cn id=239 lang=golang
 *
 * [239] 滑动窗口最大值
 */
func maxSlidingWindow(nums []int, k int) []int {
	len := len(nums)

	if len == 0 {
		return make([]int, 0)
	}

	var result = make([]int, len-k+1)

	var tempK = make([]int, len)

	i := 0
	var max int
	max = ^int(^uint(0) >> 1)

	for i < k {
		if nums[i] > max {
			max = nums[i]
			tempK = tempK[0:0]
			tempK = append(tempK, nums[i])
		} else {
			for ind, v := range tempK {
				if nums[i] > v {
					tempK = tempK[0:ind]
					break
				}
			}
			tempK = append(tempK, nums[i])
		}
		i++
	}

	resultI := 0
	result[resultI] = max
	resultI++

	for i < len {
		if nums[i] > max {
			max = nums[i]
			tempK = tempK[0:0]
			tempK = append(tempK, nums[i])
		} else {
			for ind, v := range tempK {
				if nums[i] > v {
					tempK = tempK[0:ind]
					break
				}
			}
			tempK = append(tempK, nums[i])
		}

		if nums[i-k] == max {
			max = tempK[1]
			tempK = tempK[1:]
		}

		result[resultI] = max
		resultI++
		i++
	}

	return result
}
