#
# @lc app=leetcode.cn id=565 lang=python3
#
# [565] 数组嵌套
#

# @lc code=start
class Solution:
    def arrayNesting(self, nums) -> int:
        """
        关键在与避免重复计算
        一旦一个元素被访问过，那么他的路径的长度就已经是决定的了
        比如 path = 0-5-6-2 ，那么无论是从0开始访问 还是从6开始访问，访问的路径长度都是6 并且元素都是 0562这四个
        """
        res = 0
        hited = set()
        for num in nums:
            cur = num
            visted = set()
            # path = ""
            while cur not in visted and cur not in hited:
                visted.add(cur)
                hited.add(cur)
                # path += str(cur)
                cur = nums[cur]
            # print(path)
            res = max(res, len(visted))
        return res

    def test(self):
        print(self.arrayNesting([5, 4, 0, 3, 1, 6, 2]))


# @lc code=end

