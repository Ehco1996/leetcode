package easy

/*
 * @lc app=leetcode.cn id=506 lang=golang
 *
 * [506] 相对名次
 *
 * https://leetcode-cn.com/problems/relative-ranks/description/
 *
 * algorithms
 * Easy (50.98%)
 * Likes:    27
 * Dislikes: 0
 * Total Accepted:    4.4K
 * Total Submissions: 8.5K
 * Testcase Example:  '[5,4,3,2,1]'
 *
 * 给出 N 名运动员的成绩，找出他们的相对名次并授予前三名对应的奖牌。前三名运动员将会被分别授予 “金牌”，“银牌” 和“ 铜牌”（"Gold
 * Medal", "Silver Medal", "Bronze Medal"）。
 *
 * (注：分数越高的选手，排名越靠前。)
 *
 * 示例 1:
 *
 *
 * 输入: [5, 4, 3, 2, 1]
 * 输出: ["Gold Medal", "Silver Medal", "Bronze Medal", "4", "5"]
 * 解释: 前三名运动员的成绩为前三高的，因此将会分别被授予 “金牌”，“银牌”和“铜牌” ("Gold Medal", "Silver Medal"
 * and "Bronze Medal").
 * 余下的两名运动员，我们只需要通过他们的成绩计算将其相对名次即可。
 *
 * 提示:
 *
 *
 * N 是一个正整数并且不会超过 10000。
 * 所有运动员的成绩都不相同。
 *
 *
 */

import "sort"
import "strconv"

func findRelativeRanks(nums []int) []string {
	bak := make([]int, len(nums))
	for idx, num := range nums {
		bak[idx] = num
	}

	sort.Sort(sort.Reverse(sort.IntSlice(bak)))
	h := make(map[int]int)
	for idx, num := range bak {
		h[num] = idx
	}
	res := make([]string, len(nums))
	for idx, num := range nums {
		var word string
		if h[num] == 0 {
			word = "Gold Medal"
		} else if h[num] == 1 {
			word = "Silver Medal"
		} else if h[num] == 2 {
			word = "Bronze Medal"
		} else {
			word = strconv.Itoa(h[num] + 1)
		}
		res[idx] = word
	}
	return res

}
