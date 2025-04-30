/*
 * @lc app=leetcode.cn id=54 lang=golang
 * @lcpr version=
 *
 * [54] 螺旋矩阵
 */
package main

// @lc code=start
func spiralOrder(matrix [][]int) []int {

	// sl: 类似状态机，记录方向和边界
	directions := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	row, col := len(matrix), len(matrix[0])

	visted := make([][]bool, row)
	for i := 0; i < row; i++ {
		visted[i] = make([]bool, col)
	}

	result := make([]int, row*col)
	x, y, dir := 0, 0, 0

	for i := 0; i < row*col; i++ {

		result[i] = matrix[x][y]
		visted[x][y] = true

		direction := directions[dir]
		// 结算下个位置
		nextx, nexty := x+direction[0], y+direction[1]

		if nextx < 0 || nextx >= row || nexty < 0 || nexty >= col || visted[nextx][nexty] {
			dir = (dir + 1) % 4
			// 超出边界，或者访问过。节切换
			nextx, nexty = x+directions[dir][0], y+directions[dir][1]
		}

		x, y = nextx, nexty
	}
	return result
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[1,2,3,4],[5,6,7,8],[9,10,11,12]]\n
// @lcpr case=end

*/
