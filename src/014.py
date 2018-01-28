'''
LeetCode: Longest Common Prefix


description:
Write a function to find the longest common prefix string amongst an array of strings.

author: Ehco1996
time: 2017-10-30
'''


class Solution:
    def longestCommonPrefix(self, strs):
        """
        :type strs: List[str]
        :rtype: str
        """
        if len(strs)==0:
            return ""

        for i in range(len(strs[0])):
            for string in strs[1:]:
                if i >= len(string) or string[i] != strs[0][i]:
                    return strs[0][:i]
        return strs[0]


test = ['abc', 'abccc', 'ab']

print(Solution().longestCommonPrefix(test))
