'''
LeetCode: Maximum Product Subarray

description: Find the contiguous subarray within an array (containing at least one number) which has the largest product.



author: Ehco1996
time: 2017-08-12
'''


class Solution(object):

    def is_same_sign(self, a, b):
        return True if a * b > 0 else False

    def maxProduct(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """

        res = nums[0]
        imax = res
        imin = res
        
        for i in nums[1:]:
            #当这个数为负时，数字越小成绩越大，数子越大乘积越小，所以交换imax，imin
            if i < 0:
                imax, imin = imin, imax
            res = res * i
            imax = max(i, i * imax)
            imin = min(i, i * imin)
            # 判断谁最大
            res = max(res, imax)

        return res


l = [-2, 3, -4]
print(Solution().maxProduct(l))
