/*
 * @lc app=leetcode.cn id=300 lang=golang
 * @lcpr version=
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	/* 动态规划解法思路：
	1. 定义dp数组：dp[i]表示以nums[i]结尾的最长递增子序列的长度
	2. 初始化：dp[i] = 1，每个元素自身就是一个长度为1的递增子序列
	3. 状态转移：
	   - 对于每个位置i，遍历其之前的所有位置j
	   - 如果nums[j] < nums[i]，则可以将nums[i]接在nums[j]后面
	   - dp[i] = max(dp[i], dp[j] + 1)
	4. 最终结果：返回dp数组中的最大值
	*/
	n := len(nums)
	if n == 0 {
		return 0
	}

	dp := make([]int, n)
	for i := range dp {
		dp[i] = 1
	}

	maxLen := 1

	for i := 1; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)
			}
		}
		maxLen = max(maxLen, dp[i])
	}

	return maxLen

}

// @lc code=end

/*
// @lcpr case=start
// [10,9,2,5,3,7,101,18]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,0,3,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [7,7,7,7,7,7,7]\n
// @lcpr case=end

*/

