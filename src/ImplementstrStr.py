'''
LeetCode: Implement strStr()

description:
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Input: haystack = "hello", needle = "ll"
Output: 2
Input: haystack = "aaaaa", needle = "bba"
Output: -1

author: Ehco1996
time: 2017-11-04
'''


class Solution:
    def strStr(self, haystack, needle):
        """
        :type haystack: str
        :type needle: str
        :rtype: int
        """
        if needle not in haystack:
            return -1

        return haystack.index(needle)


print(Solution().strStr('aaaaa', 'bball'))
