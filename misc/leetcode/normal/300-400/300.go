package _00_400

import (
	"math"
	"sort"
)

// 300
func lengthOfLIS(nums []int) int {

	//使得f按照升序排列
	f := []int{}
	for _, e := range nums {
		//h := e
		if i := sort.SearchInts(f, e); i < len(f) {
			f[i] = e
		} else {
			f = append(f, e)
		}
	}
	return len(f)
}

// 307
type NumArray struct {
	t fenwick
}

func Constructor307(nums []int) NumArray {
	t := newFenwickTree(int64(len(nums)+1), nil)
	for i := int64(0); i < int64(len(nums)); i++ {
		t.add(i+1, int64(nums[i]))
	}
	return NumArray{t}
}

func (this *NumArray) Update(index int, val int) {
	v := this.t.query(int64(index+1), int64(index+1))
	this.t.add(int64(index+1), int64(val)-v)
}

func (this *NumArray) SumRange(left int, right int) int {
	return int(this.t.query(int64(left+1), int64(right+1)))
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

// 309
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	n := len(prices)
	f0, f1, f2 := -prices[0], 0, 0
	for i := 1; i < n; i++ {
		newf0 := max(f0, f2-prices[i])
		newf1 := f0 + prices[i]
		newf2 := max(f1, f2)
		f0, f1, f2 = newf0, newf1, newf2
	}
	return max(f1, f2)
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 310
func findMinHeightTrees(n int, edges [][]int) []int {
	if n == 1 {
		return []int{0}
	}

	g := make([][]int, n)
	deg := make([]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x)
		deg[x]++
		deg[y]++
	}

	q := []int{}
	for i, d := range deg {
		if d == 1 {
			q = append(q, i)
		}
	}

	remainNodes := n
	for remainNodes > 2 {
		remainNodes -= len(q)
		tmp := q
		q = nil
		for _, x := range tmp {
			for _, y := range g[x] {
				deg[y]--
				if deg[y] == 1 {
					q = append(q, y)
				}
			}
		}
	}
	return q
}

// 313
func nthSuperUglyNumber(n int, primes []int) int {
	dp := make([]int, n+1)
	m := len(primes)
	pointers := make([]int, m)
	nums := make([]int, m)
	for i := range nums {
		nums[i] = 1
	}
	for i := 1; i <= n; i++ {
		minNum := math.MaxInt64
		for j := range pointers {
			minNum = min(minNum, nums[j])
		}
		dp[i] = minNum
		for j := range nums {
			if nums[j] == minNum {
				pointers[j]++
				nums[j] = dp[pointers[j]] * primes[j]
			}
		}
	}
	return dp[n]
}

// 322
func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	dp := make([]int, amount+1)
	for i := 0; i <= amount; i++ {
		dp[i] = math.MaxInt32
	}
	for _, v := range coins {
		if v > amount {
			continue
		}
		dp[v] = 1
	}
	for i := 0; i <= amount; i++ {
		for _, c := range coins {
			if c > i {
				continue
			}
			dp[i] = min(dp[i], dp[i-c]+1)
		}
	}
	if dp[amount] == math.MaxInt32 {
		return -1
	}
	return dp[amount]
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// 315
func countSmaller(nums []int) []int {

	ids := make([]int, len(nums))
	for i := range nums {
		ids[i] = i
	}
	sort.Slice(ids, func(i, j int) bool {
		return nums[ids[i]] < nums[ids[j]] || (nums[ids[i]] == nums[ids[j]] && ids[i] < ids[j])
	})
	t := newFenwickTree(int64(len(nums)), nil)
	ans := make([]int, len(nums))
	for _, v := range ids {
		ans[v] = int(t.query(int64(v)+1, int64(len(nums))))
		t.add(int64(v)+1, 1)
	}

	return ans
}

// 316
func removeDuplicateLetters(s string) string {
	h, exit := make(map[byte]int), make(map[byte]bool)
	stack := []byte{}

	for i := range s {
		h[s[i]]++
	}
	for i := range s {
		c := s[i]
		h[c]--
		if !exit[c] {
			for len(stack) > 0 && stack[len(stack)-1] > c {
				if h[stack[len(stack)-1]] == 0 {
					break
				}
				exit[stack[len(stack)-1]] = false
				stack = stack[:len(stack)-1]
			}
			stack = append(stack, c)
			exit[c] = true
		}
	}

	return string(stack)
}
