/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

/**
 * 需要多个记录位置的指针
 * 记录当前处理的k的前后指针curStart、curEnd
 * 记录处理过的最后一个节点的指针dealedTail
 */
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil {
		return nil
	}

	var result *ListNode
	result = new(ListNode)
	var dealedTail = result

	var n = 0
	var curNode = head

	var partStart = head
	for curNode != nil {
		tempNode := curNode

		curNode = curNode.Next
		n++
		if n == k {
			reversePart(partStart, tempNode)
			dealedTail.Next = tempNode
			dealedTail = partStart
			partStart = curNode
			n = 0
		}
	}
	dealedTail.Next = partStart
	return result.Next
}

func reversePart(start *ListNode, end *ListNode) {
	var pre *ListNode
	pre = nil
	cur := start
	for cur != end {
		temp := cur
		cur = cur.Next
		temp.Next = pre
		pre = temp
	}
	end.Next = pre
}
