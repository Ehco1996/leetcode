/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 *
 * https://leetcode-cn.com/problems/palindrome-linked-list/description/
 *
 * algorithms
 * Easy (38.20%)
 * Likes:    262
 * Dislikes: 0
 * Total Accepted:    37K
 * Total Submissions: 95.7K
 * Testcase Example:  '[1,2]'
 *
 * 请判断一个链表是否为回文链表。
 *
 * 示例 1:
 *
 * 输入: 1->2
 * 输出: false
 *
 * 示例 2:
 *
 * 输入: 1->2->2->1
 * 输出: true
 *
 *
 * 进阶：
 * 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？
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

func isPalindrome(head *ListNode) bool {
	// 1.快慢指针找到链表的中点
	// 2.翻转链表前半部分
	// 3.回文校验
	if head == nil || head.Next == nil {
		return true
	}

	slow := head.Next
	fast := head.Next.Next
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}

	var prev *ListNode
	for head != slow {
		head.Next, prev, head = prev, head, head.Next
	}

	//如果是奇数个节点，去掉后半部分的第一个节点
	if fast != nil {
		slow = slow.Next
	}

	for prev != nil {
		if prev.Val != slow.Val {
			return false
		}
		prev = prev.Next
		slow = slow.Next
	}
	return true
}

// @lc code=end

