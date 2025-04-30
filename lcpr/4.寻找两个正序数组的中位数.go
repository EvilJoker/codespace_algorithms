/*
 * @lc app=leetcode.cn id=4 lang=golang
 * @lcpr version=
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 暴力： 合并两个数组，找出中间两个数
	merged := append(nums1, nums2...)
	sort.Ints(merged)

	l := len(merged)
	if l%2 == 1 {
		return float64(merged[l/2])
	} else {
		return float64((merged[l/2] + merged[(l/2)-1])) / 2
	}

}

// @lc code=end

/*
// @lcpr case=start
// [1,3]\n[2]\n
// @lcpr case=end

// @lcpr case=start
// [1,2]\n[3,4]\n
// @lcpr case=end

*/

