/*
 * @lc app=leetcode.cn id=167 lang=golang
 * @lcpr version=
 *
 * [167] 两数之和 II - 输入有序数组
 */
package main

// @lc code=start
func twoSum(numbers []int, target int) []int {

	// 思路: 双指针
	n := len(numbers)
	left, right := 0, n-1
	for left < right {
		if numbers[left]+numbers[right] == target {
			return []int{left + 1, right + 1}
		} else if numbers[left]+numbers[right] < target {
			left++
		} else {
			right--
		}
	}

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
