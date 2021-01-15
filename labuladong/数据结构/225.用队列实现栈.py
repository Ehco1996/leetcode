#
# @lc app=leetcode.cn id=225 lang=python3
#
# [225] 用队列实现栈
#

# @lc code=start
import queue


class MyStack:
    def __init__(self):
        """
        Initialize your data structure here.
        """
        self.queue = queue.SimpleQueue()
        self._top = 0

    def push(self, x: int) -> None:
        """
        Push element x onto stack.
        """
        self.queue.put_nowait(x)
        self._top = x

    def pop(self) -> int:
        """
        Removes the element on top of the stack and returns that element.
        将除了队列最里面的元素之外全部pop出来在塞进队列里
        这样原本第一个元素就变成了最后一个
        """
        size = self.queue.qsize()
        for _ in range(size - 2):
            self.queue.put_nowait(self.queue.get_nowait())
        self._top = self.queue.get_nowait()
        self.queue.put_nowait(self._top)
        return self.queue.get_nowait()

    def top(self) -> int:
        """
        Get the top element.
        """
        return self._top

    def empty(self) -> bool:
        """
        Returns whether the stack is empty.
        """
        return self.queue.empty()


# Your MyStack object will be instantiated and called as such:
# obj = MyStack()
# obj.push(x)
# param_2 = obj.pop()
# param_3 = obj.top()
# param_4 = obj.empty()
# @lc code=end

