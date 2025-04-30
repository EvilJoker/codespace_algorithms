/*
 * @lc app=leetcode.cn id=35 lang=golang
 * @lcpr version=
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {

	/*sl 二分*/
	n := len(nums)
	left, right := 0, n-1
	mid := (left-right)/2 + right
	for left <= right {
		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (left-right)/2 + right
	}
	// 最终找到区间。将left 插入
	return left

}

// @lc code=end

/*
// @lcpr case=start
// [1,3,5,6]\n5\n
// @lcpr case=end

// @lcpr case=start
// [1,3,5,6]\n2\n
// @lcpr case=end

// @lcpr case=start
// [1,3,5,6]\n7\n
// @lcpr case=end

*/

