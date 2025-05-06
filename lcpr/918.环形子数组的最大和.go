/*
 * @lc app=leetcode.cn id=918 lang=golang
 * @lcpr version=
 *
 * [918] 环形子数组的最大和
 */
package main

// @lc code=start
// 思路：
// 1. 环形数组的最大子数组和有两种情况：
//    - 普通的最大子数组和（不跨越首尾）
//    - 跨越首尾的情况，此时最大和 = 总和 - 最小子数组和
// 2. 使用 Kadane 算法分别计算最大子数组和和最小子数组和
// 3. 如果最大子数组和小于0，说明数组全为负数，直接返回最大子数组和
// 4. 否则返回两种情况中的较大值

func maxSubarraySumCircular(nums []int) int {
	n := len(nums)
	// 计算数组总和
	total := nums[0]

	// 初始化最大子数组和相关的变量
	currMax, maxSum := nums[0], nums[0]
	// 初始化最小子数组和相关的变量
	currMin, minSum := nums[0], nums[0]

	// 遍历数组，同时计算最大子数组和和最小子数组和
	for i := 1; i < n; i++ {
		num := nums[i]
		total += num

		// 使用 Kadane 算法计算最大子数组和
		currMax = max(num, currMax+num)
		maxSum = max(maxSum, currMax)

		// 使用 Kadane 算法计算最小子数组和
		currMin = min(num, currMin+num)
		minSum = min(minSum, currMin)
	}

	// 处理全为负数的情况
	if maxSum < 0 {
		return maxSum
	}

	// 返回两种情况中的较大值
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
