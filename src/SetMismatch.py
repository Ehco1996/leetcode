'''
LeetCode: Set Mismatch

description: The set S originally contains numbers from 1 to n. But unfortunately, due to the data error, one of the numbers in the set got duplicated to another number in the set, which results in repetition of one number and loss of another number.

Given an array nums representing the data status of this set after the error. Your task is to firstly find the number occurs twice and then find the number that is missing. Return them in the form of an array.


author: Ehco1996
time: 2017-08-23
'''

from collections import Counter
class Solution(object):
    def findErrorNums(self, nums):
        """
        :type nums: List[int]
        :rtype: List[int]
        """
        res = []

        nums = sorted(nums)
        l = set(nums)
        s = set(i for i in range(1, len(nums) + 1))
       
        # 找到少的那个数字
        x = list(s - l)[0]
        # 找到多的那个数字
        r = Counter(nums).most_common()[0][0]
        
        res.append(r)
        res.append(x)

        return res


# best solution
def findErrorNums(A):
    N = len(A)
    count = [0] * (N+1)
    for x in A:
      count[x] += 1
    for x in range(1, len(A)+1):
        if count[x] == 2:
            twice = x
        if count[x] == 0:
            never = x
    return twice, never

# TEST:
nums=[1,2,2,4]

#print(Solution().findErrorNums(nums))

print(findErrorNums(nums))
