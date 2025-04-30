/*
 * @lc app=leetcode.cn id=162 lang=golang
 * @lcpr version=
 *
 * [162] 寻找峰值
 */
package main

// @lc code=start
func findPeakElement(nums []int) int {
	//sl: 比左右大的元素就是峰值，
	//1. 顺序遍历
	//2. 优化二分法 ：比较 mid 和邻居，如果
	/*
		a. 比左右高就是 mid 就是峰值
		b. 如果num[ mid + 1] >num[ mid], 峰值在右
		c. 反之 峰值在左
	*/
	left, right := 0, len(nums)-1

	for left < right {

		mid := left + (right-left)/2

		if nums[mid] > nums[mid+1] {
			// 峰值在左侧（包含 mid）
			right = mid
		} else {
			// 峰值在右侧
			left = mid + 1
		}

	}

	// 最终 left ==right , 就是峰值索引
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
