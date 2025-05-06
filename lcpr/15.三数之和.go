/*
 * @lc app=leetcode.cn id=15 lang=golang
 * @lcpr version=
 *
 * [15] 三数之和
 */
package main

import "sort"

// @lc code=start
func threeSum(nums []int) [][]int {
	// 解题思路：
	// 1. 对数组排序，使用双指针；去重让每个数尽量滑动到不重复位置
	// 2. 时间复杂度：O(n^2)，空间复杂度：O(1)

	// 初始化结果数组
	res := [][]int{}

	// 对数组进行排序，便于后续双指针操作
	sort.Ints(nums)
	n := len(nums)

	// 遍历第一个数 a
	for a := 0; a < n-2; a++ { // 优化：a 只需要遍历到倒数第三个元素
		// 去重：如果当前数字与前一个数字相同，跳过
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}

		// 初始化双指针：b 从 a+1 开始，c 从末尾开始
		b, c := a+1, n-1

		// 双指针向中间移动
		for b < c {
			sum := nums[a] + nums[b] + nums[c]

			if sum == 0 {
				// 找到符合条件的组合，加入结果集
				res = append(res, []int{nums[a], nums[b], nums[c]})
				
				// 移动双指针
				b++
				c--

				// 去重：跳过重复的 b
				for b < c && nums[b] == nums[b-1] {
					b++
				}
				// 去重：跳过重复的 c
				for b < c && nums[c] == nums[c+1] {
					c--
				}
			} else if sum < 0 {
				// 和小于0，需要增大，移动左指针
				b++
			} else {
				// 和大于0，需要减小，移动右指针
				c--
			}
		}
	}

	return res

}

// @lc code=end

/*
// @lcpr case=start
// [-1,0,1,2,-1,-4]\n
// @lcpr case=end

// @lcpr case=start
// [0,1,1]\n
// @lcpr case=end

// @lcpr case=start
// [0,0,0]\n
// @lcpr case=end

*/
