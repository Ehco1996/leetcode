#
# @lc app=leetcode.cn id=25 lang=python3
#
# [25] K 个一组翻转链表
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, val=0, next=None):
#         self.val = val
#         self.next = next
class Solution:
    def reverseList(self, a, b):
        """反转a,b之前的链表"""

        pre = None
        cur = a
        nxt = a
        while cur != b:
            nxt = cur.next
            cur.next = pre
            pre = cur
            cur = nxt
        return pre

    def reverseKGroup(self, head, k):
        if not head:
            return head
        a = b = head
        # 找到k个
        for i in range(k):
            if not b:  # 不足k个 不用反转
                return head
            b = b.next
        # 反转前k个元素
        new_head = self.reverseList(a, b)
        a.next = self.reverseKGroup(b, k)
        return new_head

# @lc code=end

