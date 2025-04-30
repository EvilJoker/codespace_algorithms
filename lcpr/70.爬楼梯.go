/*
 * @lc app=leetcode.cn id=70 lang=golang
 * @lcpr version=
 *
 * [70] 爬楼梯
 */
package main

// @lc code=start
func climbStairs(n int) int {
	if n <= 1 {
		return 1
	}

	if n <= 2 {
		return 2
	}
	// 斐波那契

	dp := make([]int, n)
	dp[0], dp[1] = 1, 2

	for i := 2; i < n; i++ {
		dp[i] = dp[i-1] + dp[i-2]
	}
	return dp[n-1]
}

// @lc code=end

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 3\n
// @lcpr case=end

*/
