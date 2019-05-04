/*
 * @lc app=leetcode.cn id=20 lang=golang
 *
 * [20] 有效的括号
 *
 * https://leetcode-cn.com/problems/valid-parentheses/description/
 *
 * algorithms
 * Easy (37.32%)
 * Total Accepted:    65.2K
 * Total Submissions: 174.3K
 * Testcase Example:  '"()"'
 *
 * 给定一个只包括 '('，')'，'{'，'}'，'['，']' 的字符串，判断字符串是否有效。
 *
 * 有效字符串需满足：
 *
 *
 * 左括号必须用相同类型的右括号闭合。
 * 左括号必须以正确的顺序闭合。
 *
 *
 * 注意空字符串可被认为是有效字符串。
 *
 * 示例 1:
 *
 * 输入: "()"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: "()[]{}"
 * 输出: true
 *
 *
 * 示例 3:
 *
 * 输入: "(]"
 * 输出: false
 *
 *
 * 示例 4:
 *
 * 输入: "([)]"
 * 输出: false
 *
 *
 * 示例 5:
 *
 * 输入: "{[]}"
 * 输出: true
 *
 */
// package main

// import "fmt"

// func main() {
// 	s := "()"
// 	result := isValid(s)
// 	fmt.Print("Result: ", result)
// }

func isValid(s string) bool {
	var byteArr []byte
	byteArr = []byte(s)

	index := 0
	var stack []byte
	stack = make([]byte, len(s))

	for _, e := range byteArr {
		if e == byte('{') || e == byte('(') || e == byte('[') {
			stack[index] = e
			index++
		} else {
			if e == byte('}') && index > 0 && stack[index-1] == byte('{') {
				index--
				continue
			}
			if e == byte(']') && index > 0 && stack[index-1] == byte('[') {
				index--
				continue
			}
			if e == byte(')') && index > 0 && stack[index-1] == byte('(') {
				index--
				continue
			}
			return false
		}
	}
	if index == 0 {
		return true
	}
	return false
}
