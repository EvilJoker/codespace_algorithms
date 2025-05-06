/*
 * @lc app=leetcode.cn id=27 lang=golang
 * @lcpr version=
 *
 * [27] 移除元素
 */
package main

// @lc code=start
func removeElement(nums []int, val int) int {
	// 解题思路：
	// 1. 使用双指针技巧，p指向当前要保存元素的位置
	// 2. 遍历数组，当遇到不等于val的元素时，将其复制到p位置
	// 3. 最后返回不等于val的元素个数
	// 时间复杂度：O(n)，空间复杂度：O(1)

	// p指向当前要保存元素的位置
	p := 0

	// 遍历数组
	for i := 0; i < len(nums); i++ {
		// 如果当前元素不等于val
		if nums[i] != val {
			// 将当前元素复制到p位置
			nums[p] = nums[i]
			// p指针后移
			p++
		}
	}

	// 返回不等于val的元素个数
	return p
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,2,3]\n3\n
// @lcpr case=end

// @lcpr case=start
// [0,1,2,2,3,0,4,2]\n2\n
// @lcpr case=end

*/
