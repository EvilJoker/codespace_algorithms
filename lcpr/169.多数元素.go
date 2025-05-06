/*
 * @lc app=leetcode.cn id=169 lang=golang
 * @lcpr version=
 *
 * [169] 多数元素
 */

// @lc code=start
/*
思路：Boyer-Moore 投票算法
1. 将第一个元素设为候选众数，计数器设为1
2. 遍历数组：
   - 如果当前元素等于候选数，计数器+1
   - 如果当前元素不等于候选数，计数器-1
   - 如果计数器为0，更换候选数为当前元素，计数器重置为1
3. 最后剩下的候选数就是众数
时间复杂度：O(n)
空间复杂度：O(1)
*/
func majorityElement(nums []int) int {
	// 初始化候选众数和计数器
	candidate := nums[0]
	count := 1

	// 遍历数组
	for i := 1; i < len(nums); i++ {
		if nums[i] == candidate {
			count++
		} else {
			count--
			// 当计数器为0时，更换候选数
			if count == 0 {
				candidate = nums[i]
				count = 1
			}
		}
	}

	return candidate
}

// @lc code=end

/*
// @lcpr case=start
// [3,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [2,2,1,1,1,2,2]\n
// @lcpr case=end

*/

