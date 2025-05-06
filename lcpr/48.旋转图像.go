/*
 * @lc app=leetcode.cn id=48 lang=golang
 * @lcpr version=
 *
 * [48] 旋转图像
 */
package main

// @lc code=start
func rotate(matrix [][]int) {
	// 解题思路：
	// 1. 顺时针旋转90度可以通过两次翻转实现：
	//    - 先沿主对角线翻转（左上到右下）
	//    - 再沿垂直中线左右翻转
	// 2. 时间复杂度：O(n²)，空间复杂度：O(1)
	// 3. 注意：对角线翻转时只需要遍历上三角或下三角区域

	n := len(matrix)
	if n <= 1 {
		return
	}

	// 第一步：沿主对角线翻转
	// 只需要遍历下三角区域（包括对角线）
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			// 交换对称位置的元素
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}

	// 第二步：沿垂直中线左右翻转
	// 对每一行进行翻转
	for i := 0; i < n; i++ {
		// 只需要遍历到中间位置
		for j := 0; j < n/2; j++ {
			// 交换对称位置的元素
			matrix[i][j], matrix[i][n-j-1] = matrix[i][n-j-1], matrix[i][j]
		}
	}
}

// @lc code=end

/*
// @lcpr case=start
// [[1,2,3],[4,5,6],[7,8,9]]\n
// @lcpr case=end

// @lcpr case=start
// [[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]\n
// @lcpr case=end

*/
