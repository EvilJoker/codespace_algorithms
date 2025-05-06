/*
 * @lc app=leetcode.cn id=17 lang=golang
 * @lcpr version=
 *
 * [17] 电话号码的字母组合
 */
package main

// @lc code=start

var phoneMap = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}

func letterCombinations(digits string) []string {
	// 解题思路：
	// 1. 使用哈希表存储数字对应的字母组合，便于快速查找
	// 2. 使用回溯法生成所有可能的组合，通过递归实现
	// 3. 对于每个数字，遍历其对应的所有字母，并与之前的结果组合
	// 4. 时间复杂度：O(4^n)，其中n为输入数字的长度，最坏情况下每个数字对应4个字母
	// 5. 空间复杂度：O(n)，递归调用栈的深度为n

	// 处理空输入的情况
	if len(digits) == 0 {
		return []string{}
	}
	// 初始化结果数组
	combinas := []string{}
	// 开始回溯，从第一个数字开始，初始组合为空字符串
	var pushback func(index int, combina string)
	pushback = func(index int, combina string) {
		// 当处理完所有数字时，将当前组合加入结果数组
		if index == len(digits) {
			combinas = append(combinas, combina)
		}
		// 获取当前数字对应的字母集合
		letters, _ := phoneMap[string(digits[index])]

		// 遍历当前数字对应的所有字母
		for i := 0; i < len(letters); i++ {
			// 递归处理下一个数字，将当前字母加入组合
			pushback(index+1, combina+string(letters[i]))
		}
	}

	pushback(0, "")

	return combinas
}

// @lc code=end

/*
// @lcpr case=start
// "23"\n
// @lcpr case=end

// @lcpr case=start
// ""\n
// @lcpr case=end

// @lcpr case=start
// "2"\n
// @lcpr case=end

*/
