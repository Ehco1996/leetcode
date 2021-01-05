#
# @lc app=leetcode.cn id=341 lang=python3
#
# [341] 扁平化嵌套列表迭代器
#

# @lc code=start
# """
# This is the interface that allows for creating nested lists.
# You should not implement it, or speculate about its implementation
# """
# class NestedInteger:
#    def isInteger(self) -> bool:
#        """
#        @return True if this NestedInteger holds a single integer, rather than a nested list.
#        """
#
#    def getInteger(self) -> int:
#        """
#        @return the single integer that this NestedInteger holds, if it holds a single integer
#        Return None if this NestedInteger holds a nested list
#        """
#
#    def getList(self) -> [NestedInteger]:
#        """
#        @return the nested list that this NestedInteger holds, if it holds a nested list
#        Return None if this NestedInteger holds a single integer
#        """


class NestedIterator:
    def __init__(self, nestedList):
        self.res = nestedList

    def next(self) -> int:
        # hasNext 方法保证了第一个元素一定是整数类型
        return self.res.pop(0).getInteger()

    def hasNext(self) -> bool:
        while len(self.res) > 0 and not self.res[0].isInteger():
            first = self.res.pop(0).getList()
            print(first)
            for i in first[::-1]:
                self.res.insert(0, i)
        return len(self.res) > 0

    @classmethod
    def test(cls, l=[[1, 1], 2, [1, 1]]):
        i, v = NestedIterator(l), []
        while i.hasNext():
            v.append(i.next())
        return v


# Your NestedIterator object will be instantiated and called as such:
# i, v = NestedIterator(nestedList), []
# while i.hasNext(): v.append(i.next())
# @lc code=end

