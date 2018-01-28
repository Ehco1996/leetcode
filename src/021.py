'''
LeetCode: Merge Two Sorted Lists

description:
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.


author: Ehco1996
time: 2017-11-2
'''


# Definition for singly-linked list.
class ListNode:
    def __init__(self, x):
        self.val = x
        self.next = None


class Solution:
    def mergeTwoLists(self, l1, l2):
        """
        :type l1: ListNode
        :type l2: ListNode
        :rtype: ListNode
        """
        if l1 == None:
            return l2
        if l2 == None:
            return l1
        ret = ListNode(0)
        head = ret
        while l1 != None and l2 != None:
            if l1.val > l2.val:
                ret.next = l2
                l2 = l2.next
            else:
                ret.next = l1
                l1 = l1.next
            ret = ret.next

        if l1 == None:
            ret.next = l2

        if l2 == None:
            ret.next = l1

        return head.next


def pprint(node):
    '''从当前节点开始显示所有连接上的链表'''

    while node:
        print('{} -> '.format(node.val), end='')
        node = node.next
        if node is None:
            break


nodea = ListNode('1')
nodeb = ListNode('2')
nodec = ListNode('3')

noded = ListNode('1')
nodee = ListNode('4')

nodea.next = nodeb
nodeb.next = nodec

noded.next = nodee
pprint(Solution().mergeTwoLists(nodea, noded))
