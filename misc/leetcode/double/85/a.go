package double_85

import (
	"math"
	"strings"
)

// 1
func minimumRecolors(blocks string, k int) int {

	ans := math.MaxInt64
	for i := 0; i <= len(blocks)-k; i++ {
		ans = min(ans, strings.Count(blocks[i:i+k], "W"))
	}
	return ans
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// 2
func secondsToRemoveOccurrences(s string) int {

	ans := 0
	d := []byte(s)
	for {
		i, found := 1, false

		for i < len(d) {
			if d[i] == '1' && d[i-1] == '0' {
				d[i] = '0'
				d[i-1] = '1'
				i += 2
				found = true
				continue
			}
			i++
		}
		if !found {
			break
		}
		ans++
	}
	return ans
}

// 3
func shiftingLetters(s string, shifts [][]int) string {
	n := len(s)
	d := make([]int, n+5)
	for _, p := range shifts {
		s, t, e := p[0], p[1], p[2]
		if e == 0 {
			e = -1
		}
		d[s] += e
		d[t+1] -= e
	}
	const mod = 26
	c := 0
	t := []byte(s)
	for i := range t {
		c += d[i]
		c = (c%mod + mod) % mod
		t[i] -= 'a'
		t[i] = (t[i] + byte(c)) % mod
		t[i] += 'a'
	}
	return string(t)
}

// 4
type uf struct {
	fa, sz  []int
	groups  int
	maxSize int
}

func newUnionFind(n int, a []int) uf {
	fa := make([]int, n)
	sz := make([]int, n)
	for i := range fa {
		fa[i] = i
		sz[i] = a[i]
	}
	return uf{fa, sz, n, 1}
}

func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *uf) merge(from, to int) (isNewMerge int) {
	x, y := u.find(from), u.find(to)
	if x == y {
		return 0
	}
	u.fa[x] = y
	u.sz[y] += u.sz[x]
	if u.sz[y] > u.maxSize {
		u.maxSize = u.sz[y]
	}
	u.groups--
	return u.sz[y]
}

func (u uf) same(x, y int) bool { return u.find(x) == u.find(y) }
func (u uf) size(x int) int     { return u.sz[u.find(x)] }

func maximumSegmentSum(a []int, removeQueries []int) (ans []int64) {
	n := len(a)
	a = append(a, 0)

	ans = make([]int64, n)
	u := newUnionFind(n+1, a)
	for i := len(removeQueries) - 1; i > 0; i-- {
		v := removeQueries[i]
		u.merge(v, v+1)
		end := u.find(v + 1)
		sz := u.size(v)
		sz -= a[end]
		ans[i-1] = int64(max(int64(ans[i]), int64(sz)))
	}
	return
}

func max(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}

// 并查集
// maximumSegmentSum([]int{1,2,5,6,1}, []int{0,3,2,4,1})
func maximumSegmentSumS2(nums []int, removeQueries []int) (ans []int64) {
	n := len(nums)
	fa := make([]int, n+1)
	for i := range fa {
		fa[i] = i
	}

	sum := make([]int64, n+1)
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	ans = make([]int64, n)
	for i := n - 1; i > 0; i-- {
		x := removeQueries[i]
		to := find(x + 1)
		fa[x] = to // 合并 x 和 x+1
		sum[to] += sum[x] + int64(nums[x])
		ans[i-1] = max(ans[i], sum[to])
	}
	return
}
