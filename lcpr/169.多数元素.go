/*
 * @lc app=leetcode.cn id=169 lang=golang
 * @lcpr version=
 *
 * [169] 多数元素
 */

// @lc code=start
func majorityElement(nums []int) int {
	/* 思路： 将每个元素设置为候选人，如果遇到相同就 count ++,不同的 就 count --,如果 =0就切换候选人
	众数候选人，总会剩下来
	*/
	candida := nums[0]
	count := 1

	for i := 1; i < len(nums); i++ {
		if nums[i] == candida {
			count++
			continue
		}
		if nums[i] != candida {
			count--
		}
		if count == 0 {
			candida = nums[i]
			count++
		}
	}
	return candida
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

