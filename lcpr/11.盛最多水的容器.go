/*
 * @lc app=leetcode.cn id=11 lang=golang
 * @lcpr version=
 *
 * [11] 盛最多水的容器
 */
package main

import (
	"math"
)

// @lc code=start
func maxArea(height []int) int {
	// 思路： 贪心，总是向比较高的一段移动

	n := len(height)
	left, end := 0, n-1

	max_area := math.MinInt32

	for left < end {
		max_area = max(max_area, (end-left)*min(height[left], height[end]))
		if height[left] < height[end] {
			left++
		} else {
			end--
		}
	}

	return max_area

}

// @lc code=end

/*
// @lcpr case=start
// [1,8,6,2,5,4,8,3,7]\n
// @lcpr case=end

// @lcpr case=start
// [1,1]\n
// @lcpr case=end

*/
