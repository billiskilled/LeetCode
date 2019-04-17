package main

import "fmt"

/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] k个一组翻转链表
 *
 * https://leetcode-cn.com/problems/reverse-nodes-in-k-group/description/
 *
 * algorithms
 * Hard (50.49%)
 * Total Accepted:    9.3K
 * Total Submissions: 18.4K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给出一个链表，每 k 个节点一组进行翻转，并返回翻转后的链表。
 *
 * k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么将最后剩余节点保持原有顺序。
 *
 * 示例 :
 *
 * 给定这个链表：1->2->3->4->5
 *
 * 当 k = 2 时，应当返回: 2->1->4->3->5
 *
 * 当 k = 3 时，应当返回: 3->2->1->4->5
 *
 * 说明 :
 *
 *
 * 你的算法只能使用常数的额外空间。
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

// func reverseKGroup(head *ListNode, k int) *ListNode {
// 	var pre *ListNode
// 	var result *ListNode
// 	var len int

// 	var start *ListNode
// 	var end *ListNode
// 	start = head
// 	end = head

// 	for head != nil {
// 		len++
// 		if len%k == 0 {
// 			end = head

// 			var temp *ListNode

// 			temp = reverseKList(start, end)

// 			// if result == nil {
// 			// 	result = tempEnd
// 			// 	pre = tempStart
// 			// } else {
// 			// 	pre.Next = tempEnd
// 			// 	pre = tempStart
// 			// }

// 			// head = head.Next
// 			// start = head
// 		}
// 	}

// 	if len%k != 0 {
// 		pre.Next = start
// 	}

// 	return result
// }

func reverseKList(start *ListNode, end *ListNode) *ListNode {
	if start == end {
		return start
	}

	var cur *ListNode
	cur = start

	var pre *ListNode
	var next *ListNode
	next = start.Next

	for next != end {
		next = cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	cur.Next = pre
	end.Next = cur

	return end
}

func main() {
	var a *ListNode = new(ListNode)
	var b *ListNode = new(ListNode)
	var c *ListNode = new(ListNode)
	a.Val = 1
	b.Val = 2
	c.Val = 3
	a.Next = b
	b.Next = c

	var temp *ListNode = new(ListNode)
	temp = reverseKList(a, c)
	for temp != nil {
		fmt.Printf("%d, ", temp.Val)
		temp = temp.Next
	}
}
