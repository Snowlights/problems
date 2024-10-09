package main

import (
	"sort"
)

func resultArray(nums []int) []int {
	a, h := make([]int, 0, len(nums)), make(map[int]bool)
	for _, v := range nums {
		if h[v] {
			continue
		}
		a = append(a, v)
		h[v] = true
	}
	sort.Ints(a)
	n := len(a)
	at, bt := newFenwickTree(n, nil), newFenwickTree(n, nil)
	r1, r2 := []int{nums[0]}, []int{nums[1]}
	at.add(sort.SearchInts(a, nums[0])+1, 1)
	bt.add(sort.SearchInts(a, nums[1])+1, 1)
	for _, v := range nums[2:] {
		idx := sort.SearchInts(a, v) + 1
		v1, v2 := len(r1)-at.query(1, idx), len(r2)-bt.query(1, idx)
		if v1 > v2 || v1 == v2 && len(r1) <= len(r2) {
			r1 = append(r1, v)
			at.add(idx, 1)
		} else {
			r2 = append(r2, v)
			bt.add(idx, 1)
		}
	}
	return append(r1, r2...)
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
	val1 := i * val
	for ; i < int(len(f.tree)); i += i & -i {
		f.tree[i] += val
		f.diff[i] += val1
	}
}

// 求前缀和 [0,i]
// 0<=i<=n
func (f fenwick) sum(i int) (res int) {
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
