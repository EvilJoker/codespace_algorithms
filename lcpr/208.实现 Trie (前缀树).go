/*
 * @lc app=leetcode.cn id=208 lang=golang
 * @lcpr version=
 *
 * [208] 实现 Trie (前缀树)
 */
package main

// @lc code=start
type Trie struct {
	end  bool
	dict [26]*Trie // 彩虹表
}

func Constructor() Trie {
	return Trie{}

}

func (this *Trie) Insert(word string) {
	// 递归的节点
	cur := this

	for _, c := range word {
		idx := c - 'a'
		if cur.dict[idx] == nil {
			cur.dict[idx] = &Trie{}
		}
		cur = cur.dict[idx]
	}

	cur.end = true

}

func (this *Trie) Search(word string) bool {

	cur := this
	for _, c := range word {
		idx := c - 'a'
		if cur.dict[idx] == nil {
			return false
		}
		cur = cur.dict[idx]
	}
	return cur.end

}

func (this *Trie) StartsWith(prefix string) bool {

	cur := this
	for _, c := range prefix {
		idx := c - 'a'
		if cur.dict[idx] == nil {
			return false
		}
		cur = cur.dict[idx]
	}
	return true
}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);
 */
// @lc code=end
