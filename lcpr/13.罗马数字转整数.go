/*
 * @lc app=leetcode.cn id=13 lang=golang
 * @lcpr version=
 *
 * [13] 罗马数字转整数
 */
package main

// @lc code=start

var dict = map[byte]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	// 解题思路：
	// 1. 遍历字符串，使用字典记录每个罗马数字对应的值
	// 2. 如果当前罗马数字比下一个罗马数字小，则减去当前值，否则加上当前值
	// 3. 最后加上最后一个罗马数字的值
	// 时间复杂度：O(n)，空间复杂度：O(1)

	// 遍历字符串，使用字典记录每个罗马数字对应的值
	n := len(s)

	sum := 0
	for i := 0; i < n-1; i++ {
		right, _ := dict[s[i+1]]
		cur, _ := dict[s[i]]

		if cur < right {
			sum -= cur
		} else {
			sum += cur
		}
	}
	t, _ := dict[s[n-1]]
	sum += t

	return sum

}

// @lc code=end

/*
// @lcpr case=start
// "III"\n
// @lcpr case=end

// @lcpr case=start
// "IV"\n
// @lcpr case=end

// @lcpr case=start
// "IX"\n
// @lcpr case=end

// @lcpr case=start
// "LVIII"\n
// @lcpr case=end

// @lcpr case=start
// "MCMXCIV"\n
// @lcpr case=end

*/
