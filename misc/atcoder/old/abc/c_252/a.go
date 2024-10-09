package c_252

import (
	"bufio"
	"container/heap"
	"fmt"
	"io"
	"sort"
)

// https://atcoder.jp/contests/abc252/tasks/abc252_f
//
//输入 n (2≤n≤2e5) 和 L(≤1e15)，长为 n 的数组 a (1≤a[i]≤1e9, sum(a)≤L)。
//有一根长为 L 的面包，需要分给 n 个小孩，每个小孩需要长度恰好为 a[i] 的面包。
//对于任意一根长为 k 的面包，你可以切成两段，要求每段长度都为整数，切的花费为 k。
//输出最小花费。

func AtCoderABC251E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, l, v int
	fmt.Fscan(r, &n, &l)
	hp := &hp2{}
	for i := int(0); i < n; i++ {
		fmt.Fscan(r, &v)
		heap.Push(hp, v)
		l -= v
	}
	if l > 0 {
		heap.Push(hp, l)
	}

	ans := int(0)
	for hp.Len() > 1 {
		a, b := heap.Pop(hp).(int), heap.Pop(hp).(int)
		ans += a + b
		heap.Push(hp, a+b)
	}
	fmt.Fprintln(w, ans)
}

type hp2 struct{ sort.IntSlice }

func (h hp2) Less(i, j int) bool  { return h.IntSlice[i] < h.IntSlice[j] }
func (h *hp2) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp2) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
