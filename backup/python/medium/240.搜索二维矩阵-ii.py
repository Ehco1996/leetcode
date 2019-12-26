#
# @lc app=leetcode.cn id=240 lang=python3
#
# [240] 搜索二维矩阵 II
#
# https://leetcode-cn.com/problems/search-a-2d-matrix-ii/description/
#
# algorithms
# Medium (36.79%)
# Likes:    126
# Dislikes: 0
# Total Accepted:    21K
# Total Submissions: 56.2K
#
# 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target。该矩阵具有以下特性：
#
#
# 每行的元素从左到右升序排列。
# 每列的元素从上到下升序排列。
#
#
# 示例:
#
# 现有矩阵 matrix 如下：
#
# [
# ⁠ [1,   4,  7, 11, 15],
# ⁠ [2,   5,  8, 12, 19],
# ⁠ [3,   6,  9, 16, 22],
# ⁠ [10, 13, 14, 17, 24],
# ⁠ [18, 21, 23, 26, 30]
# ]
#
#
# 给定 target = 5，返回 true。
#
# 给定 target = 20，返回 false。
#
#
class Solution:
    def searchMatrix(self, matrix, target):
        """
        :type matrix: List[List[int]]
        :type target: int
        :rtype: bool
        """
        m = len(matrix)
        for i in range(m - 1, -1, -1):
            l = matrix[i]
            if not l:
                return False
            if l[0] == target:
                return True
            if l[0] < target and self.bs_find(l, target):
                return True
        return False

    def bs_find(self, l, target):
        if not l:
            return False

        lens = len(l)
        idx = lens // 2
        val = l[idx]
        if val == target:
            return True

        if val < target:
            return self.bs_find(l[idx + 1 :], target)
        if val > target:
            return self.bs_find(l[:idx], target)

    # def bs_find(self, l, target):
    #     start_idx = 0
    #     end_idx = len(l) - 1

    #     while start_idx <= end_idx:
    #         mid = (start_idx + end_idx) // 2
    #         val = l[mid]
    #         if val == target:
    #             return True
    #         elif val < target:
    #             start_idx = mid + 1
    #         else:
    #             end_idx = mid - 1
    #     return False

