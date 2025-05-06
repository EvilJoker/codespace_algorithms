/*
 * @lc app=leetcode.cn id=11 lang=golang
 * @lcpr version=
 *
 * [11] 盛最多水的容器
 */
package main

// @lc code=start
func maxArea(height []int) int {
	// 解题思路：双指针法
	// 1. 使用左右两个指针从数组两端向中间移动
	// 2. 每次移动高度较小的指针，因为移动高度较大的指针一定会导致面积减小
	// 3. 在移动过程中记录最大面积
	// 时间复杂度：O(n)，空间复杂度：O(1)

	left, right := 0, len(height)-1
	maxArea := 0

	for left < right {
		// 计算当前容器的面积
		area := (right - left) * min(height[left], height[right])
		maxArea = max(maxArea, area)

		// 移动高度较小的指针
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxArea
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
