'''
LeetCode: Reverse Integer

description: Reverse digits of an integer.



author: Ehco1996
time: 2017-08-10
'''


class Solution(object):
    def reverse(self, x):
        """
        :type x: int
        :rtype: int
        """

        if x >= 0:
            y = int(str(x)[::-1])
            if y > 2**31:
                return 0
            else:
                return y
        else:
            y = int(str(x)[1:][::-1])
            if y > 2 ** 31:
                return 0
            else:
                return -y


print(Solution().reverse(112))
