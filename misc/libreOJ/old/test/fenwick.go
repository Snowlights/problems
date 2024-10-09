package test

import (
	"bufio"
	"fmt"
	"io"
)

// 单点修改、区间查询
func LibreOJP130(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, q int64
	fmt.Fscan(r, &n, &q)
	a := make([]int64, n+1)
	t := newFenwickTree(n, nil)
	for i := int64(1); i <= n; i++ {
		fmt.Fscan(r, &a[i])
		t.add(i, a[i])
	}
	var op, left, right int64
	var x int64
	for i := int64(0); i < q; i++ {
		fmt.Fscan(r, &op)
		switch op {
		case 1:
			fmt.Fscan(r, &left, &x)
			t.add(left, x)
		case 2:
			fmt.Fscan(r, &left, &right)
			fmt.Fprintln(w, t.query(left, right))
		}
	}

}

// 区间修改、单点查询
func LibreOJP131(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, q int64
	fmt.Fscan(r, &n, &q)
	a := make([]int64, n+1)
	for i := int64(1); i <= n; i++ {
		fmt.Fscan(r, &a[i])
	}
	var op, left, right int64
	var x int64
	t := newFenwickTree(n, a)
	for i := int64(0); i < q; i++ {
		fmt.Fscan(r, &op)
		switch op {
		case 1:
			fmt.Fscan(r, &left, &right, &x)
			t.addRange(left, right, x)
		case 2:
			fmt.Fscan(r, &left)
			fmt.Fprintln(w, t.queryOne(left))
		}
	}

}

// 区间修改，区间查询
func LibreOJP132(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, q int64
	fmt.Fscan(r, &n, &q)
	a := make([]int64, n+1)
	t := newFenwickTree(n, nil)
	for i := int64(1); i <= n; i++ {
		fmt.Fscan(r, &a[i])
		t.addRange(i, i, a[i])
	}
	var op, left, right int64
	var x int64
	for i := int64(0); i < q; i++ {
		fmt.Fscan(r, &op)
		switch op {
		case 1:
			fmt.Fscan(r, &left, &right, &x)
			t.addRange(left, right, x)
		case 2:
			fmt.Fscan(r, &left, &right)
			fmt.Fprintln(w, t.queryDiff(left, right))
		}
	}

}

type fenwick struct {
	a    []int64 // 原始数据，单点查询+区间更新传入
	tree []int64 // 树状数组
	diff []int64 // 差分数组, 用于区间加、区间求和
}

func newFenwickTree(n int64, a []int64) fenwick {
	return fenwick{a, make([]int64, n+1), make([]int64, n+1)}
}

// 位置 i 增加 val
// 1<=i<=n
func (f fenwick) add(i, val int64) {
	val1 := i * val
	for ; i < int64(len(f.tree)); i += i & -i {
		f.tree[i] += val
		f.diff[i] += val1
	}
}

// 求前缀和 [0,i]
// 0<=i<=n
func (f fenwick) sum(i int64) (res int64) {
	for ; i > 0; i &= i - 1 {
		res += f.tree[i]
	}
	return
}

func (f fenwick) sumDiff(i int64) (res int64) {
	for ; i > 0; i &= i - 1 {
		res += f.diff[i]
	}
	return
}

// 求区间和 [l,r]
// 1<=l<=r<=n
func (f fenwick) query(l, r int64) int64 {
	return f.sum(r) - f.sum(l-1)
}

func (f fenwick) queryDiff(l, r int64) int64 {
	return (r+1)*f.sum(r) - l*f.sum(l-1) - (f.sumDiff(r) - f.sumDiff(l-1))
}

// 求权值树状数组第 k 小的数（k > 0）
// 这里 tree[i] 表示 i 的个数
// 返回最小的 x 满足 ∑i=[1..x] tree[i] >= k
// 思路类似倍增的查询，不断寻找 ∑<k 的数，最后 +1 就是答案
// - 代码见下面的 rangeMex
func (f fenwick) kth(k int64) (res int64) {
	const log = 17 // bits.Len(uint(n))
	for b := int64(1 << (log - 1)); b > 0; b >>= 1 {
		if next := res | b; next < int64(len(f.tree)) && k > f.tree[next] {
			k -= f.tree[next]
			res = next
		}
	}
	return res + 1
}

// 差分树状数组，可用于区间更新+单点查询 queryOne(i) = a[i] + sum(i) // a 从 1 开始
// r+1 即使超过 n 也没关系，因为不会用到
// i >= 1
func (f fenwick) queryOne(i int64) int64 { return f.a[i] + f.sum(i) }

// [l,r]
func (f fenwick) addRange(l, r, val int64) { f.add(l, val); f.add(r+1, -val) }
