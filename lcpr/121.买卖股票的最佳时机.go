package main

import "math"

/*
 * @lc app=leetcode.cn id=121 lang=golang
 * @lcpr version=
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) int {
	// 从前到后，维护最小值，维护最大值
	minv, maxv := math.MaxInt32, math.MinInt32

	for i := 0; i < len(prices); i++ {
		minv = min(minv, prices[i])
		maxv = max(maxv, prices[i]-minv)
	}

	return maxv
}

// @lc code=end

/*
// @lcpr case=start
// [7,1,5,3,6,4]\n
// @lcpr case=end

// @lcpr case=start
// [7,6,4,3,1]\n
// @lcpr case=end

*/
