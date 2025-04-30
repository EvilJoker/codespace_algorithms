/*
 * @lc app=leetcode.cn id=46 lang=golang
 * @lcpr version=
 *
 * [46] 全排列
 */
package main

// @lc code=start
func permute(nums []int) [][]int {
	//sl ： 排列是有顺序的，

	result := [][]int{}

	n := len(nums)

	var dfs func(path []int, used map[int]bool)

	dfs = func(path []int, used map[int]bool) {
		// 到结尾
		if len(path) == n {
			tmp := make([]int, n)
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		for _, v := range nums {
			if _, ok := used[v]; ok {
				continue
			}
			used[v] = true
			dfs(append(path, v), used)
			// 回溯，撤销选择. 或者copy
			delete(used, v)
		}
	}

	dfs([]int{}, map[int]bool{})
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [1,2,3]\n
// @lcpr case=end

// @lcpr case=start
// [0,1]\n
// @lcpr case=end

// @lcpr case=start
// [1]\n
// @lcpr case=end

*/
