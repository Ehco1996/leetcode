'''
LeetCode: Reverse Integer

description: Given an integer, write a function to determine if it is a power of three..



author: Ehco1996
time: 2017-08-11
'''


import math


class Solution(object):
    def judgeSquareSum(self, c):
        """
        :type c: int
        :rtype: bool
        """
        num = int(math.sqrt(c))
        for a in range(0, num+1):
            b = c - a**2
            if math.sqrt(b) == int(math.sqrt(b)):
                return True

        return False


print(Solution().judgeSquareSum(2**32))