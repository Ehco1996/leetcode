#
# @lc app=leetcode.cn id=232 lang=python3
#
# [232] 用栈实现队列
#

# @lc code=start
from collections import deque


class MyQueue:
    """
    用两个栈实现队列
    s1,s2
    push就不停的往s1里添加元素
    当s2为空时，
    可以把s1的所有元素取出再添加进s2，这时候s2中元素就是先进先出顺序了
    """

    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.s1 = deque()
        self.s2 = deque()

    def push(self, x: int) -> None:
        """
        Push element x to the back of queue.
        """
        self.s1.append(x)

    def pop(self) -> int:
        """
        Removes the element from in front of queue and returns that element.
        """
        self.peek()
        return self.s2.pop()

    def _full_s2(self) -> int:
        """
        Get the front element.
        """
        if len(self.s2) == 0:
            while len(self.s1) > 0:
                self.s2.append(self.s1.pop())

    def peek(self):
        self._full_s2()
        x = self.s2.pop()
        self.s2.append(x)
        return x

    def empty(self) -> bool:
        """
        Returns whether the queue is empty.
        """
        return (len(self.s1) + len(self.s2)) == 0


# Your MyQueue object will be instantiated and called as such:
# obj = MyQueue()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.peek()
# param_4 = obj.empty()
# @lc code=end

