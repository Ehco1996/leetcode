package easy

/*
 * @lc app=leetcode.cn id=232 lang=golang
 *
 * [232] 用栈实现队列
 *
 * https://leetcode-cn.com/problems/implement-queue-using-stacks/description/
 *
 * algorithms
 * Easy (60.69%)
 * Likes:    93
 * Dislikes: 0
 * Total Accepted:    18.6K
 * Total Submissions: 30.5K
 * Testcase Example:  '["MyQueue","push","push","peek","pop","empty"]\n[[],[1],[2],[],[],[]]'
 *
 * 使用栈实现队列的下列操作：
 *
 *
 * push(x) -- 将一个元素放入队列的尾部。
 * pop() -- 从队列首部移除元素。
 * peek() -- 返回队列首部的元素。
 * empty() -- 返回队列是否为空。
 *
 *
 * 示例:
 *
 * MyQueue queue = new MyQueue();
 *
 * queue.push(1);
 * queue.push(2);
 * queue.peek();  // 返回 1
 * queue.pop();   // 返回 1
 * queue.empty(); // 返回 false
 *
 * 说明:
 *
 *
 * 你只能使用标准的栈操作 -- 也就是只有 push to top, peek/pop from top, size, 和 is empty
 * 操作是合法的。
 * 你所使用的语言也许不支持栈。你可以使用 list 或者 deque（双端队列）来模拟一个栈，只要是标准的栈操作即可。
 * 假设所有操作都是有效的 （例如，一个空的队列不会调用 pop 或者 peek 操作）。
 *
 *
 */
type MyQueue struct {
	Stack []int
}

/** Initialize your data structure here. */
func Constructor() MyQueue {
	return MyQueue{}

}

/** Push element x to the back of queue. */
func (this *MyQueue) Push(x int) {
	this.Stack = append(this.Stack, x)

}

/** Removes the element from in front of queue and returns that element. */
func (this *MyQueue) Pop() int {
	if len(this.Stack) > 0 {
		res := this.Peek()
		this.Stack = this.Stack[1:]
		return res
	}
	return -1
}

/** Get the front element. */
func (this *MyQueue) Peek() int {
	if len(this.Stack) > 0 {
		return this.Stack[0]
	}
	return -1
}

/** Returns whether the queue is empty. */
func (this *MyQueue) Empty() bool {
	return len(this.Stack) == 0
}
