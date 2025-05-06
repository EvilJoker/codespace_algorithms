package main

/*
 * @lc app=leetcode.cn id=74 lang=golang
 * @lcpr version=
 *
 * [74] 搜索二维矩阵
 */

// @lc code=start
func searchMatrix(matrix [][]int, target int) bool {

	/*
		解题思路：二分查找
		1. 将二维矩阵映射成一维数组：
		   - 矩阵中的每个元素matrix[i][j]可以映射到一维数组的index = i*col + j
		   - 一维数组的index可以映射回矩阵的i = index/col, j = index%col
		2. 在一维数组上进行二分查找：
		   - 左边界left = 0
		   - 右边界right = row*col-1
		   - 中间位置mid = (left+right)/2
		3. 时间复杂度：O(log(m*n))，空间复杂度：O(1)
	*/

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
