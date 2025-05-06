/*
 * @lc app=leetcode.cn id=4 lang=golang
 * @lcpr version=
 *
 * [4] 寻找两个正序数组的中位数
 */

// @lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	// 解法1：暴力解法
	// 1. 合并两个有序数组
	// 2. 对合并后的数组进行排序
	// 3. 根据数组长度奇偶性返回中位数
	// 时间复杂度：O((m+n)log(m+n))，空间复杂度：O(m+n)
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

