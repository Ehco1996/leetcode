package medium

/*
 * @lc app=leetcode.cn id=86 lang=golang
 *
 * [86] 分隔链表
 *
 * https://leetcode-cn.com/problems/partition-list/description/
 *
 * algorithms
 * Medium (51.51%)
 * Likes:    120
 * Dislikes: 0
 * Total Accepted:    15.2K
 * Total Submissions: 28.9K
 * Testcase Example:  '[1,4,3,2,5,2]\n3'
 *
 * 给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。
 *
 * 你应当保留两个分区中每个节点的初始相对位置。
 *
 * 示例:
 *
 * 输入: head = 1->4->3->2->5->2, x = 3
 * 输出: 1->2->2->4->3->5
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

func partition(head *ListNode, x int) *ListNode {
	// 将链表拆成两个部分，最后拼接
	if head == nil {
		return nil
	}

	before_head := &ListNode{
		Val:  -1,
		Next: nil,
	}
	after_head := &ListNode{
		Val:  -1,
		Next: nil,
	}

	before := before_head
	after := after_head

	for head != nil {
		if head.Val < x {
			before.Next = head
			before = before.Next
		} else {
			after.Next = head
			after = after.Next
		}
		head = head.Next
	}
	after.Next = nil
	before.Next = after_head.Next
	return before_head.Next
}
