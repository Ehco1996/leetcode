'''
LeetCode: Two Sum

description: Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

author: Ehco1996
time: 2017-08-09
'''


class Solution:
    def twoSum(self, nums, target):
        """
        :type nums: List[int]
        :type target: int
        :rtype: List[int]
        """

        for i, num in enumerate(nums):
            x = target - num
            if x in nums and i != nums.index(x):
                return sorted([i, nums.index(x)])


# test
print(Solution().twoSum([1, 9, 7, 100], 101))
