#
# @lc app=leetcode.cn id=92 lang=python3
#
# [92] 反转链表 II
#

# @lc code=start
# Definition for singly-linked list.
# class ListNode:
#     def __init__(self, x):
#         self.val = x
#         self.next = None


class Solution:
    def reverseBetween(self, head, m: int, n: int):
        # base case
        if m == 1:
            res, _ = self.reverseN(head, n)
            return res
        # 前进到反转的起点触发 base case
        head.next = self.reverseBetween(head.next, m - 1, n - 1)
        return head

    def reverseN(self, head, n):
        """反转前n个节点,返回新的节点和剩下的节点"""
        if n == 1:
            return head, head.next
        # 以 head.next 为起点，需要反转前 n - 1 个节点
        new_head, rest_head = self.reverseN(head.next, n - 1)

        # 将后驱节点连上自己
        head.next.next = head

        # 让反转之后的 head 节点和后面的节点连起来
        head.next = rest_head

        return new_head, rest_head


# @lc code=end

