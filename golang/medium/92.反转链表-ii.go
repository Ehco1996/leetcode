package medium

/*
 * @lc app=leetcode.cn id=92 lang=golang
 *
 * [92] 反转链表 II
 *
 * https://leetcode-cn.com/problems/reverse-linked-list-ii/description/
 *
 * algorithms
 * Medium (45.95%)
 * Likes:    217
 * Dislikes: 0
 * Total Accepted:    20.5K
 * Total Submissions: 43.7K
 * Testcase Example:  '[1,2,3,4,5]\n2\n4'
 *
 * 反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。
 *
 * 说明:
 * 1 ≤ m ≤ n ≤ 链表长度。
 *
 * 示例:
 *
 * 输入: 1->2->3->4->5->NULL, m = 2, n = 4
 * 输出: 1->4->3->2->5->NULL
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummy := new(ListNode)
	head, dummy.Next = dummy, head
	for i := 0; i < m-1; i++ {
		head = head.Next
	}
	newHead, _ := reverseList(head.Next, n-m+1)
	head.Next = newHead
	return dummy.Next
}

// return new head and the head of the rest
func reverseList(head *ListNode, cnt int) (*ListNode, *ListNode) {
	if cnt == 1 {
		return head, head.Next
	}
	newHead, restHead := reverseList(head.Next, cnt-1)
	head.Next.Next = head
	head.Next = restHead
	return newHead, restHead
}
