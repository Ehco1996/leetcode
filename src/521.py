'''
LeetCode: Longest Uncommon Subsequence I


description:
Given a group of two strings, you need to find the longest uncommon subsequence of this group of two strings. The longest uncommon subsequence is defined as the longest subsequence of one of these strings and this subsequence should not be any subsequence of the other strings.
A subsequence is a sequence that can be derived from one sequence by deleting some characters without changing the order of the remaining elements. Trivially, any string is a subsequence of itself and an empty string is a subsequence of any string.
The input will be two strings, and the output needs to be the length of the longest uncommon subsequence. If the longest uncommon subsequence doesn't exist, return -1.

Input: "aba", "cdc"
Output: 3
Explanation: The longest uncommon subsequence is "aba" (or "cdc"), 
because "aba" is a subsequence of "aba", 
but not a subsequence of any other strings in the group of two strings. 

author: Ehco1996
time: 2017-12-6
'''


class Solution:
    def genSubsequence(self, s):
        '''
        genrate the subsequence list of the string
        :type s: str
        :rtype sorted(list) 
        '''
        res = ['', ]
        for i in range(len(s)):
            tmp_s = s[i:]
            for j in range(len(tmp_s)):
                res.append(tmp_s[j:])
        return sorted(list(set(res)), key=len, reverse=True)

    def findLUSlength(self, a, b):
        """
        :type a: str
        :type b: str]
        :rtype: int
        """
        a_list = self.genSubsequence(a)
        b_list = self.genSubsequence(b)
        a_res = -1
        b_res = -1
        for w in a_list:
            if w not in b_list:
                a_res = len(w)
                break
        for w in b_list:
            if w not in a_list:
                b_res = len(w)
                break
        return a_res if a_res > b_res else b_res


s = Solution()

s1 = 'aefawfawfawfaw'
s2 = 'aefawfeawfwafwaef'

s1=s2 = 'aaa'
a = s.genSubsequence(s1)
b = s.genSubsequence(s2)
print(a)
print(b)

c = s.findLUSlength(s1, s2)
print(c)
