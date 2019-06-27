#
# @lc app=leetcode.cn id=83 lang=python3
#
# [83] 删除排序链表中的重复元素
#
# https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/description/
#
# algorithms
# Easy (46.07%)
# Likes:    158
# Dislikes: 0
# Total Accepted:    29.4K
# Total Submissions: 63.8K
# Testcase Example:  '[1,1,2]'
#
# 给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。
#
# 示例 1:
#
# 输入: 1->1->2
# 输出: 1->2
#
#
# 示例 2:
#
# 输入: 1->1->2->3->3
# 输出: 1->2->3
#
#
# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None

    def __repr__(self):
        res = ""
        node = self
        while node.next:
            res += f"{node.val}->"
            node = node.next
        res += f"{node.val}"
        return res


def build_single_list_from_list(l):
    head = ListNode(l[0])
    node = head
    for i in l[1:]:
        n = ListNode(i)
        node.next = n
        node = n
    return head


class Solution:
    def deleteDuplicates(self, head: ListNode) -> ListNode:
        if not head:
            return []

        prev = head
        node = head
        seen = set()

        while node:
            if node.val not in seen:
                seen.add(node.val)
                prev = node
                node = node.next
            else:
                next_node = node.next
                prev.next = next_node
                node.next = None
                node = next_node
        return head

