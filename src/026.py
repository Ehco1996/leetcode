'''
LeetCode: Remove Duplicates from Sorted Array

description:
Given a sorted array, remove the duplicates in place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this in place with constant memory.

For example,
Given input array nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively. It doesn't matter what you leave beyond the new length.


author: Ehco1996
time: 2017-11-3
'''

import time
import random


class MySolution:
    def removeDuplicates(self, nums):
        """
        :type nums: List[int]
        :rtype: int
        """
        if len(nums) == 0:
            return 0
        # seen = set()
        # for num in nums:
        #     if num not in seen:
        #         seen.add(num)
        return len(set(nums))


class Solution:
    def removeDuplicates(self, A):
        if not A:
            return 0

        last, i = 0, 1
        while i < len(A):
            if A[last] != A[i]:
                last += 1
                A[last] = A[i]
            i += 1

        return last + 1


def gen_lists(x):
    '''
    生成指定长度随机递增的列表
    '''
    ret = [1]
    for i in range(0, x - 1):
        min_num = ret[i]
        random_num = min_num + random.randint(0, 5)
        ret.append(random.randint(min_num, random_num))
    return ret


random_lists = []
# 随机生成100个随机长度的递增列表
for i in range(10000):
    random_lists.append(gen_lists(random.randint(5, 30)))


t1 = time.time()
for l in random_lists:
    MySolution().removeDuplicates(l)
t2 = time.time()

t3 = time.time()
for l in random_lists:
    Solution().removeDuplicates(l)
t4 = time.time()


print('我的程序运行的时间：{}'.format(t2 - t1))
print('他程序运行的时间：{}'.format(t4 - t3))
