/*
 * @lc app=leetcode.cn id=54 lang=golang
 * @lcpr version=
 *
 * [54] 螺旋矩阵
 */
package main

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	// 解题思路：
	// 1. 使用模拟法，按照顺时针方向遍历矩阵
	// 2. 定义四个方向：右、下、左、上，使用方向数组表示
	// 3. 使用visited数组记录已访问的位置，避免重复访问
	// 4. 当遇到边界或已访问位置时，顺时针旋转90度改变方向
	// 5. 时间复杂度：O(m*n)，空间复杂度：O(m*n)，其中m和n为矩阵的行数和列数
	// 定义四个方向：右、下、左、上
	// 使用二维数组表示方向向量，每个方向包含 x 和 y 的偏移量
	directions := [][]int{
		{0, 1},  // 向右移动
		{1, 0},  // 向下移动
		{0, -1}, // 向左移动
		{-1, 0}, // 向上移动
	}

	// 获取矩阵的行数和列数
	row, col := len(matrix), len(matrix[0])

	// 创建访问标记数组，用于记录已访问的位置
	visited := make([][]bool, row)
	for i := 0; i < row; i++ {
		visited[i] = make([]bool, col)
	}

	// 初始化结果数组，大小为矩阵元素总数
	result := make([]int, row*col)

	// 初始化当前位置和方向
	x, y := 0, 0 // 当前位置坐标
	dir := 0     // 当前方向索引

	// 遍历所有元素
	for i := 0; i < row*col; i++ {
		// 记录当前元素
		result[i] = matrix[x][y]
		visited[x][y] = true

		// 计算下一个位置
		direction := directions[dir]
		nextX, nextY := x+direction[0], y+direction[1]

		// 检查是否需要改变方向
		// 1. 超出边界
		// 2. 已经访问过
		if nextX < 0 || nextX >= row || nextY < 0 || nextY >= col || visited[nextX][nextY] {
			// 顺时针旋转90度，切换到下一个方向
			dir = (dir + 1) % 4
			// 使用新方向重新计算下一个位置
			nextX, nextY = x+directions[dir][0], y+directions[dir][1]
		}

		// 更新当前位置
		x, y = nextX, nextY
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
