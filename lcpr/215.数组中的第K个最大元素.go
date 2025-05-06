/*
 * @lc app=leetcode.cn id=215 lang=golang
 * @lcpr version=30104
 *
 * [215] 数组中的第K个最大元素
 */

// @lc code=start
func findKthLargest(nums []int, k int) int {
	/*
		解题思路：
		1. 使用小顶堆维护k个最大元素
		2. 遍历数组，将元素加入堆中
		3. 当堆的大小超过k时，弹出堆顶元素（最小值）
		4. 最终堆顶元素即为第k个最大元素
	*/

	// 使用container/heap包中的堆
	h := &IntHeap{}
	heap.Init(h)

	// 遍历数组
	for _, num := range nums {
		heap.Push(h, num)
		// 当堆的大小超过k时，弹出堆顶元素
		if h.Len() > k {
			heap.Pop(h)
		}
	}

	// 返回堆顶元素，即第k个最大元素
	return (*h)[0]
}

// 实现container/heap.Interface接口
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

// @lc code=end

/*
// @lcpr case=start
// 2\n
// @lcpr case=end

// @lcpr case=start
// 4\n
// @lcpr case=end

*/

