/*
 * @lc app=leetcode.cn id=88 lang=golang
 * @lcpr version=
 *
 * [88] 合并两个有序数组
 */

// @lc code=start
func merge(nums1 []int, m int, nums2 []int, n int) {
	/*
		思路: 倒叙开始，最大的往 nums1 (m,n )的屁股放
	*/
	p1, p2 := m-1, n-1

	for i := 0; i < m+n; i++ {
		if p1 >= 0 && p2 >= 0 {
			if nums1[p1] > nums2[p2] {
				nums1[m+n-i-1] = nums1[p1]
				p1--
			} else {
				nums1[m+n-i-1] = nums2[p2]
				p2--
			}
		} else if p1 < 0 && p2 >= 0 {
			nums1[m+n-i-1] = nums2[p2]
			p2--
		} else if p1 >= 0 && p2 < 0 {
			nums1[m+n-i-1] = nums1[p1]
			p1--
		}
	}

}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,0,0,0]\n3\n[2,5,6]\n3\n
// @lcpr case=end

// @lcpr case=start
// [1]\n1\n[]\n0\n
// @lcpr case=end

// @lcpr case=start
// [0]\n0\n[1]\n1\n
// @lcpr case=end

*/

