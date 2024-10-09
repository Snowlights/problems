package c_221

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

// https://atcoder.jp/contests/abc221/tasks/abc221_e
//
//输入 n(≤3e5) 和长为 n 的数组 a (1≤a[i]≤1e9)。
//输出有多少个 a 的长度至少为 2 的子序列，满足子序列的第一项 ≤ 子序列的最后一项。
//由于答案很大，输出答案模 998244353 的结果。
//
//注：子序列不要求连续。

// https://atcoder.jp/contests/abc221/submissions/35791225
//这是一道基于逆序对的变形题。
//用树状数组实现，这里的问题是，如何把 2 的幂次（中间的子序列的个数）也考虑进去。
//我们可以把 2^(i-j) 转换成 2^i / 2^j，那么把 2^j 的逆元加到树状数组中即可。
//注意需要离散化。

func AtCoderABC221E(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	const (
		mod  int64 = 998244353
		mod2 int64 = (mod + 1) / 2
	)
	// 逆元的定义
	// mod2 * a % mod = 1, mod2 就是a的逆元
	// mod2 = (mod + 1) / a
	var n int
	fmt.Fscan(r, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(r, &a[i])
	}
	b := append([]int{}, a...)
	sort.Ints(b)

	tree := make([]int64, n+1)
	p, invp, ans := int64(1), int64(1), int64(0)
	for _, v := range a {
		v = sort.SearchInts(b, v) + 1
		for j := v; j > 0; j &= j - 1 {
			ans = (ans + tree[j]*p) % mod
		}
		p = p * 2 % mod
		invp = invp * mod2 % mod
		for ; v <= n; v += v & -v {
			tree[v] = (tree[v] + invp) % mod
		}
	}
	fmt.Fprintln(w, ans)
}

const (
	mod  int64 = 998244353
	mod2 int64 = (mod + 1) / 2
	// 逆元的定义
	// mod2 * a % mod = 1, mod2 就是a的逆元
	// mod2 = (mod + 1) / a
)

func AtCoderABC221EV2(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	fmt.Fscan(r, &n)
	a := make([]int, n)
	for i := range a {
		fmt.Fscan(r, &a[i])
	}
	b := append([]int{}, a...)
	sort.Ints(b)

	t := newFenwickTree(int64(n+1), nil)
	p, invp, ans := int64(1), int64(1), int64(0)
	for _, v := range a {
		v = sort.SearchInts(b, v) + 1
		val := t.query(1, int64(v))
		ans = (ans + val*p%mod + mod) % mod
		p = p * 2 % mod
		invp = invp * mod2 % mod
		t.add(int64(v), invp)
	}
	fmt.Fprintln(w, ans)
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
	val1 := i * val % mod
	for ; i < int64(len(f.tree)); i += i & -i {
		f.tree[i] += val % mod
		f.diff[i] += val1 % mod
	}
}

// 求前缀和 [0,i]
// 0<=i<=n
func (f fenwick) sum(i int64) (res int64) {
	for ; i > 0; i &= i - 1 {
		res += f.tree[i] % mod
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
	return (f.sum(r)%mod - f.sum(l-1)%mod + mod) % mod
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
