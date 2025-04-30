/*
 * @lc app=leetcode.cn id=238 lang=golang
 * @lcpr version=
 *
 * [238] 除自身以外数组的乘积
 */
package main

// @lc code=start
func productExceptSelf(nums []int) []int {
	//*思路： 数可以等于两边数的乘积。左右两次遍历，计算乘积
	n := len(nums)
	lefts, rights := []int{nums[0]}, []int{nums[n-1]}
	for i := 1; i < n; i++ {
		lefts = append(lefts, nums[i]*lefts[i-1])
		rights = append(rights, nums[n-1-i]*rights[i-1])
	}

	res := []int{}
	for i := 0; i < n; i++ {
		// 左边的数
		left, right := 1, 1
		if i-1 >= 0 {
			left = lefts[i-1]
		}
		// 右边的数
		if n-i-2 >= 0 {
			right = rights[n-i-2]
		}
		res = append(res, left*right)
	}
	return res
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3,4]\n
// @lcpr case=end

// @lcpr case=start
// [-1,1,0,-3,3]\n
// @lcpr case=end

*/
