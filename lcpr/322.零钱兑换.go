/*
 * @lc app=leetcode.cn id=322 lang=golang
 * @lcpr version=
 *
 * [322] 零钱兑换
 */

// @lc code=start
func coinChange(coins []int, amount int) int {

	/* 思路，动态规划,类似单词拆分
	1. dp[i] 代表 amount 为 i 时，最少个数
	2. 初始化 dp[i]=amout +1，到达不了的数, dp[0]=0
	3. 递推： 遍历coins 如果 dp[i-coin]!=-1, dp[i]在所有硬币间取最小值。

	*/
	// 初始化
	dp := make([]int, amount+1)
	for i := 0; i < amount+1; i++ {
		dp[i] = amount + 1
	}
	// 特殊点，如果 能追溯到 dp[n-1],说明是可用的
	dp[0] = 0

	// 递归
	for i := 1; i < amount+1; i++ {

		for _, coin := range coins {
			// 试图找到最少得硬币数。
			if i-coin >= 0 {
				dp[i] = min(dp[i], dp[i-coin]+1) // 如果之前不能凑整，值就是amout +1,
			}
		}

	}
	// 找不到
	if dp[amount] == amount+1 {
		return -1
	}

	return dp[amount]

}

// @lc code=end

/*
// @lcpr case=start
// [1, 2, 5]\n11\n
// @lcpr case=end

// @lcpr case=start
// [2]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n0\n
// @lcpr case=end

*/

