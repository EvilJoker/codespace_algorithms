/*
 * @lc app=leetcode.cn id=300 lang=golang
 * @lcpr version=
 *
 * [300] 最长递增子序列
 */

// @lc code=start
func lengthOfLIS(nums []int) int {
	// 思路：动态规划
	n := len(nums)

	// 目前为止最长的子序列，初始值 dp[0]=1
	dp := make([]int, n)
	for i := 0; i < n; i++ {
		dp[i] = 1
	}

	maxlen := 1

	// 递推 遍历所有前面的数，找到小于当前数的数，num[i-1]> 0..i-2; dp[i]= max(dp[i],dp[i-x]+1)

	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = max(dp[i], dp[j]+1)

				maxlen = max(dp[i], maxlen)
			}
		}
	}

	return maxlen

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

