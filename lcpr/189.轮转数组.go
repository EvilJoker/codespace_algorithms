/*
 * @lc app=leetcode.cn id=189 lang=golang
 * @lcpr version=
 *
 * [189] 轮转数组
 */

// @lc code=start
/*
思路：使用三次翻转实现数组轮转
1. 首先将整个数组翻转
2. 然后翻转前k个元素
3. 最后翻转剩余元素
时间复杂度：O(n)
空间复杂度：O(1)
*/

func rotate(nums []int, k int) {
	// 处理k大于数组长度的情况
	n := len(nums)
	k = k % n

	// 执行三次翻转
	reverse(nums, 0, n-1) // 整体翻转
	reverse(nums, 0, k-1) // 翻转前k个元素
	reverse(nums, k, n-1) // 翻转剩余元素
}

// reverse 函数用于翻转数组指定范围内的元素
func reverse(nums []int, start, end int) {
	for start < end {
		// 交换首尾元素
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4,5,6,7]\n3\n
// @lcpr case=end

// @lcpr case=start
// [-1,-100,3,99]\n2\n
// @lcpr case=end

*/

