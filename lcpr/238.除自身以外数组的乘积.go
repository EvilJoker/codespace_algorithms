/*
 * @lc app=leetcode.cn id=238 lang=golang
 * @lcpr version=
 *
 * [238] 除自身以外数组的乘积
 */
package main

// @lc code=start
/*
思路：
1. 对于数组中的每个位置i，我们需要计算除了nums[i]之外所有数的乘积
2. 可以将问题分解为：每个位置的答案 = 左边所有数的乘积 * 右边所有数的乘积
3. 使用两次遍历：
   - 第一次从左到右计算每个位置左边所有数的乘积
   - 第二次从右到左计算每个位置右边所有数的乘积
4. 最后将左右乘积相乘得到最终结果
*/
func productExceptSelf(nums []int) []int {
	n := len(nums)
	// leftProducts[i] 表示位置i左边所有数的乘积
	leftProducts := make([]int, n)
	leftProducts[0] = 1 // 第一个数左边没有数，乘积为1

	// 计算左边乘积
	for i := 1; i < n; i++ {
		leftProducts[i] = leftProducts[i-1] * nums[i-1]
	}

	// rightProduct 表示当前位置右边所有数的乘积
	rightProduct := 1
	// 从右向左遍历，同时计算最终结果
	for i := n - 1; i >= 0; i-- {
		leftProducts[i] = leftProducts[i] * rightProduct
		rightProduct *= nums[i]
	}

	return leftProducts
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
