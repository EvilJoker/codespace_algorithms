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
	// 思路：标记符号，如果 右边比自己大，那就减， 比如 IV
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
