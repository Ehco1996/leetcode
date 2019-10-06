/*
 * @lc app=leetcode.cn id=295 lang=golang
 *
 * [295] 数据流的中位数
 *
 * https://leetcode-cn.com/problems/find-median-from-data-stream/description/
 *
 * algorithms
 * Hard (36.87%)
 * Likes:    77
 * Dislikes: 0
 * Total Accepted:    6.1K
 * Total Submissions: 15.4K
 * Testcase Example:  '["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]\n' +
  '[[],[1],[2],[],[3],[]]'
 *
 * 中位数是有序列表中间的数。如果列表长度是偶数，中位数则是中间两个数的平均值。
 *
 * 例如，
 *
 * [2,3,4] 的中位数是 3
 *
 * [2,3] 的中位数是 (2 + 3) / 2 = 2.5
 *
 * 设计一个支持以下两种操作的数据结构：
 *
 *
 * void addNum(int num) - 从数据流中添加一个整数到数据结构中。
 * double findMedian() - 返回目前所有元素的中位数。
 *
 *
 * 示例：
 *
 * addNum(1)
 * addNum(2)
 * findMedian() -> 1.5
 * addNum(3)
 * findMedian() -> 2
 *
 * 进阶:
 *
 *
 * 如果数据流中所有整数都在 0 到 100 范围内，你将如何优化你的算法？
 * 如果数据流中 99% 的整数都在 0 到 100 范围内，你将如何优化你的算法？
 *
 *
*/

// @lc code=start
type MedianFinder struct {
	Nums []int
}

func Constructor() MedianFinder {
	return MedianFinder{Nums: []int{}}

}

func (this *MedianFinder) AddNum(num int) {
	// 插入排序
	if len(this.Nums) == 0 || this.Nums[len(this.Nums)-1] <= num {
		this.Nums = append(this.Nums, num)
		return
	}

	var idx int
	for i := 0; i < len(this.Nums); i++ {
		if this.Nums[i] <= num && this.Nums[i+1] >= num {
			idx = i + 1
			break
		}
	}
	this.Nums = append(this.Nums, 0)
	copy(this.Nums[idx+1:], this.Nums[idx:])
	this.Nums[idx] = num
}

func (this *MedianFinder) FindMedian() float64 {
	n := len(this.Nums)
	mid := n / 2
	if n%2 == 1 {
		return float64(this.Nums[mid])
	}
	return float64(this.Nums[mid]+this.Nums[mid-1]) / 2.0
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */
// @lc code=end

