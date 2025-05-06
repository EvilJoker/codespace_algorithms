/*
 * @lc app=leetcode.cn id=35 lang=golang
 * @lcpr version=
 *
 * [35] 搜索插入位置
 */
package main

// @lc code=start
// 二分查找
// 思路：
// 1. 使用二分查找在有序数组中查找目标值
// 2. 如果找到目标值，直接返回其索引
// 3. 如果没找到，left 指针会指向第一个大于目标值的位置，这就是应该插入的位置
// 4. 时间复杂度：O(log n)，空间复杂度：O(1)
func searchInsert(nums []int, target int) int {
	n := len(nums)
	left, right := 0, n-1

	for left <= right {
		mid := left + (right-left)/2 // 防止整数溢出的写法

		if nums[mid] == target {
			return mid
		}

		if target > nums[mid] {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}

	return left // 返回插入位置
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
