/*
 * @lc app=leetcode.cn id=167 lang=golang
 * @lcpr version=
 *
 * [167] 两数之和 II - 输入有序数组
 */
package main

// @lc code=start
// 思路：由于数组是有序的，可以使用双指针技巧
// 1. 左指针从数组开始，右指针从数组末尾
// 2. 如果两数之和等于目标值，返回两个指针位置（注意题目要求从1开始计数）
// 3. 如果两数之和小于目标值，左指针右移（因为需要更大的数）
// 4. 如果两数之和大于目标值，右指针左移（因为需要更小的数）
func twoSum(numbers []int, target int) []int {
	// 初始化左右指针
	left, right := 0, len(numbers)-1

	// 当左右指针相遇时结束循环
	for left < right {
		sum := numbers[left] + numbers[right]

		if sum == target {
			// 找到目标值，返回位置（从1开始计数）
			return []int{left + 1, right + 1}
		} else if sum < target {
			// 和小于目标值，左指针右移
			left++
		} else {
			// 和大于目标值，右指针左移
			right--
		}
	}

	// 未找到解
	return []int{}
}

// @lc code=end

/*
// @lcpr case=start
// [2,7,11,15]\n9\n
// @lcpr case=end

// @lcpr case=start
// [2,3,4]\n6\n
// @lcpr case=end

// @lcpr case=start
// [-1,0]\n-1\n
// @lcpr case=end

*/
