/*
 * @lc app=leetcode.cn id=918 lang=golang
 * @lcpr version=
 *
 * [918] 环形子数组的最大和
 */
package main

// @lc code=start
func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	total := nums[0]

	curMax, maxSum := nums[0], nums[0]
	curMin, minSum := nums[0], nums[0]

	for i := 1; i < n; i++ {
		num := nums[i]
		total += num

		// Kadane: 最大子数组
		curMax = max(num, curMax+num)
		maxSum = max(maxSum, curMax)

		// Kadane: 最小子数组
		curMin = min(num, curMin+num)
		minSum = min(minSum, curMin)
	}

	// 全为负数的特例
	if maxSum < 0 {
		return maxSum
	}

	return max(maxSum, total-minSum)
}

// @lc code=end

/*
// @lcpr case=start
// [1,-2,3,-2]\n
// @lcpr case=end

// @lcpr case=start
// [5,-3,5]\n
// @lcpr case=end

// @lcpr case=start
// [3,-2,2,-3]\n
// @lcpr case=end

*/
