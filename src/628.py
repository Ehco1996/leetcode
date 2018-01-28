'''
LeetCode: Maximum Product of Three Numbers

description: Given an integer array, find three numbers whose product is maximum and output the maximum product.



author: Ehco1996
time: 2017-08-12
'''


class Solution(object):
    def maximumProduct(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """

        '''
        res = 1
        for i in range(3):
            a = max(nums)
            nums.remove(a)
            res = res * a
        return res
        '''

        nums.sort(reverse=True)
        print(nums)
        lens = len(nums)-1
        return max(nums[0] * nums[1] * nums[2], nums[lens] * nums[0] * nums[lens-1])


l = [i for i in range(2**20)]

print(Solution().maximumProduct(l))
