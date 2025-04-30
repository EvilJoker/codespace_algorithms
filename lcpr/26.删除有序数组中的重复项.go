/*
 * @lc app=leetcode.cn id=26 lang=golang
 * @lcpr version=
 *
 * [26] 删除有序数组中的重复项
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	/* 思路：比较前一位元素*/
	if len(nums) == 0 {
		return 0
	}

	p := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[p-1] {
			nums[p] = nums[i]
			p++
		}
	}

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

