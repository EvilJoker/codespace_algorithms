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
var combinas []string

func letterCombinations(digits string) []string {
	// 组合，回溯
	/* 1. hash 存储 kv 表
	/* 2. 回溯框架:result, path
	/* 3. 终止条件 判断是否满足 path 长度
	*/
	if len(digits) == 0 {
		return []string{}
	}
	combinas = []string{}
	pushback(digits, 0, "")

	return combinas
}

func pushback(digits string, index int, combina string) {
	// 加入
	if len(digits) == index {
		// 终止条件加入
		combinas = append(combinas, combina)
		return
	}

	letters, _ := phoneMap[string(digits[index])]

	for i := 0; i < len(letters); i++ {
		pushback(digits, index+1, combina+string(letters[i]))
	}

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
