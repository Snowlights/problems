package interview

import "strconv"

// 面试题 10-10
type StreamRank struct {
	t fenwick
}

func Constructor() StreamRank {
	return StreamRank{newFenwickTree(50000+1, nil)}
}

func (this *StreamRank) Track(x int) {
	this.t.add(int64(x+1), 1)
}

func (this *StreamRank) GetRankOfNumber(x int) int {
	return int(this.t.query(1, int64(x+1)))
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

// 面试题 13

func reach(b [][]bool, i, j, m, n int) bool {
	return (i-1 >= 0 && b[i-1][j]) ||
		(j-1 >= 0 && b[i][j-1]) ||
		(i+1 < m && b[i+1][j]) ||
		(j+1 < n && b[i][j+1])
}

func dig(a, b int) int {
	ans := 0
	for a > 0 {
		ans += a % 10
		a /= 10
	}
	for b > 0 {
		ans += b % 10
		b /= 10
	}
	return ans
}

func movingCount(m int, n int, k int) int {

	b, ans := make([][]bool, m), 1
	for i := range b {
		b[i] = make([]bool, n)
	}
	b[0][0] = true

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dig(i, j) <= k && reach(b, i, j, m, n) {
				b[i][j] = true
				ans++
			}
		}
	}
	return ans
}

// 面试题 16-07
func numberOf2sInRange(n int) int {
	s := strconv.Itoa(n)
	m := len(s)
	dp := make([][]int, m)
	for i := range dp {
		dp[i] = make([]int, m)
		for j := range dp[i] {
			dp[i][j] = -1
		}
	}

	var f func(i, num int, limit bool) int
	f = func(i, num int, limit bool) (res int) {
		if i == m {
			return num
		}

		if !limit {
			dv := &dp[i][num]
			if *dv >= 0 {
				return *dv
			}
			defer func() {
				*dv = res
			}()
		}

		up := 9
		if limit {
			up = int(s[i] - '0')
		}

		for j := 0; j <= up; j++ {
			c := num
			if j == 2 {
				c++
			}
			res += f(i+1, c, limit && j == up)
		}
		return
	}

	return f(0, 0, true)
}
