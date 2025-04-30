/*
 * @lc app=leetcode.cn id=200 lang=golang
 * @lcpr version=
 *
 * [200] 岛屿数量
 */
package main

// @lc code=start
func numIslands(grid [][]byte) int {
	//sl: bfs 遍历，每个节点都尽可能染色。遇到几个不染色的就是成功的
	count, row, col := 0, len(grid), len(grid[0])

	var render func(i, j int)

	render = func(i, j int) {
		directions := [][]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

		queue := [][]int{{i, j}} // 待访问的节点

		for len(queue) > 0 {
			head := queue[0]
			queue = queue[1:]
			x, y := head[0], head[1]
			if grid[x][y] != '1' {
				continue
			}
			grid[x][y] = 'x' // 染色
			for _, v := range directions {
				nx, ny := x+v[0], y+v[1]
				if nx >= 0 && nx < row && ny >= 0 && ny < col {
					queue = append(queue, []int{nx, ny})
				}

			}

		}

	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if grid[i][j] == '1' {
				count++
				render(i, j)
			}
		}
	}
	return count
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
