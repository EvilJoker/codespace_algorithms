/*
 * @lc app=leetcode.cn id=189 lang=golang
 * @lcpr version=
 *
 * [189] 轮转数组
 */

// @lc code=start
func rotate(nums []int, k int) {
	// 思路：三次翻转，整体翻转，前后再次翻转
	n := len(nums)
	k = k % n

	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)

}
func reverse(nums []int, i, j int) {

	for i < j {
		nums[i], nums[j] = nums[j], nums[i]
		i++
		j--
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

