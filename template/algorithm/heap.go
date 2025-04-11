package algorithm

import (
	"container/heap"
	"sort"
)

type pair struct{ sum, i int }
type hp []pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	return h[i].sum < h[j].sum || (h[i].sum == h[i].sum && h[j].i < h[j].i)
}
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

// 大根堆
type hp2 struct{ sort.IntSlice } // 继承 sort.IntSlice 的 Len Less Swap，这样就只需要实现 Push 和 Pop

func (h hp2) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 加上这行变成最大堆
func (h *hp2) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp2) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *hp2) push(v int)        { heap.Push(h, v) }
func (h *hp2) pop() int          { return heap.Pop(h).(int) } // 稍微封装一下，方便使用

// EXTRA: 参考 Python，引入下面两个效率更高的方法（相比调用 push + pop）
// replace 弹出并返回堆顶，同时将 v 入堆
// 需保证 h 非空
func (h *hp2) replace(v int) int {
	top := h.IntSlice[0]
	h.IntSlice[0] = v
	heap.Fix(h, 0)
	return top
}

// pushPop 先将 v 入堆，然后弹出并返回堆顶
// 使用见下面的 dynamicMedians
func (h *hp2) pushPop(v int) int {
	if h.Len() > 0 && v > h.IntSlice[0] { // 最大堆改成 v < h.IntSlice[0]
		v, h.IntSlice[0] = h.IntSlice[0], v
		heap.Fix(h, 0)
	}
	return v
}

// 自定义类型（int32 可以替换成其余类型）
type hp32 []int32

func (h hp32) Len() int           { return len(h) }
func (h hp32) Less(i, j int) bool { return h[i] < h[j] } // > 为最大堆
func (h hp32) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *hp32) Push(v any)        { *h = append(*h, v.(int32)) }
func (h *hp32) Pop() any          { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *hp32) push(v int32)      { heap.Push(h, v) }
func (h *hp32) pop() int32        { return heap.Pop(h).(int32) } // 稍微封装一下，方便使用

// 支持修改、删除指定元素的堆
// 用法：调用 push 会返回一个 *viPair 指针，记作 p
// 将 p 存于他处（如 slice 或 map），可直接在外部修改 p.v 后调用 fix(p.hi)，从而做到修改堆中指定元素
// 调用 remove(p.hi) 可以从堆中删除 p
// 例题 https://atcoder.jp/contests/abc170/tasks/abc170_e
// 模拟 multiset https://codeforces.com/problemset/problem/1106/E
type viPair struct {
	v  int
	hi int // *viPair 在 mh 中的下标，可随着 Push Pop 等操作自动改变
}
type mh []*viPair // mh 指 modifiable heap

func (h mh) Len() int              { return len(h) }
func (h mh) Less(i, j int) bool    { return h[i].v < h[j].v } // > 为最大堆
func (h mh) Swap(i, j int)         { h[i], h[j] = h[j], h[i]; h[i].hi = i; h[j].hi = j }
func (h *mh) Push(v any)           { *h = append(*h, v.(*viPair)) }
func (h *mh) Pop() any             { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }
func (h *mh) push(v int) *viPair   { p := &viPair{v, len(*h)}; heap.Push(h, p); return p }
func (h *mh) pop() *viPair         { return heap.Pop(h).(*viPair) }
func (h *mh) fix(i int)            { heap.Fix(h, i) }
func (h *mh) remove(i int) *viPair { return heap.Remove(h, i).(*viPair) }

// 懒删除堆
// LC716 https://leetcode.cn/problems/max-stack/
// LC3092 https://leetcode.cn/problems/most-frequent-ids/
// https://codeforces.com/problemset/problem/1883/D 1500
// https://codeforces.com/problemset/problem/796/C 1900
// https://codeforces.com/problemset/problem/1732/D2 2400 简化版懒删除堆
type lazyHeap struct {
	sort.IntSlice
	todo map[int]int
	size int // 实际大小
	sum  int // 实际元素和（可选）
}

func (h lazyHeap) Less(i, j int) bool { return h.IntSlice[i] > h.IntSlice[j] } // 最大堆
func (h *lazyHeap) Push(v any)        { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *lazyHeap) Pop() any          { a := h.IntSlice; v := a[len(a)-1]; h.IntSlice = a[:len(a)-1]; return v }
func (h *lazyHeap) del(v int)         { h.todo[v]++; h.size--; h.sum -= v } // 懒删除
func (h *lazyHeap) push(v int) {
	if h.todo[v] > 0 {
		h.todo[v]--
	} else {
		heap.Push(h, v)
	}
	h.size++
	h.sum += v
}
func (h *lazyHeap) _do() {
	for h.Len() > 0 && h.todo[h.IntSlice[0]] > 0 {
		h.todo[h.IntSlice[0]]--
		heap.Pop(h)
	}
}
func (h *lazyHeap) pop() int    { h._do(); h.size--; v := heap.Pop(h).(int); h.sum -= v; return v }
func (h *lazyHeap) top() int    { h._do(); return h.IntSlice[0] }
func (h *lazyHeap) empty() bool { return h.size == 0 }
func (h *lazyHeap) pushPop(v int) int {
	if h.size > 0 && v < h.top() { // 最大堆，v 比堆顶小就替换堆顶
		h.sum += v - h.IntSlice[0]
		v, h.IntSlice[0] = h.IntSlice[0], v
		heap.Fix(h, 0)
	}
	return v
}

// 对顶堆：滑动窗口前 k 小元素和
// 保证 1 <= k <= windowSize <= n
// 返回数组 kthSum，其中 kthSum[i] 为 a[i:i+windowSize] 的前 k 小元素和
// - [3013. 将数组分成最小总代价的子数组 II](https://leetcode.cn/problems/divide-an-array-into-subarrays-with-minimum-cost-ii/) 2540
// - https://leetcode.cn/problems/find-x-sum-of-all-k-long-subarrays-ii/
// 另见 treap_kthsum.go
func slidingWindowKthSum(a []int, windowSize, k int) []int {
	h := newKthHeap()
	// 注：也可以 copy 一份 a[:k] 然后堆化
	for _, v := range a[:k] {
		h.l.push(v)
	}
	for _, v := range a[k:windowSize] {
		h.add(v)
	}
	kthSum := make([]int, len(a)-windowSize+1)
	kthSum[0] = h.l.sum
	for r := windowSize; r < len(a); r++ {
		l := r - windowSize // 前一个窗口的左端点
		h.add(a[r])
		h.del(a[l]) // 先加再删（注意 windowSize=1 的情况）
		kthSum[l+1] = h.l.sum
	}
	return kthSum
}

type kthHeap struct {
	l *lazyHeap // 最大堆
	r *lazyHeap // 最小堆，所有元素取反
}

func newKthHeap() *kthHeap {
	return &kthHeap{&lazyHeap{todo: map[int]int{}}, &lazyHeap{todo: map[int]int{}}}
}

func (h *kthHeap) empty() bool {
	return h.l.size == 0 && h.r.size == 0
}

func (h *kthHeap) size() int {
	return h.l.size + h.r.size
}

func (h *kthHeap) l2r() {
	if h.l.size == 0 {
		panic("h.l is empty")
	}
	h.r.push(-h.l.pop())
}

func (h *kthHeap) r2l() {
	if h.r.size == 0 {
		panic("h.r is empty")
	}
	h.l.push(-h.r.pop())
}

// 保证 h.l 大小不变
func (h *kthHeap) add(v int) {
	h.r.push(-h.l.pushPop(v))
}

// 保证 h.l 大小不变
func (h *kthHeap) del(v int) {
	if v <= h.l.top() {
		h.l.del(v)
		h.r2l()
	} else {
		h.r.del(-v)
	}
}

// 把 h.l 的大小调整为 k
func (h *kthHeap) balance(k int) {
	for h.l.size > k {
		h.l2r()
	}
	for h.l.size < k {
		h.r2l()
	}
}
