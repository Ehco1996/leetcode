#
# @lc app=leetcode.cn id=6 lang=python3
#
# [6] Z 字形变换
#
# https://leetcode-cn.com/problems/zigzag-conversion/description/
#
# algorithms
# Medium (42.77%)
# Likes:    317
# Dislikes: 0
# Total Accepted:    37.5K
# Total Submissions: 87.2K
# Testcase Example:  '"PAYPALISHIRING"\n3'
#
# 将一个给定字符串根据给定的行数，以从上往下、从左到右进行 Z 字形排列。
#
# 比如输入字符串为 "LEETCODEISHIRING" 行数为 3 时，排列如下：
#
# L   C   I   R
# E T O E S I I G
# E   D   H   N
#
#
# 之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："LCIRETOESIIGEDHN"。
#
# 请你实现这个将字符串进行指定行数变换的函数：
#
# string convert(string s, int numRows);
#
# 示例 1:
#
# 输入: s = "LEETCODEISHIRING", numRows = 3
# 输出: "LCIRETOESIIGEDHN"
#
#
# 示例 2:
#
# 输入: s = "LEETCODEISHIRING", numRows = 4
# 输出: "LDREOEIIECIHNTSG"
# 解释:
#
# L     D     R
# E   O E   I I
# E C   I H   N
# T     S     G
#
#


class Solution:
    def convert1(self, s: str, numRows: int) -> str:
        from collections import defaultdict

        if numRows < 2:
            return s

        h = defaultdict(str)
        loop_num = 2 * (numRows - 1)
        while s:
            t = s[:loop_num]
            s = s[loop_num:]
            idx = 0
            for i in range(len(t)):
                idx = i if i < numRows else idx - 1
                h[idx] += t[i]

        res = ""
        for idx in range(numRows):
            res += h[idx]
        return res

    def convert(self, s, numRows):
        if numRows == 1 or numRows >= len(s):
            return s

        L = [""] * numRows
        index, step = 0, 1

        for x in s:
            L[index] += x
            if index == 0:
                step = 1
            elif index == numRows - 1:
                step = -1
            index += step

        return "".join(L)

    def convert_list(nums, n):
        """
        数组 Z 型输出；

        给一串数组，将其以“Z 形”放入 n 个桶中

        例：数组：[1, 2, 3, 4, 5, 6, 7]，将其放入 3 个桶中

        输出
        [
           [1, 6, 7],
           [2, 5],
           [3, 4],
        ]
        """
        import math

        L = [[] for i in range(n)]
        index_list = list(range(n))
        mirror = index_list[::-1]

        for i in range(math.ceil(len(nums) / n)):
            index_list.extend(mirror)
            mirror = mirror[::-1]
        for idx, num in zip(index_list, nums):
            L[idx].append(num)
        return L

