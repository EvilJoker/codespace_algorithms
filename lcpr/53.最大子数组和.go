/*
 * @lc app=leetcode.cn id=53 lang=golang
 * @lcpr version=
 *
 * [53] 最大子数组和
 */
package main

// @lc code=start
func maxSubArray(nums []int) int {
	// 使用动态规划的思想
	// maxSum: 记录全局最大子数组和
	// curSum: 记录当前连续子数组的和

	// 初始化：第一个元素作为初始值
	maxSum := nums[0]
	curSum := nums[0]

	// 从第二个元素开始遍历
	for i := 1; i < len(nums); i++ {
		// 对于每个位置，有两种选择：
		// 1. 将当前数字加入现有子数组
		// 2. 以当前数字开始新的子数组
		// 选择两者中的较大值
		curSum = max(nums[i], curSum+nums[i])

		// 更新全局最大和
		maxSum = max(curSum, maxSum)
	}

	return maxSum
}

// @lc code=end

/*
// @lcpr case=start
// [-2,1,-3,4,-1,2,1,-5,4]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

// @lcpr case=start
// [5,4,-1,7,8]\n
// @lcpr case=end

*/
