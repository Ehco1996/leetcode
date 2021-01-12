#
# @lc app=leetcode.cn id=295 lang=python3
#
# [295] 数据流的中位数
#

# @lc code=start
import heapq


class MinHeap:
    def __init__(self):
        self.nums = []

    @property
    def size(self):
        return len(self.nums)

    def pop(self):
        return heapq.heappop(self.nums)

    def peek(self):
        return heapq.nsmallest(1, self.nums)[0]

    def push(self, num):
        return heapq.heappush(self.nums, num)


class MaxHeap:
    def __init__(self):
        self.nums = []

    @property
    def size(self):
        return len(self.nums)

    def pop(self):
        return heapq.heappop(self.nums) * -1

    def peek(self):
        return heapq.nsmallest(1, self.nums)[0] * -1

    def push(self, num):
        return heapq.heappush(self.nums, num * -1)


class MedianFinder:
    def __init__(self):
        """
        将nums拆分为两个堆
            大顶堆：数据从小到大排列
            小顶堆: 数据从大到小排列

        每次加入数据时，选择两堆之一加入(优先往大顶堆里加)，
                并且保证
                1. 两个堆是有序的，并且两堆元素数量相差不超过1
                2. large堆的堆顶元素要大于等于small堆的堆顶元素
        这样中位数就是
            1.偶数 两堆堆顶的平均值
            2 奇数 元素数量多的那个堆顶元素
        """
        self.large = MinHeap()
        self.small = MaxHeap()

    def addNum(self, num: int) -> None:

        if self.small.size >= self.large.size:
            # small较大就要往large里加元素了
            self.small.push(num)
            self.large.push(self.small.pop())
        else:
            self.large.push(num)
            self.small.push(self.large.pop())

    def findMedian(self) -> float:
        if self.small.size == self.large.size:
            return (self.small.peek() + self.large.peek()) / 2.0
        elif self.small.size > self.large.size:
            return self.small.peek()
        else:
            return self.large.peek()

    @classmethod
    def test(cls):
        m = cls()
        m.addNum(1)
        m.addNum(2)
        m.addNum(3)
        print(m.findMedian())

        m = cls()
        m.addNum(-1)
        m.addNum(-2)
        m.addNum(-3)
        print(m.findMedian())


# Your MedianFinder object will be instantiated and called as such:
# obj = MedianFinder()
# obj.addNum(num)
# param_2 = obj.findMedian()
# @lc code=end
