/*
 * @lc app=leetcode.cn id=46 lang=golang
 * @lcpr version=
 *
 * [46] 全排列
 */
package main

// @lc code=start
func permute(nums []int) [][]int {
	// 解题思路：
	// 1. 使用回溯算法，通过DFS遍历所有可能的排列
	// 2. 使用used map记录已使用的数字，避免重复使用
	// 3. 当path长度等于nums长度时，表示找到一个完整排列
	// 4. 每次选择未使用的数字加入path，递归后回溯
	// 时间复杂度：O(n!)，空间复杂度：O(n)

	result := [][]int{}
	n := len(nums)

	// 定义DFS函数，path记录当前路径，used记录已使用的数字
	var dfs func(path []int, used map[int]bool)
	dfs = func(path []int, used map[int]bool) {
		// 找到一个完整排列
		if len(path) == n {
			// 创建path的副本并加入结果集
			tmp := make([]int, n)
			copy(tmp, path)
			result = append(result, tmp)
			return
		}

		// 遍历所有数字
		for _, num := range nums {
			// 跳过已使用的数字
			if _, exists := used[num]; exists {
				continue
			}

			// 选择当前数字
			used[num] = true
			// 递归搜索
			dfs(append(path, num), used)
			// 回溯，撤销选择
			delete(used, num)
		}
	}

	// 从空路径开始搜索
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
