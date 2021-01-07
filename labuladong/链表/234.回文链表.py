#
# @lc app=leetcode.cn id=234 lang=python3
#
# [234] 回文链表
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def reverseList(self, head):
        if not head or not head.next:
            return head
        last = self.reverseList(head.next)
        head.next.next = head
        head.next = None
        return last

    def isPalindrome(self, head) -> bool:
        """
        1.快慢指针找到链表的中点
	    2.翻转链表后半部分
	    3.回文校验
        """

        if not head or not head.next:
            return True

        slow = head
        fast = head
        while fast and fast.next:
            slow = slow.next
            fast = fast.next.next
        # 如果fast指针没有指向null，说明链表长度为奇数，slow还要再前进一步：
        if fast:
            slow = slow.next

        left = head
        right = self.reverseList(slow)

        while right:
            if left.val != right.val:
                return False
            left = left.next
            right = right.next
        return True

# @lc code=end

