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
	//思路：子数组是连续的意思
	// 滑动窗口，右滑左缩，维持 min len

	min_len := math.MaxInt32

	left, end, n := 0, 0, len(nums)
	sum := 0

	for ; end < n; end++ {
		// 加入新元素
		sum += nums[end]
		// >=时更新
		for sum >= target {
			min_len = min(end-left+1, min_len) // 更新

			sum -= nums[left] // 缩减
			left++

		}

	}
	if min_len == math.MaxInt32 {
		return 0
	}

	return min_len
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
