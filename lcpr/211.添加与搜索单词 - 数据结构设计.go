/*
 * @lc app=leetcode.cn id=211 lang=golang
 * @lcpr version=
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */
package main

// @lc code=start
type WordDictionary struct {
	val  bool
	dict [26]*WordDictionary
}

func Constructor() WordDictionary {
	return WordDictionary{}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this
	for _, c := range word {
		idx := c - 'a'
		if cur.dict[idx] == nil {
			cur.dict[idx] = &WordDictionary{}
		}
		cur = cur.dict[idx]
	}

	cur.val = true

}

func (this *WordDictionary) Search(word string) bool {
	if this == nil {
		return false
	}
	if len(word) == 0 {
		return this.val
	}

	c := word[0]

	if c != '.' {
		idx := c - 'a'
		if this.dict[idx] == nil {
			return false
		}
		return this.dict[idx].Search(word[1:])
	} else {
		for _, child := range this.dict {
			// 命中
			if child != nil && child.Search(word[1:]) {
				return true
			}
		}
	}
	return false
}

/**
 * Your WordDictionary object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddWord(word);
 * param_2 := obj.Search(word);
 */
// @lc code=end

/*
// @lcpr case=start
// ["WordDictionary","addWord","addWord","addWord","search","search","search","search"]\n[[],["bad"],["dad"],["mad"],["pad"],["bad"],[".ad"],["b.."]]\n
// @lcpr case=end

*/
