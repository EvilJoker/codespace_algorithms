/*
 * @lc app=leetcode.cn id=198 lang=golang
 * @lcpr version=
 *
 * [198] 打家劫舍
 */
package main

// @lc code=start
func rob(nums []int) int {
	// 思路：动态规划
	/*	dp[i] 标识，目前到 房间i 为止的最大金额
		初始值：
		1. i=0, dp[i] = w[i], 一个房间只有一个
		2. i=1, dp[i] = max(1,2) 前两个房间选最大
		递归表达：
		如果偷当前的。dp[i] = dp[i-2] + w[i],不偷 dp[i]= dp[i-1]
		dp[i]= max(两个最大值
	*/
	n := len(nums)
	if n == 1 {
		return nums[0]
	}

	dp := make([]int, n)
	dp[0] = nums[0]
	dp[1] = max(nums[0], nums[1])

	if n == 2 {
		return dp[1]
	}

	// 递归

	for i := 2; i < n; i++ {
		dp[i] = max(dp[i-2]+nums[i], dp[i-1])

	}

	return dp[n-1]
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [2,7,9,3,1]\n
// @lcpr case=end

*/
