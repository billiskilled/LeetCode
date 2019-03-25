# -*- coding: utf-8 -*-
#
# @lc app=leetcode.cn id=24 lang=python3
#
# [24] 两两交换链表中的节点
#
# https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
#
# algorithms
# Medium (56.70%)
# Total Accepted:    16.1K
# Total Submissions: 27.7K
# Testcase Example:  '[1,2,3,4]'
#
# 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
# 
# 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
# 
# 
# 
# 示例:
# 
# 给定 1->2->3->4, 你应该返回 2->1->4->3.
# 
# 
#
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None

class Solution:
    def swapPairs(self, head):
        if not head: return None

        first = None
        second = None
        newHead = None
        prev = None
        while head:
            if head.next != None:
                first = head
                second = head.next
                head = head.next.next
                second.next = first
                if not newHead:
                    newHead = second
                if prev:
                    prev.next = second
                prev = first
                prev.next = None
            else:
                if not newHead:
                    newHead = head
                if prev:
                    prev.next = head
                head = head.next
        return newHead

# head = ListNode(1)
# head.next = ListNode(2)
# head.next.next = ListNode(3)
# head.next.next.next = ListNode(4)

# while head:
#     print('%s;' % head.val)
#     head = head.next

# solu = Solution()

# result = solu.swapPairs(head)

# while result:
#     print('%s;' % result.val)
#     result = result.next
