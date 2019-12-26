package easy

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
	if head == nil {
		return head
	}
	for head != nil && head.Val == val {
		head = head.Next
	}
	pre := head
	now := head
	for now != nil {
		if now.Val == val {
			pre.Next = now.Next
		} else {
			pre = now
		}
		now = now.Next
	}
	return head
}
