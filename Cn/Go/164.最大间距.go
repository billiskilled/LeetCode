package main

/*
 * @lc app=leetcode.cn id=164 lang=golang
 *
 * [164] 最大间距
 */
func maximumGap(nums []int) int {
	numsLen := len(nums)
	if numsLen < 2 {
		return 0
	}

	const maxInt = int(^uint(0) >> 1)
	const minInt = ^maxInt

	var min = maxInt
	var max = minInt

	for _, v := range nums {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	buckSize := getMax(1, (max-min)/(numsLen-1))

	type buc struct {
		min int
		max int
	}

	bucksNum := (max-min)/buckSize + 1
	bucks := make([]buc, bucksNum)
	bucksFlag := make([]int, bucksNum)
	for i := 0; i < bucksNum; i++ {
		bucks[i].min = maxInt
		bucks[i].max = minInt
	}
	// for i, v := range bucks {
	// 	fmt.Println(i)
	// 	v.min = maxInt
	// 	v.max = minInt
	// }
	for _, v := range nums {
		loc := (v - min) / buckSize

		bucksFlag[loc] = 1
		if v < bucks[loc].min {
			bucks[loc].min = v
		}
		if v > bucks[loc].max {
			bucks[loc].max = v
		}
	}

	maxGap := minInt

	var prev buc
	var prevFlag int
	for i, v := range bucks {
		if bucksFlag[i] == 0 {
			continue
		}
		maxGap = getMax(v.max-v.min, maxGap)

		if prevFlag != 0 {
			maxGap = getMax(v.min-prev.max, maxGap)
		}
		prev = v
		prevFlag = 1
	}

	return maxGap
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

// func main() {
// 	a := []int{3, 6, 9, 1}
// 	fmt.Println(maximumGap(a))
// }
