package simple_340

import (
	"math"
	"sort"
)

// 1
func isPrime(n int) bool {
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return n >= 2 // 1 不是质数
}

func diagonalPrime(nums [][]int) (ans int) {
	for i, row := range nums {
		if x := row[i]; x > ans && isPrime(x) {
			ans = x
		}
		if x := row[len(nums)-1-i]; x > ans && isPrime(x) {
			ans = x
		}
	}
	return
}

// 2
func distance(nums []int) []int64 {
	groups := map[int][]int{}
	for i, x := range nums {
		groups[x] = append(groups[x], i) // 相同元素分到同一组，记录下标
	}
	ans := make([]int64, len(nums))
	for _, a := range groups {
		n := len(a)
		s := int64(0)
		for _, x := range a {
			s += int64(x - a[0]) // a[0] 到其它下标的距离之和
		}
		ans[a[0]] = s
		for i := 1; i < n; i++ {
			// 从计算 a[i-1] 到计算 a[i]，考虑 s 增加了多少
			s += int64(i*2-n) * int64(a[i]-a[i-1])
			ans[a[i]] = s
		}
	}
	return ans
}

// 3
func minimizeMax(nums []int, p int) int {
	sort.Ints(nums)
	return sort.Search(1e9, func(mx int) bool {
		cnt := 0
		for i := 0; i < len(nums)-1; i++ {
			if nums[i+1]-nums[i] <= mx { // 都选
				cnt++
				i++
			}
		}
		return cnt >= p
	})
}

// 4
func minimumVisitedCells(grid [][]int) (mn int) {
	m, n := len(grid), len(grid[0])
	type pair struct{ x, i int }
	colSt := make([][]pair, n) // 每列的单调栈
	for i := m - 1; i >= 0; i-- {
		st := []pair{} // 当前行的单调栈
		for j := n - 1; j >= 0; j-- {
			st2 := colSt[j]
			mn = math.MaxInt32
			if i == m-1 && j == n-1 { // 特殊情况：已经是终点
				mn = 0
			} else if g := grid[i][j]; g > 0 {
				// 在单调栈上二分
				k := j + g
				k = sort.Search(len(st), func(i int) bool { return st[i].i <= k })
				if k < len(st) {
					mn = min(mn, st[k].x)
				}
				k = i + g
				k = sort.Search(len(st2), func(i int) bool { return st2[i].i <= k })
				if k < len(st2) {
					mn = min(mn, st2[k].x)
				}
			}

			if mn < math.MaxInt32 {
				mn++ // 加上 (i,j) 这个格子
				// 插入单调栈
				for len(st) > 0 && mn <= st[len(st)-1].x {
					st = st[:len(st)-1]
				}
				st = append(st, pair{mn, j})
				for len(st2) > 0 && mn <= st2[len(st2)-1].x {
					st2 = st2[:len(st2)-1]
				}
				colSt[j] = append(st2, pair{mn, i})
			}
		}
	}
	// 最后一个算出的 mn 就是 f[0][0]
	if mn == math.MaxInt32 {
		return -1
	}
	return
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}
