#
# @lc app=leetcode.cn id=797 lang=python3
#
# [797] 所有可能的路径
#

# @lc code=start
from typing import List
from copy import copy


class Solution:
    def allPathsSourceTarget(self, graph: List[List[int]]) -> List[List[int]]:
        """dfs"""
        res = []
        visited = set()

        def backtrack(path, can_move):
            if path[-1] == len(graph) - 1 and str(path) not in visited:
                visited.add(str(path))
                res.append(path)
            for p in can_move:
                path.append(p)
                backtrack(copy(path), graph[p])
                path.pop()

        for start in graph[0]:
            backtrack([0, start], graph[start])

        return res

    def test(self):
        return self.allPathsSourceTarget([[2], [], [1]])


# @lc code=end
