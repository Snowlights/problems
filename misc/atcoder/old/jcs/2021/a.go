package _021

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

// https://atcoder.jp/contests/jsc2021/tasks/jsc2021_f
//
//输入 n m q (≤2e5)，初始你有长为 n 的数组 a，长为 m 的数组 b，元素值都为 0，下标从 1 开始。
//然后输入 q 个询问，每个询问形如 t x y (1≤y≤1e8)。
//t=1，表示把 a[x]=y；t=2，表示把 b[x]=y。
//每次修改后，输出 ∑∑max(a[i],b[j])，这里 i 取遍 [1,n]，j 取遍 [1,m]。

// 提示 1：元素的顺序并没有那么重要。
//提示 2：每次修改时，只需要知道另一个数组中，有多少个数比自己小（这些数的个数乘上自己），以及≥自己的数的和是多少。
//最好写的应该是离散化+树状数组。

func AtCoderJCS2021F(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, q int
	fmt.Fscan(r, &n, &m, &q)
	a, b := make([]int64, n), make([]int64, m)
	t, x, y := make([]int64, q), make([]int64, q), make([]int64, q)
	h := make(map[int64]int64)
	for i := 0; i < q; i++ {
		fmt.Fscan(r, &t[i], &x[i], &y[i])
		h[y[i]] = 1
	}
	yList, yMap := make([]int, len(h)), make(map[int64]int64)
	for y := range h {
		yList = append(yList, int(y))
	}
	sort.Ints(yList)
	for i, v := range yList {
		yMap[int64(v)] = int64(i)
	}
	fa, fb := newFenwickTree(int64(len(yList)), nil), newFenwickTree(int64(len(yList)), nil)
	ca, cb := newFenwickTree(int64(len(yList)), nil), newFenwickTree(int64(len(yList)), nil)

	sum := int64(0)
	ca.add(1, int64(n))
	cb.add(1, int64(m))

	for i, v := range t {
		x, y := x[i]-1, y[i]
		switch v {
		case 1:
			sum -= cb.query(1, yMap[a[x]]) * a[x]
			sum += cb.query(1, yMap[y]) * y
			sum += fb.query(1, yMap[a[x]])
			sum -= fb.query(1, yMap[y])
			fa.add(yMap[a[x]], -int64(a[x]))
			fa.add(yMap[y], int64(y))
			ca.add(yMap[a[x]], -1)
			ca.add(yMap[y], 1)
			a[x] = y
		case 2:
			sum -= ca.query(1, yMap[b[x]]) * b[x]
			sum += ca.query(1, yMap[y]) * y
			sum += fa.query(1, yMap[b[x]])
			sum -= fa.query(1, yMap[y])
			fb.add(yMap[b[x]], -int64(b[x]))
			fb.add(yMap[y], int64(y))
			cb.add(yMap[b[x]], -1)
			cb.add(yMap[y], 1)
			b[x] = y
		}
		fmt.Fprintln(w, sum)
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

// 差分树状数组，可用于区间更新+单点查询 queryOne(i) = a[i] + sum(i) // a 从 1 开始
// r+1 即使超过 n 也没关系，因为不会用到
// i >= 1
func (f fenwick) queryOne(i int64) int64 { return f.a[i] + f.sum(i) }

// [l,r]
func (f fenwick) addRange(l, r, val int64) { f.add(l, val); f.add(r+1, -val) }
