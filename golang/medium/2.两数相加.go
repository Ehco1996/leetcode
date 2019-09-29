package medium

/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 *
 * https://leetcode-cn.com/problems/add-two-numbers/description/
 *
 * algorithms
 * Medium (34.32%)
 * Likes:    2576
 * Dislikes: 0
 * Total Accepted:    146.6K
 * Total Submissions: 424.3K
 * Testcase Example:  '[2,4,3]\n[5,6,4]'
 *
 * 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
 *
 * 如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
 *
 * 您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
 *
 * 示例：
 *
 * 输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
 * 输出：7 -> 0 -> 8
 * 原因：342 + 465 = 807
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

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	// carry 表示十分有进一位
	ret := &ListNode{Val: 0}
	cur := ret
	carry := 0
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		// 看是否要进位
		if carry == 1 {
			sum += 1
			carry = 0
		}
		if sum > 9 {
			sum -= 10
			carry = 1
		}

		cur.Next = &ListNode{Val: sum}
		cur = cur.Next
		if l1 == nil && l2 == nil && carry == 1 {
			cur.Next = &ListNode{Val: 1}
		}
	}

	if ret.Next != nil {
		return ret.Next
	}
	return ret
}
