/*
 * @lc app=leetcode.cn id=209 lang=golang
 * @lcpr version=
 *
 * [209] 长度最小的子数组
 */
package main

import "math"

// @lc code=start
func minSubArrayLen(target int, nums []int) int {
	// 解题思路：
	// 1. 使用滑动窗口算法，维护一个窗口[left, right]
	// 2. 右指针不断向右移动，扩大窗口，累加和
	// 3. 当窗口和 >= target 时，更新最小长度，并尝试左指针右移缩小窗口
	// 4. 重复步骤2-3直到遍历完整个数组

	minLength := math.MaxInt32
	left, right := 0, 0
	windowSum := 0

	// 遍历数组，移动右指针
	for right < len(nums) {
		// 扩大窗口，累加和
		windowSum += nums[right]

		// 当窗口和满足条件时，尝试缩小窗口
		for windowSum >= target {
			// 更新最小长度
			minLength = min(right-left+1, minLength)
			// 缩小窗口，左指针右移
			windowSum -= nums[left]
			left++
		}
		right++
	}

	// 如果没有找到满足条件的子数组，返回0
	if minLength == math.MaxInt32 {
		return 0
	}
	return minLength
}

// @lc code=end

/*
// @lcpr case=start
// 7\n[2,3,1,2,4,3]\n
// @lcpr case=end

// @lcpr case=start
// 4\n[1,4,4]\n
// @lcpr case=end

// @lcpr case=start
// 11\n[1,1,1,1,1,1,1,1]\n
// @lcpr case=end

*/
