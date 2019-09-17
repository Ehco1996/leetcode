package easy

/*
 * @lc app=leetcode.cn id=204 lang=golang
 *
 * [204] 计数质数
 *
 * https://leetcode-cn.com/problems/count-primes/description/
 *
 * algorithms
 * Easy (29.34%)
 * Likes:    195
 * Dislikes: 0
 * Total Accepted:    25.9K
 * Total Submissions: 87K
 * Testcase Example:  '10'
 *
 * 统计所有小于非负整数 n 的质数的数量。
 *
 * 示例:
 *
 * 输入: 10
 * 输出: 4
 * 解释: 小于 10 的质数一共有 4 个, 它们是 2, 3, 5, 7 。
 *
 *
 */

func countPrimes(n int) int {
	if n < 2 {
		return 0
	}
	c := 0
	lastPrime := 2
	for num := 2; num < n; num++ {
		if isPrimesCacheVersion(num, lastPrime) {
			lastPrime = num
			c += 1
		}
	}

	return c
}

func isPrimesCacheVersion(n, lastPrime int) bool {
	// 质数（Prime number），又称素数，
	// 指在大于1的自然数中，除了1和该数自身外，无法被其他自然数整除的数
	// 只需要判段 1 ~ sqrt(n)之间的数是能被整除就可以判断了
	sqrt := int(math.Ceil(math.Sqrt(float64(n))))
	var i int
	if lastPrime>sqrt{
		i=2
	}else{
		i=lastPrime
	}
	for ;i <= sqrt; i++ {
		if n%i == 0 && i!=n {
			return false
		}
	}
	return true
}

// func main() {
// 	fmt.Println(countPrimes(2))
// 	// fmt.Println(isPrimesCacheVersion(2,2))
// }
