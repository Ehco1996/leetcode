'''
LeetCode: Palindrome Number

description: Determine whether an integer is a palindrome. Do this without extra space.


author: Ehco1996
time: 2017-10-27
'''


class Solution:
    def isPalindrome(self, x):
        """
        :type x: int
        :rtype: bool
        """
        if x < 0:
            return False
        elif x == 0:
            return True
        else:
            # 反转这个数并比较是否相等
            bakx = x
            num = 0
            while x != 0:
                num = num * 10 + x % 10
                x = x // 10
            return bakx == num




print(Solution().isPalindrome(0))