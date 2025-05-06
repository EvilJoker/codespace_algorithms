/*
 * @lc app=leetcode.cn id=80 lang=golang
 * @lcpr version=
 *
 * [80] 删除有序数组中的重复项 II
 */
package main

// @lc code=start
func removeDuplicates(nums []int) int {
	// 不能和上一位和上上一位相等, 特殊点需要和 p 比较

	n := len(nums)
	if n < 3 {
		return n
	}

	p := 2

	for i := 2; i < n; i++ {
		if nums[i] != nums[p-2] { //
			nums[p] = nums[i] // 需要交换
			p++
		}
	}

	return p

}

// @lc code=end

/*
// @lcpr case=start
// [1,1,1,2,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,1,1,1,1,2,3,3]\n
// @lcpr case=end

*/
