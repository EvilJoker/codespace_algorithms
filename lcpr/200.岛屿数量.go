/*
 * @lc app=leetcode.cn id=200 lang=golang
 * @lcpr version=
 *
 * [200] 岛屿数量
 */
package main

// @lc code=start
/*
思路：
1. 使用BFS遍历每个岛屿
2. 对于每个未访问的陆地('1')，将其标记为已访问('x')
3. 从该点开始BFS，将所有相连的陆地都标记为已访问
4. 统计未访问的陆地数量即为岛屿数量
*/
func numIslands(grid [][]byte) int {
	// 初始化变量
	islandCount := 0
	rows, cols := len(grid), len(grid[0])

	// 定义四个方向：上、下、左、右
	directions := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	// BFS函数：从给定位置开始，标记所有相连的陆地
	var bfs func(startRow, startCol int)
	bfs = func(startRow, startCol int) {
		queue := [][]int{{startRow, startCol}}

		for len(queue) > 0 {
			// 取出队首元素
			current := queue[0]
			queue = queue[1:]
			row, col := current[0], current[1]

			// 如果当前位置不是陆地，跳过
			if grid[row][col] != '1' {
				continue
			}

			// 标记当前位置为已访问
			grid[row][col] = 'x'

			// 检查四个方向的相邻位置
			for _, dir := range directions {
				newRow, newCol := row+dir[0], col+dir[1]
				// 确保新位置在网格范围内
				if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
					queue = append(queue, []int{newRow, newCol})
				}
			}
		}
	}

	// 遍历整个网格
	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			// 发现未访问的陆地，增加岛屿计数并开始BFS
			if grid[i][j] == '1' {
				islandCount++
				bfs(i, j)
			}
		}
	}

	return islandCount
}

// @lc code=end

/*
// @lcpr case=start
// [\n["1","1","1","1","0"],\n["1","1","0","1","0"],\n["1","1","0","0","0"],\n["0","0","0","0","0"]\n]\n
// @lcpr case=end

// @lcpr case=start
// [\n["1","1","0","0","0"],\n["1","1","0","0","0"],\n["0","0","1","0","0"],\n["0","0","0","1","1"]\n]\n
// @lcpr case=end

*/
