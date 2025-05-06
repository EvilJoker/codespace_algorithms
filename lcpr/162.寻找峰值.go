/*
 * @lc app=leetcode.cn id=162 lang=golang
 * @lcpr version=
 *
 * [162] 寻找峰值
 */
package main

// @lc code=start
/*
思路：
1. 峰值定义：比相邻元素都大的元素
2. 使用二分查找优化：
   - 比较 mid 和 mid+1 元素
   - 如果 nums[mid] > nums[mid+1]，说明峰值在左侧（包含mid）
   - 如果 nums[mid] < nums[mid+1]，说明峰值在右侧
3. 最终 left == right 时，即为峰值位置
*/
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		if nums[mid] > nums[mid+1] {
			// 峰值在左侧（包含mid）
			right = mid
		} else {
			// 峰值在右侧
			left = mid + 1
		}
	}

	// 当 left == right 时，即为峰值位置
	return left
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,1]\n
// @lcpr case=end

// @lcpr case=start
// [1,2,1,3,5,6,4]\n
// @lcpr case=end

*/
