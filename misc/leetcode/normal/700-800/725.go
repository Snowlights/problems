package _00_800

import (
	"container/heap"
	"math"
	"strconv"
)

// 733

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	var (
		dx = []int{1, 0, 0, -1}
		dy = []int{0, 1, -1, 0}
	)
	var dfs func(image [][]int, x, y, color, newColor int)
	dfs = func(image [][]int, x, y, color, newColor int) {
		if image[x][y] == color {
			image[x][y] = newColor
			for i := 0; i < 4; i++ {
				mx, my := x+dx[i], y+dy[i]
				if mx >= 0 && mx < len(image) && my >= 0 && my < len(image[0]) {
					dfs(image, mx, my, color, newColor)
				}
			}
		}
	}
	currColor := image[sr][sc]
	if currColor != newColor {
		dfs(image, sr, sc, currColor, newColor)
	}
	return image
}

// 738

func monotoneIncreasingDigits(n int) int {
	s := []byte(strconv.Itoa(n))

	i := 1
	for i < len(s) && s[i] >= s[i-1] {
		i++
	}
	if i != len(s) {
		for i > 0 && s[i] < s[i-1] {
			s[i-1]--
			i--
		}
		for i++; i < len(s); i++ {
			s[i] = '9'
		}
	}

	n, _ = strconv.Atoi(string(s))
	return n
}

// 740
func deleteAndEarn(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	sum := make([]int, 1e4+5)
	for _, v := range nums {
		sum[v] += v
	}
	dp := make([]int64, 1e4+5)
	dp[1] = int64(sum[1])
	dp[2] = max(int64(dp[1]), int64(sum[2]))
	for i := 3; i < 1e4+5; i++ {
		dp[i] = max(dp[i-2]+int64(sum[i]), dp[i-1])
	}
	return int(dp[1e4+4])
}

// 743

func networkDelayTime(times [][]int, n int, k int) int {
	type edge struct{ to, wt int }
	const (
		inf int64 = 1e18
	)
	dijkstra := func(g [][]edge, start int) []int64 {
		dist := make([]int64, n)
		for i := range dist {
			dist[i] = inf
		}
		dist[start] = 0
		// 虽然可以用 dist 来判断是否需要 relax，但是对于一些变形题，用 vis 是最稳的
		vis := make([]bool, n)
		fa := make([]int, n)
		for i := range fa {
			fa[i] = -1
		}
		h := vdHeap{{start, 0}}
		for len(h) > 0 {
			p := h.pop()
			v := p.v
			if vis[v] { // p.dis > dist[v]
				continue
			}
			vis[v] = true
			for _, e := range g[v] {
				w := e.to
				if newD := dist[v] + int64(e.wt); newD < dist[w] {
					dist[w] = newD
					fa[w] = v
					h.push(vdPair{w, dist[w]})
				}
			}
		}
		return dist
	}
	g := make([][]edge, n)
	for _, v := range times {
		from, to, weight := v[0]-1, v[1]-1, v[2]
		g[from] = append(g[from], edge{to: to, wt: weight})
	}

	dist := dijkstra(g, k-1)

	ans := 0
	for i, v := range dist {
		if i == k-1 {
			continue
		}
		if v == inf {
			return -1
		}
		ans = int(max(int64(ans), int64(v)))
	}
	return int(ans)
}

type vdPair struct {
	v   int
	dis int64
}
type vdHeap []vdPair

func (h vdHeap) Len() int              { return len(h) }
func (h vdHeap) Less(i, j int) bool    { return h[i].dis < h[j].dis }
func (h vdHeap) Swap(i, j int)         { h[i], h[j] = h[j], h[i] }
func (h *vdHeap) Push(v interface{})   { *h = append(*h, v.(vdPair)) }
func (h *vdHeap) Pop() (v interface{}) { a := *h; *h, v = a[:len(a)-1], a[len(a)-1]; return }
func (h *vdHeap) push(v vdPair)        { heap.Push(h, v) }
func (h *vdHeap) pop() vdPair          { return heap.Pop(h).(vdPair) }

// 746
func minCostClimbingStairs(cost []int) int {
	dp := make([]int, len(cost))
	for i := range dp {
		dp[i] = math.MaxInt32
	}
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i, v := range cost {
		if i > 0 {
			dp[i] = min(dp[i], dp[i-1]+v)
		}
		if i > 1 {
			dp[i] = min(dp[i], dp[i-2]+v)
		}
	}

	return min(dp[len(cost)-1], dp[len(cost)-2])
}
