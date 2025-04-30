/*
 * @lc app=leetcode.cn id=146 lang=golang
 * @lcpr version=
 *
 * [146] LRU 缓存
 */
package main

/*
1. 使用链表实现LRU,先定义数据结构，有成员 head,tail,容量 cap
2. 实现 map 影响，获取 成员时，直接获取而不是遍历，实现o1
3. 访问，更新，都会将元素重新插入 head,相当于最新记录
4. 超限时，删除队尾元素
5. 需要实现 add。和remove ，作为基础函数， get 和 put 是上层
6. get 和 remove 需要考虑边界情况



*/
// @lc code=start
type DNode struct {
	Key  int
	Val  int
	Prev *DNode
	Next *DNode
}

type LRUCache struct {
	capacity   int
	used       int
	head, tail *DNode // 所有的节点都是 head 和tail 中间的节点。
	dict       map[int]*DNode
}

func Constructor(capacity int) LRUCache {

	head := &DNode{Prev: nil, Next: nil}
	tail := &DNode{Prev: nil, Next: nil}
	head.Next = tail
	tail.Prev = head

	return LRUCache{capacity: capacity, used: 0, head: head, tail: tail, dict: map[int]*DNode{}}
}

func (this *LRUCache) add(key int, value int) {
	// 增加一个新元素
	newNode := &DNode{Key: key, Val: value, Prev: nil, Next: nil}

	next := this.head.Next

	this.head.Next = newNode
	newNode.Prev = this.head
	newNode.Next = next
	next.Prev = newNode
	this.used += 1
	this.dict[key] = newNode

	if this.used > this.capacity {
		deleted := this.tail.Prev
		prev := deleted.Prev
		prev.Next = this.tail
		this.tail.Prev = prev
		this.used -= 1
		delete(this.dict, deleted.Key)
	}
}

func (this *LRUCache) remove(key int) {
	if deleted, exist := this.dict[key]; exist {

		prev := deleted.Prev
		next := deleted.Next

		prev.Next = next
		next.Prev = prev
		this.used -= 1
		delete(this.dict, key)
	}

}

func (this *LRUCache) Get(key int) int {

	if node, exist := this.dict[key]; exist {
		this.remove(key)
		this.add(node.Key, node.Val)
		return node.Val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {

	if _, exist := this.dict[key]; exist {
		this.remove(key)

	}

	this.add(key, value)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
// @lc code=end
