package main

/*
 * @lc app=leetcode.cn id=74 lang=golang
 * @lcpr version=
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {

	//sl: 将二维映射成一维

	row, col := len(matrix), len(matrix[0])
	left, right := 0, row*col-1

	mid := (left-right)/2 + right
	for left <= right {
		i, j := indexToIj(col, mid)
		if matrix[i][j] == target {
			return true
		}

		if target > matrix[i][j] {
			left = mid + 1
		} else {
			right = mid - 1
		}
		mid = (left-right)/2 + right
	}
	return false

}

func ijToIndex(col, i, j int) int {
	return i*col + j
}

func indexToIj(col, index int) (int, int) {
	return index / col, index % col
}

// @lc code=end

/*
// @lcpr case=start
// [[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n3\n
// @lcpr case=end

// @lcpr case=start
// [[1,3,5,7],[10,11,16,20],[23,30,34,60]]\n13\n
// @lcpr case=end

*/
