// package golang

/*
 * @lc app=leetcode.cn id=447 lang=golang
 *
 * [447] 回旋镖的数量
 *
 * https://leetcode-cn.com/problems/number-of-boomerangs/description/
 *
 * algorithms
 * Easy (54.28%)
 * Likes:    58
 * Dislikes: 0
 * Total Accepted:    5.9K
 * Total Submissions: 10.8K
 * Testcase Example:  '[[0,0],[1,0],[2,0]]'
 *
 * 给定平面上 n 对不同的点，“回旋镖” 是由点表示的元组 (i, j, k) ，其中 i 和 j 之间的距离和 i 和 k
 * 之间的距离相等（需要考虑元组的顺序）。
 *
 * 找到所有回旋镖的数量。你可以假设 n 最大为 500，所有点的坐标在闭区间 [-10000, 10000] 中。
 *
 * 示例:
 *
 *
 * 输入:
 * [[0,0],[1,0],[2,0]]
 *
 * 输出:
 * 2
 *
 * 解释:
 * 两个回旋镖为 [[1,0],[0,0],[2,0]] 和 [[1,0],[2,0],[0,0]]
 *
 *
 */
func dist(x,y []int) int{
	return (x[0]-y[0])*(x[0]-y[0]) + (x[1]-y[1])*(x[1]-y[1])
}


 func numberOfBoomerangs(points [][]int) int {
	//两重循环，用 hash 表记录出现过的距离
	//遍历每个点和它所有点的距离，用 hash 表记录，如果出现重复的，说明可以形成“回旋镖”

	count := 0
	for i:=0;i<len(points);i++{
		// k:distance v: number of the point
		dis := make(map[int]int)
		for j:=0;j<len(points);j++{
			if j!=i{
				dis[dist(points[i],points[j])] +=1
			}
		}
		for _,v := range dis{
			//计算排列组合公式 n * (n - 1)
			count += v*(v-1)
		}
	}
	return count
}
