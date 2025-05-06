/*
 * @lc app=leetcode.cn id=211 lang=golang
 * @lcpr version=
 *
 * [211] 添加与搜索单词 - 数据结构设计
 */
package main

// @lc code=start

// 解题思路：
// 1. 使用前缀树(Trie)数据结构来存储单词
// 2. 每个节点包含26个子节点，对应26个字母
// 3. 节点中的val标记是否为单词结尾
// 4. 搜索时支持通配符'.'，需要遍历所有可能的子节点

// WordDictionary 前缀树节点结构
type WordDictionary struct {
	isEnd    bool                // 标记是否为单词结尾
	children [26]*WordDictionary // 子节点数组，对应26个字母
}

// Constructor 初始化前缀树
func Constructor() WordDictionary {
	return WordDictionary{}
}

// AddWord 添加单词到前缀树
func (this *WordDictionary) AddWord(word string) {
	node := this
	for _, ch := range word {
		idx := ch - 'a'
		if node.children[idx] == nil {
			node.children[idx] = &WordDictionary{}
		}
		node = node.children[idx]
	}
	node.isEnd = true
}

// Search 搜索单词
// 支持通配符'.'，可以匹配任意字母
func (this *WordDictionary) Search(word string) bool {
	if this == nil {
		return false
	}
	if len(word) == 0 {
		return this.isEnd
	}

	ch := word[0]
	if ch != '.' {
		// 精确匹配
		idx := ch - 'a'
		if this.children[idx] == nil {
			return false
		}
		return this.children[idx].Search(word[1:])
	} else {
		// 通配符匹配，遍历所有子节点
		for _, child := range this.children {
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
