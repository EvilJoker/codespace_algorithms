/*
 * @lc app=leetcode.cn id=27 lang=golang
 * @lcpr version=
 *
 * [27] 移除元素
 */

// @lc code=start
func removeElement(nums []int, val int) int {
	/* 思路：指针将不相等的 元素，保存下来 */
	p := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[p] = nums[i]
			p++
		}
	}

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

