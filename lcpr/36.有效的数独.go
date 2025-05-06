/*
 * @lc app=leetcode.cn id=36 lang=golang
 * @lcpr version=
 *
 * [36] 有效的数独
 */
package main

import "strconv"

// @lc code=start
/*
解题思路：
1. 数独规则：
   - 每行数字1-9不能重复
   - 每列数字1-9不能重复
   - 每个3x3宫格内数字1-9不能重复
   - '.'表示空格

2. 验证方法：
   - 分别验证行、列、3x3宫格
   - 使用数组记录已出现的数字
   - 遇到'.'跳过
   - 数字范围必须在1-9之间
   - 重复数字返回false
*/

func isValidSudoku(board [][]byte) bool {
	// 验证行和列
	for i := 0; i < len(board); i++ {
		if !(rowVerfy(board, i) && colVerfy(board, i)) {
			return false
		}
	}

	// 验证3x3宫格
	for i := 0; i < len(board); i += 3 {
		for j := 0; j < len(board); j += 3 {
			if !squreVerfy(board, i, j) {
				return false
			}
		}
	}
	return true
}

// 验证行
func rowVerfy(board [][]byte, row int) bool {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < 9; i++ {
		c := string(board[row][i])
		if c == "." {
			continue
		}
		t, _ := strconv.Atoi(c)
		if t < 1 || t > 9 {
			return false
		}
		if nums[t-1] == -1 {
			return false
		}
		nums[t-1] = -1 // 标记数字已使用
	}
	return true
}

// 验证列
func colVerfy(board [][]byte, col int) bool {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := 0; i < 9; i++ {
		c := string(board[i][col])
		if c == "." {
			continue
		}
		t, _ := strconv.Atoi(c)
		if t < 1 || t > 9 {
			return false
		}
		if nums[t-1] == -1 {
			return false
		}
		nums[t-1] = -1 // 标记数字已使用
	}
	return true
}

// 验证3x3宫格
func squreVerfy(board [][]byte, row, col int) bool {
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}

	for i := row; i < row+3; i++ {
		for j := col; j < col+3; j++ {
			c := string(board[i][j])
			if c == "." {
				continue
			}
			t, _ := strconv.Atoi(c)
			if t < 1 || t > 9 {
				return false
			}
			if nums[t-1] == -1 {
				return false
			}
			nums[t-1] = -1 // 标记数字已使用
		}
	}
	return true
}

// @lc code=end

/*
// @lcpr case=start
// \n[["5","3",".",".","7",".",".",".","."]\n,["6",".",".","1","9","5",".",".","."]\n,[".","9","8",".",".",".",".","6","."]\n,["8",".",".",".","6",".",".",".","3"]\n,["4",".",".","8",".","3",".",".","1"]\n,["7",".",".",".","2",".",".",".","6"]\n,[".","6",".",".",".",".","2","8","."]\n,[".",".",".","4","1","9",".",".","5"]\n,[".",".",".",".","8",".",".","7","9"]]\n
// @lcpr case=end

// @lcpr case=start
// \n[["8","3",".",".","7",".",".",".","."]\n,["6",".",".","1","9","5",".",".","."]\n,[".","9","8",".",".",".",".","6","."]\n,["8",".",".",".","6",".",".",".","3"]\n,["4",".",".","8",".","3",".",".","1"]\n,["7",".",".",".","2",".",".",".","6"]\n,[".","6",".",".",".",".","2","8","."]\n,[".",".",".","4","1","9",".",".","5"]\n,[".",".",".",".","8",".",".","7","9"]]\n
// @lcpr case=end

*/
