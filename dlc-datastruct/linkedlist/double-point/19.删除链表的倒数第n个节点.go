// package medium

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第N个节点
 *
 * https://leetcode-cn.com/problems/remove-nth-node-from-end-of-list/description/
 *
 * algorithms
 * Medium (34.89%)
 * Likes:    525
 * Dislikes: 0
 * Total Accepted:    73.8K
 * Total Submissions: 207.8K
 * Testcase Example:  '[1,2,3,4,5]\n2'
 *
 * 给定一个链表，删除链表的倒数第 n 个节点，并且返回链表的头结点。
 *
 * 示例：
 *
 * 给定一个链表: 1->2->3->4->5, 和 n = 2.
 *
 * 当删除了倒数第二个节点后，链表变为 1->2->3->5.
 *
 *
 * 说明：
 *
 * 给定的 n 保证是有效的。
 *
 * 进阶：
 *
 * 你能尝试使用一趟扫描实现吗？
 *
 */
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 双指针  p q ,当q到达终点，q和p相距n时 满足要求

	dummy := &ListNode{Val: -1, Next: head}

	p := dummy
	q := dummy
	// q 先走n步 这时两者相差n
	for i := 0; i <= n; i++ {
		q = q.Next
	}

	for q != nil {
		q = q.Next
		p = p.Next
	}

	node := p.Next
	p.Next = node.Next
	return dummy.Next
}
