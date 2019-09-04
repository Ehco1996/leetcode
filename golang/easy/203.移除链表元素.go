package main

/*
 * @lc app=leetcode.cn id=203 lang=golang
 *
 * [203] 移除链表元素
 *
 * https://leetcode-cn.com/problems/remove-linked-list-elements/description/
 *
 * algorithms
 * Easy (41.96%)
 * Likes:    278
 * Dislikes: 0
 * Total Accepted:    36.4K
 * Total Submissions: 86K
 * Testcase Example:  '[1,2,6,3,4,5,6]\n6'
 *
 * 删除链表中等于给定值 val 的所有节点。
 *
 * 示例:
 *
 * 输入: 1->2->6->3->4->5->6, val = 6
 * 输出: 1->2->3->4->5
 *
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeElements(head *ListNode, val int) *ListNode {
	ret := head
	last := head
	for head != nil {
		if head.Val == val {
			last.Next = head.Next
		} else {
			last = head
		}
		head = head.Next

		if ret.Val == val {
			ret = head
		}
	}
	return ret
}
