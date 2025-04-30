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
	// 思路: 对数组排序，使用双指针；去重让每个数尽量滑动到不重复位置
	res := [][]int{}

	sort.Ints(nums)
	n := len(nums)

	for a := 0; a < n; a++ { // `a` 只需要遍历到倒数第三个元素

		// 去重：跳过重复的 `a`
		if a > 0 && nums[a] == nums[a-1] {
			continue
		}

		b, c := a+1, n-1

		for b < c {
			sum := nums[a] + nums[b] + nums[c]

			if sum == 0 {
				res = append(res, []int{nums[a], nums[b], nums[c]})
				b++
				c--

				// 去重：跳过重复的 `b`
				for b < c && nums[b] == nums[b-1] {
					b++
				}
				// 去重：跳过重复的 `c`
				for b < c && nums[c] == nums[c+1] {
					c--
				}

			} else if sum < 0 {
				b++
			} else {
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
