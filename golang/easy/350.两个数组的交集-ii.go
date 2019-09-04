package main

import "fmt"

/*
 * @lc app=leetcode.cn id=350 lang=golang
 *
 * [350] 两个数组的交集 II
 *
 * https://leetcode-cn.com/problems/intersection-of-two-arrays-ii/description/
 *
 * algorithms
 * Easy (43.22%)
 * Likes:    183
 * Dislikes: 0
 * Total Accepted:    45.3K
 * Total Submissions: 103.1K
 * Testcase Example:  '[1,2,2,1]\n[2,2]'
 *
 * 给定两个数组，编写一个函数来计算它们的交集。
 *
 * 示例 1:
 *
 * 输入: nums1 = [1,2,2,1], nums2 = [2,2]
 * 输出: [2,2]
 *
 *
 * 示例 2:
 *
 * 输入: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
 * 输出: [4,9]
 *
 * 说明：
 *
 *
 * 输出结果中每个元素出现的次数，应与元素在两个数组中出现的次数一致。
 * 我们可以不考虑输出结果的顺序。
 *
 *
 * 进阶:
 *
 *
 * 如果给定的数组已经排好序呢？你将如何优化你的算法？
 * 如果 nums1 的大小比 nums2 小很多，哪种方法更优？
 * 如果 nums2 的元素存储在磁盘上，磁盘内存是有限的，并且你不能一次加载所有的元素到内存中，你该怎么办？
 *
 *
 */
func intersectsort(nums1 []int, nums2 []int) []int {
	h := make(map[int]int)

	for i := 0; i < len(nums1); {
		j := i
		for ; nums1[i] > nums2[j] && j < len(nums2); j++ {
		}
		num1 := nums1[i]
		num2 := nums2[j]
		if num1 == num2 {
			h[num1] = 0
		}
		i++
	}

	res := []int{}
	for num, _ := range h {
		res = append(res, num)
	}

	return res
}

func intersect(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	for _, n := range nums1 {
		if count, ok := m[n]; !ok {
			m[n] = 1
		} else {
			m[n] = count + 1
		}
	}

	var res []int
	for _, n := range nums2 {
		if count, ok := m[n]; ok && count > 0 {
			res = append(res, n)
			m[n] = count - 1
		}
	}
	return res
}

// func main() {
// 	nums1 := []int{2, 2, 3, 4, 5, 6}
// 	nums2 := []int{1, 2, 2, 3, 9, 9, 9, 9, 9}
// 	fmt.Println(intersect(nums1, nums2))
// }
