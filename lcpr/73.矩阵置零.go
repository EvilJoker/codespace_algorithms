/*
 * @lc app=leetcode.cn id=73 lang=golang
 * @lcpr version=
 *
 * [73] 矩阵置零
 */
package main

// @lc code=start
func setZeroes(matrix [][]int) {
	/*
		解题思路：
		1. 使用矩阵的第一行和第一列作为标记数组
		2. 使用两个额外的标志位setX和setY来标记第一行和第一列是否需要置零
		3. 遍历矩阵，如果matrix[i][j]为0：
		   - 将matrix[i][0]和matrix[0][j]置为0作为标记
		4. 再次遍历矩阵，根据标记将对应行列置零
		5. 最后根据setX和setY处理第一行和第一列
		时间复杂度：O(m*n)，空间复杂度：O(1)
	*/

	setX, setY := false, false

	row, col := len(matrix), len(matrix[0])

	for i := 0; i < row; i++ {
		if matrix[i][0] == 0 {
			setY = true
			break
		}
	}

	for j := 0; j < col; j++ {
		if matrix[0][j] == 0 {
			setX = true
			break
		}
	}

	// 遍历
	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}
	// 渲染

	for i := 1; i < row; i++ {
		for j := 1; j < col; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if setX == true {
		for j := 0; j < col; j++ {
			matrix[0][j] = 0
		}
	}
	if setY == true {
		for i := 0; i < row; i++ {
			matrix[i][0] = 0
		}
	}

}

// @lc code=end

/*
// @lcpr case=start
// [[1,1,1],[1,0,1],[1,1,1]]\n
// @lcpr case=end

// @lcpr case=start
// [[0,1,2,0],[3,4,5,2],[1,3,1,5]]\n
// @lcpr case=end

*/
