package main
import "fmt"
/*
 * @lc app=leetcode.cn id=205 lang=golang
 *
 * [205] 同构字符串
 *
 * https://leetcode-cn.com/problems/isomorphic-strings/description/
 *
 * algorithms
 * Easy (45.51%)
 * Likes:    131
 * Dislikes: 0
 * Total Accepted:    16.2K
 * Total Submissions: 35.4K
 * Testcase Example:  '"egg"\n"add"'
 *
 * 给定两个字符串 s 和 t，判断它们是否是同构的。
 *
 * 如果 s 中的字符可以被替换得到 t ，那么这两个字符串是同构的。
 *
 * 所有出现的字符都必须用另一个字符替换，同时保留字符的顺序。两个字符不能映射到同一个字符上，但字符可以映射自己本身。
 *
 * 示例 1:
 *
 * 输入: s = "egg", t = "add"
 * 输出: true
 *
 *
 * 示例 2:
 *
 * 输入: s = "foo", t = "bar"
 * 输出: false
 *
 * 示例 3:
 *
 * 输入: s = "paper", t = "title"
 * 输出: true
 *
 * 说明:
 * 你可以假设 s 和 t 具有相同的长度。
 *
 */

func isIsomorphic(s string, t string) bool {
	h1 := make(map[string]int)
	pattern1 := make([]int, len(s))
	curFlag := 0
	// gen s pattern
	for idx, c := range s {
		if p, ok := h1[string(c)]; !ok {
			pattern1[idx] = curFlag
			h1[string(c)] = curFlag
			curFlag++
		} else {
			pattern1[idx] = p
		}
	}
	// fmt.Println(pattern1)

	h2 := make(map[string]int)
	curFlag = 0
	for idx,c :=range t{
		_, ok := h2[string(c)]
		if !ok {
			h2[string(c)] = curFlag
			curFlag++
		}


		if h2[string(c)] != pattern1[idx]{
			// fmt.Println(h2,idx,c)
			return false
		}
	}
	return true

}

// func main(){
// 	s :="aa"
// 	t :="ab"
// 	fmt.Println(isIsomorphic(s,t))
// }