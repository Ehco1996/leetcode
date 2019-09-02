package main

import "fmt"

/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 *
 * https://leetcode-cn.com/problems/min-stack/description/
 *
 * algorithms
 * Easy (49.33%)
 * Likes:    270
 * Dislikes: 0
 * Total Accepted:    41.2K
 * Total Submissions: 82.9K
 * Testcase Example:  '["MinStack","push","push","push","getMin","pop","top","getMin"]\n' +
  '[[],[-2],[0],[-3],[],[],[],[]]'
 *
 * 设计一个支持 push，pop，top 操作，并能在常数时间内检索到最小元素的栈。
 *
 *
 * push(x) -- 将元素 x 推入栈中。
 * pop() -- 删除栈顶的元素。
 * top() -- 获取栈顶元素。
 * getMin() -- 检索栈中的最小元素。
 *
 *
 * 示例:
 *
 * MinStack minStack = new MinStack();
 * minStack.push(-2);
 * minStack.push(0);
 * minStack.push(-3);
 * minStack.getMin();   --> 返回 -3.
 * minStack.pop();
 * minStack.top();      --> 返回 0.
 * minStack.getMin();   --> 返回 -2.
 *
 *
*/
type MinStack struct {
	Stack []int
	Min   []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		Stack: []int{},
		Min:   []int{},
	}

}

func (this *MinStack) Push(x int) {

	this.Stack = append([]int{x}, this.Stack...)

	if len(this.Min) == 0 || x <=this.Min[0] {
		this.Min = append([]int{x}, this.Min...)
	}

}

func (this *MinStack) Pop() {
	if len(this.Stack) > 0 {
		res := this.Stack[0]
		this.Stack = this.Stack[1:]
		if res == this.GetMin() {
			this.Min = this.Min[1:]
		}
	}
}

func (this *MinStack) Top() int {
	if len(this.Stack) > 0 {
		return this.Stack[0]
	}
	return 0
}

func (this *MinStack) GetMin() int {
	if len(this.Min) > 0 {
		return this.Min[0]
	}
	return 0
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(x);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

// func main() {
// 	obj := Constructor()
// 	obj.Push(2)
// 	obj.Push(0)
// 	obj.Push(3)
// 	obj.Push(0)
//     fmt.Println(obj.GetMin())
//     obj.Pop()
//     fmt.Println(obj.GetMin())
//     obj.Pop()
//     fmt.Println(obj.GetMin())
//     obj.Pop()
//     fmt.Println(obj.GetMin())
    
//     fmt.Println(obj.Top())
// 	fmt.Println(obj.GetMin())
// 	fmt.Println(obj.Top())

// }
