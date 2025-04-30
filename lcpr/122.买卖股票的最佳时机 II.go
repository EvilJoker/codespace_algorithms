/*
 * @lc app=leetcode.cn id=122 lang=golang
 * @lcpr version=
 *
 * [122] 买卖股票的最佳时机 II
 */
package main

// @lc code=start
func maxProfit(prices []int) int {
	// 思路，求递增变量
	maxp := 0
	for i := 1; i < len(prices); i++ {
		if prices[i] > prices[i-1] {
			maxp += prices[i] - prices[i-1]
		}
	}
	return maxp
}

// @lc code=end

/*
// @lcpr case=start
// [7,1,5,3,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,3,4,5]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

*/
