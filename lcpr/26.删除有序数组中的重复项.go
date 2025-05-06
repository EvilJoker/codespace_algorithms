/*
 * @lc app=leetcode.cn id=26 lang=golang
 * @lcpr version=
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	// 解题思路：
	// 1. 使用双指针技巧，p指向当前不重复元素的位置
	// 2. 遍历数组，当遇到与p-1位置不同的元素时，将其复制到p位置
	// 3. 最后返回不重复元素的个数
	// 时间复杂度：O(n)，空间复杂度：O(1)

	// 处理空数组的情况
	if len(nums) == 0 {
		return 0
	}

	// p指向当前不重复元素的位置，从1开始因为第一个元素一定是不重复的
	p := 1

	// 从第二个元素开始遍历数组
	for i := 1; i < len(nums); i++ {
		// 如果当前元素与前一个不重复元素不同
		if nums[i] != nums[p-1] {
			// 将当前元素复制到p位置
			nums[p] = nums[i]
			// p指针后移
			p++
		}
	}

	// 返回不重复元素的个数
	return p
}

// @lc code=end

/*
// @lcpr case=start
// [1,1,2]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,1,1,1,2,2,3,3,4]\n
// @lcpr case=end

*/

