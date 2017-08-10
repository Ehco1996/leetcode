'''
LeetCode: Reverse Integer

description: Given an integer, write a function to determine if it is a power of three..



author: Ehco1996
time: 2017-08-10
'''


'''
我们首先分析3的幂的特点，假设一个数Num是3的幂，那么所有Num的约数都是3的幂，
如果一个数n小于Num且是3的幂，
那么这个数n一定是Num的约数。
了解上述性质，我们只需要找到一个最大的3的幂，看看参数n是不是此最大的幂的约数就行了
'''


class Solution(object):
    def isPowerOfThree(self, n):
        """
        :type n: int
        :rtype: bool
        """
        if n <= 0:
            return False

        bigPower = 3**19

        if bigPower % n == 0:
            return True
        else:
            return False


print(Solution().isPowerOfThree(14348907))
