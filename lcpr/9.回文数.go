/*
 * @lc app=leetcode.cn id=9 lang=golang
 * @lcpr version=
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	// 思路：先转成字符串，然后进行遍历
	s := strconv.Itoa(x)
	n := len(s)

	for i := 0; i < n/2; i++ {
		if s[i] != s[n-i-1] {
			return false
		}
	}

	return true

}

// @lc code=end

/*
// @lcpr case=start
// 121\n
// @lcpr case=end

// @lcpr case=start
// -121\n
// @lcpr case=end

// @lcpr case=start
// 10\n
// @lcpr case=end

*/

