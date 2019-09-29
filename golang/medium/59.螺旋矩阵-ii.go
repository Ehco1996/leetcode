package medium

/*
 * @lc app=leetcode.cn id=59 lang=golang
 *
 * [59] 螺旋矩阵 II
 *
 * https://leetcode-cn.com/problems/spiral-matrix-ii/description/
 *
 * algorithms
 * Medium (73.97%)
 * Likes:    120
 * Dislikes: 0
 * Total Accepted:    16.1K
 * Total Submissions: 21.5K
 * Testcase Example:  '3'
 *
 * 给定一个正整数 n，生成一个包含 1 到 n^2 所有元素，且元素按顺时针顺序螺旋排列的正方形矩阵。
 *
 * 示例:
 *
 * 输入: 3
 * 输出:
 * [
 * ⁠[ 1, 2, 3 ],
 * ⁠[ 8, 9, 4 ],
 * ⁠[ 7, 6, 5 ]
 * ]
 *
 */

func generateMatrix(n int) [][]int {
	// 按照 l->r t->b r->l,b->t的顺序往里面添加数据
	// 每个方向更新完之后需要更新边界

	m := make([][]int, n)
	for i := 0; i < n; i++ {
		m[i] = make([]int, n)
	}

	// 初始化边界
	l, r, t, b := 0, n-1, 0, n-1

	for num := 1; num <= n*n; {
		for i := l; i < r+1; i++ {
			m[t][i] = num
			num++
		}
		t += 1

		for i := t; i < b+1; i++ {
			m[i][r] = num
			num++
		}
		r--

		for i := r; i > l-1; i-- {
			m[b][i] = num
			num++
		}
		b--

		for i := b; i > t-1; i-- {
			m[i][l] = num
			num++
		}
		l++
	}
	return m
}
