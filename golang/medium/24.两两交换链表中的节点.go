/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 *
 * https://leetcode-cn.com/problems/swap-nodes-in-pairs/description/
 *
 * algorithms
 * Medium (61.29%)
 * Likes:    285
 * Dislikes: 0
 * Total Accepted:    40.9K
 * Total Submissions: 65.6K
 * Testcase Example:  '[1,2,3,4]'
 *
 * 给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。
 *
 * 你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。
 *
 *
 *
 * 示例:
 *
 * 给定 1->2->3->4, 你应该返回 2->1->4->3.
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

// 1-2-3-4 -> 2-1-3-4
func swapPairs(head *ListNode) *ListNode {
	if head != nil && head.Next != nil {
		walker := head
		runner := walker.Next
		head = runner
		tail := &ListNode{Val: -1}
		for walker != nil && runner != nil {
			walker.Next, runner.Next = runner.Next, walker
			tail.Next = runner
			tail = walker
			walker = walker.Next
			if walker != nil {
				runner = walker.Next
			}
		}
	}
	return head

}

