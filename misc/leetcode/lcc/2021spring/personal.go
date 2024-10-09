package _021spring

import (
	"container/heap"
	"sort"
)

// 1
func purchasePlans(a []int, tar int) (ans int) {
	const mod int = 1e9 + 7

	sort.Ints(a)
	n := len(a)
	for i := 1; i < n; i++ {
		v := a[i]
		j := sort.SearchInts(a[:i], tar-v+1)
		ans += j
	}

	ans %= mod
	return
}

// 2
func orchestraLayout(n int, r int, c int) (ans int) {
	k := min(min(r, n-1-r), min(c, n-1-c))
	cnt := k*(4*n-4) - 4*(k-1)*k
	//println(cnt)
	if r == k {
		d := c - k + 1
		cnt += d
		ans = (cnt-1)%9 + 1
		return
	}
	cnt += n - 2*k - 1
	if c+k == n-1 {
		d := r - k + 1
		cnt += d
		ans = (cnt-1)%9 + 1
		return
	}
	cnt += n - 2*k - 1
	if r+k == n-1 {
		d := n - c - k
		cnt += d
		ans = (cnt-1)%9 + 1
		return
	}
	cnt += n - 2*k - 1
	d := n - r - k
	cnt += d
	ans = (cnt-1)%9 + 1
	return
}

// 3
type hp struct{ sort.IntSlice }

func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *hp) push(v int) { heap.Push(h, v) }
func (h *hp) pop() int   { return heap.Pop(h).(int) }
func (h hp) empty() bool { return len(h.IntSlice) == 0 }
func (h hp) top() int    { return h.IntSlice[0] }

func magicTower(a []int) (ans int) {
	sum := 1
	for _, v := range a {
		sum += v
	}
	if sum < 1 {
		return -1
	}

	s := 1
	h := &hp{}
	for _, v := range a {
		if v < 0 {
			h.push(v)
		}
		s += v
		if s < 1 {
			w := h.pop()
			s -= w
			ans++
		}
	}
	return
}

// 4
func escapeMaze(maze [][]string) bool {
	var T, n, m int
	fx := [5][2]int{{0, 1}, {1, 0}, {0, 0}, {0, -1}, {-1, 0}}
	vis := [101][55][55][2][2]bool{}
	var dfs func(x, y, t, a, b, c, d int) bool
	dfs = func(x, y, t, a, b, c, d int) bool {
		if x == n-1 && y == m-1 {
			return true
		}
		if t == T || vis[t][x][y][a][b] {
			return false
		}
		for i := 0; i < 5; i++ {
			dx, dy := x+fx[i][0], y+fx[i][1]
			if dx < 0 || dy < 0 || dx >= n || dy >= m {
				continue
			}
			if maze[t+1][dx][dy] == '#' {
				if c == dx && d == dy {
					if dfs(dx, dy, t+1, a, b, c, d) {
						return true
					}
				} else {
					if a > 0 {
						if dfs(dx, dy, t+1, 0, b, c, d) {
							return true
						}
					}
					if b > 0 {
						if dfs(dx, dy, t+1, a, 0, dx, dy) {
							return true
						}
					}
				}
			} else {
				if dfs(dx, dy, t+1, a, b, c, d) {
					return true
				}
			}
		}

		vis[t][x][y][a][b] = true
		return false
	}

	T, n, m = len(maze)-1, len(maze[0]), len(maze[0][0])
	return dfs(0, 0, 0, 1, 1, -1, -1)
}

// 5
func processTasks(tasks [][]int) int {
	sort.Slice(tasks, func(i, j int) bool {
		return tasks[i][1] < tasks[j][1]
	})
	type pair struct {
		first, second int
	}
	var parts []*pair
	n := len(tasks)
	parts = append(parts, &pair{tasks[0][1] - tasks[0][2] + 1, tasks[0][1]})
	for i := 1; i < n; i++ {
		total, begin := 0, tasks[i][0]
		for k := len(parts) - 1; k >= 0; k-- {
			if parts[k].second < begin {
				break
			}
			total += parts[k].second - max(begin, parts[k].first) + 1
		}
		if total >= tasks[i][2] {
			continue
		}
		need := tasks[i][2] - total
		begin = tasks[i][1] - need + 1
		end := tasks[i][1]
		for len(parts) > 0 {
			last := parts[len(parts)-1]
			if last.second >= begin-1 {
				parts = parts[:len(parts)-1]
				common := last.second - max(last.first, begin) + 1
				begin = min(begin, last.first) - common
			} else {
				break
			}
		}
		parts = append(parts, &pair{begin, end})
	}

	res := 0
	for _, part := range parts {
		res += part.second - part.first + 1
	}
	return res
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
