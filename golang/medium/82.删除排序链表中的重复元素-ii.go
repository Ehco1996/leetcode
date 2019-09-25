/*
 * @lc app=leetcode.cn id=82 lang=golang
 *
 * [82] 删除排序链表中的重复元素 II
 *
 * https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii/description/
 *
 * algorithms
 * Medium (42.71%)
 * Likes:    143
 * Dislikes: 0
 * Total Accepted:    19K
 * Total Submissions: 43.2K
 * Testcase Example:  '[1,2,3,3,4,4,5]'
 *
 * 给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。
 *
 * 示例 1:
 *
 * 输入: 1->2->3->3->4->4->5
 * 输出: 1->2->5
 *
 *
 * 示例 2:
 *
 * 输入: 1->1->1->2->3
 * 输出: 2->3
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	dum := &ListNode{
		Next: head,
		Val:  -1,
	}

	slow := dum
	fast := dum.Next

	for fast != nil {
		if fast.Next != nil && fast.Val == fast.Next.Val {
			temp := fast.Val
			for fast != nil && fast.Val == temp {
				fast = fast.Next
			}
		} else {
			slow.Next = fast
			slow = fast
			fast = fast.Next
		}
	}
	slow.Next = fast
	return dum.Next
}

