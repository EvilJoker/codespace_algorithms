/*
 * @lc app=leetcode.cn id=9 lang=golang
 * @lcpr version=
 *
 * [9] 回文数
 */

// @lc code=start
func isPalindrome(x int) bool {
	// 解题思路：
	// 1. 将整数转换为字符串，便于逐位比较
	// 2. 使用双指针从两端向中间移动，比较对应位置的字符
	// 3. 如果发现不相等，则不是回文数
	// 4. 时间复杂度：O(log n)，空间复杂度：O(log n)

	// 将整数转换为字符串
	s := strconv.Itoa(x)
	n := len(s)

	// 双指针从两端向中间移动，比较对应位置的字符
	for i := 0; i < n/2; i++ {
		// 如果对应位置字符不相等，则不是回文数
		if s[i] != s[n-i-1] {
			return false
		}
	}

	// 所有位置都相等，是回文数
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

