//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1334F(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var n, m int
	Fscan(in, &n)
	a := make([]int, n)
	pos := make([][]int, n+1)
	negSum := make([]int, n+1)
	for i := range a {
		Fscan(in, &a[i])
		pos[a[i]] = append(pos[a[i]], i)
	}
	cost := make([]int, n)
	for i := range cost {
		Fscan(in, &cost[i])
		negSum[i+1] = negSum[i]
		if cost[i] < 0 {
			negSum[i+1] += cost[i]
		}
	}
	Fscan(in, &m)
	b := make([]int, m)
	for i := range b {
		Fscan(in, &b[i])
	}

	t := newFenwickTree(n, nil)
	f := make([]int, n)
	for i := range f {
		f[i] = 1e18
	}
	s := 0
	for i, v := range a {
		if v > b[0] {
			if cost[i] > 0 {
				t.add(i, cost[i])
			}
		} else if v == b[0] {
			f[i] = s
		}
		s += cost[i]
	}

	pid := 1
	for pid <= b[0] {
		pid++
	}

	for bi := 1; bi < m; bi++ {
		p := pos[b[bi-1]]
		j := 0
		mn := int(1e18)
		pre := -1
		for _, i := range pos[b[bi]] {
			if pre >= 0 {
				mn += t.query(pre, i-1) + negSum[i] - negSum[pre]
			}
			for j < len(p) && p[j] < i {
				res := f[p[j]] + t.query(p[j]+1, i-1) + negSum[i] - negSum[p[j]+1]
				mn = min(mn, res)
				j++
			}
			f[i] = mn
			pre = i
		}
		for pid <= b[bi] {
			for _, i := range pos[pid] {
				if cost[i] > 0 {
					t.add(i, -cost[i])
				}
			}
			pid++
		}
	}

	ans := int(1e18)
	s = 0
	for i := n - 1; i >= 0; i-- {
		v := a[i]
		if v == b[m-1] {
			ans = min(ans, f[i]+s)
		}
		if v > b[m-1] || cost[i] < 0 {
			s += cost[i]
		}
	}
	if ans >= 1e17 {
		Fprint(out, "NO")
	} else {
		Fprintln(out, "YES")
		Fprint(out, ans)
	}

}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

type fenwick struct {
	a    []int // 原始数据，单点查询+区间更新传入
	tree []int // 树状数组
	diff []int // 差分数组, 用于区间加、区间求和
}

func newFenwickTree(n int, a []int) fenwick {
	return fenwick{a, make([]int, n+1), make([]int, n+1)}
}

// 位置 i 增加 val
// 1<=i<=n
func (f fenwick) add(i, val int) {
	i++
	val1 := i * val
	for ; i < int(len(f.tree)); i += i & -i {
		f.tree[i] += val
		f.diff[i] += val1
	}
}

// 求前缀和 [0,i]
// 0<=i<=n
func (f fenwick) sum(i int) (res int) {
	i++
	for ; i > 0; i &= i - 1 {
		res += f.tree[i]
	}
	return
}

func (f fenwick) sumDiff(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f.diff[i]
	}
	return
}

// 求区间和 [l,r]
// 1<=l<=r<=n
func (f fenwick) query(l, r int) int {
	return f.sum(r) - f.sum(l-1)
}

func (f fenwick) queryDiff(l, r int) int {
	return (r+1)*f.sum(r) - l*f.sum(l-1) - (f.sumDiff(r) - f.sumDiff(l-1))
}

// 差分树状数组，可用于区间更新+单点查询 queryOne(i) = a[i] + sum(i) // a 从 1 开始
// r+1 即使超过 n 也没关系，因为不会用到
// i >= 1
func (f fenwick) queryOne(i int) int { return f.a[i] + f.sum(i) }

// [l,r]
func (f fenwick) addRange(l, r, val int) { f.add(l, val); f.add(r+1, -val) }

func main() { CF1334F(os.Stdin, os.Stdout) }
