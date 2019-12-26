/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 *
 * https://leetcode-cn.com/problems/reverse-linked-list/description/
 *
 * algorithms
 * Easy (65.64%)
 * Likes:    673
 * Dislikes: 0
 * Total Accepted:    128.5K
 * Total Submissions: 195.2K
 * Testcase Example:  '[1,2,3,4,5]'
 *
 * 反转一个单链表。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL
 * 输出: 5->4->3->2->1->NULL
 *
 * 进阶:
 * 你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
 *
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	// 不断将节点移动到链表头部

	if head == nil {
		return head
	}

	dummy := head
	for head.Next != nil {
		node := head.Next
		head.Next = node.Next

		node.Next = dummy
		dummy = node

	}
	return dummy
}

// @lc code=end

