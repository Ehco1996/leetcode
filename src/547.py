'''
LeetCode: Judge Route Circle

description:
There are N students in a class. Some of them are friends, while some are not. Their friendship is transitive in nature. For example, if A is a direct friend of B, and B is a direct friend of C, then A is an indirect friend of C. And we defined a friend circle is a group of students who are direct or indirect friends.

Given a N*N matrix M representing the friend relationship between students in the class. If M[i][j] = 1, then the ith and jth students are direct friends with each other, otherwise not. And you have to output the total number of friend circles among all the students.

author: Ehco1996
time: 2017-08-14
'''


class Solution(object):
    def findCircleNum(self, M):
        """
        :type M: List[List[int]]
        :rtype: int
        """

        N = len(M)
        seen = set()

        def dfs(node):
            seen.add(node)
            for friend in range(N):
                if M[node][friend] and friend not in seen:
                    dfs(friend)

        circle = 0
        for node in range(N):
            if node not in seen:
                dfs(node)
                circle += 1

        return circle


l = [[1, 1, 0,1],
     [1, 1, 0,1],
     [0, 0, 1,1],
     [0, 0, 1,1]]

print(Solution().findCircleNum(l))
