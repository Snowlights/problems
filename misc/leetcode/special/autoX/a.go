package autoX

import (
	"container/heap"
	"math"
	"sort"
)

// 1
func getLengthOfWaterfallFlow(num int, block []int) int {
	h := hp{}
	for i := 0; i < num; i++ {
		heap.Push(&h, pair{
			sum: 0,
			i:   i,
		})
	}
	for _, b := range block {
		p := heap.Pop(&h).(pair)
		p.sum += b
		heap.Push(&h, p)
	}
	ans := 0
	for _, v := range h {
		ans = max(ans, v.sum)
	}
	return ans
}

type pair struct{ sum, i int }
type hp []pair

func (h hp) Len() int { return len(h) }
func (h hp) Less(i, j int) bool {
	return h[i].sum < h[j].sum || (h[i].sum == h[i].sum && h[j].i < h[j].i)
}
func (h hp) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *hp) Push(v interface{}) { *h = append(*h, v.(pair)) }
func (h *hp) Pop() interface{}   { a := *h; v := a[len(a)-1]; *h = a[:len(a)-1]; return v }

// 2
func honeyQuotes(handles [][]int) []float64 {

	total, num, h := 0.0, 0.0, make(map[float64]int)
	ans := make([]float64, 0)
	for _, handle := range handles {

		switch handle[0] {
		case 1:
			val := float64(handle[1])
			h[val]++
			num++
			total += val
		case 2:
			val := float64(handle[1])
			h[val]--
			num--
			total -= val
			if h[val] == 0 {
				delete(h, val)
			}
		case 3:
			if num == 0 {
				ans = append(ans, -1)
				continue
			}
			ans = append(ans, total/num)
		case 4:
			if num == 0 {
				ans = append(ans, -1)
				continue
			}
			tmp, avg := 0.0, total/num
			for k, v := range h {
				tmp += (k - avg) * (k - avg) * float64(v)
			}
			ans = append(ans, tmp/num)
		}
	}

	return ans
}

// 3
func minCostToTravelOnDays(days []int, tickets [][]int) int64 {
	dp := make([]int64, len(days)+1)
	for i := range dp {
		dp[i] = math.MaxInt64
	}
	dp[0] = 0
	for i := 0; i < len(days); i++ {
		for _, v := range tickets {
			d, p := v[0], v[1]
			j := sort.SearchInts(days, days[i]+d)
			dp[j] = minInt64(dp[j], dp[i]+int64(p))
		}
	}

	return dp[len(days)]
}

func maxInt64(a, b int64) int64 {
	if b > a {
		return b
	}
	return a
}

func minInt64(a, b int64) int64 {
	if a > b {
		return b
	}
	return a
}

// 4

// 点
type Point struct {
	x, y int
}

// 圆
type Circle struct {
	x, y, r int
}

// 线段和圆是否相交
func judge(p1, p2 *Point, c *Circle) bool {
	flag1 := (p1.x-c.x)*(p1.x-c.x)+(p1.y-c.y)*(p1.y-c.y) <= c.r*c.r
	flag2 := (p2.x-c.x)*(p2.x-c.x)+(p2.y-c.y)*(p2.y-c.y) <= c.r*c.r
	//情况一、两点都在圆内 :一定不相交
	if flag1 && flag2 {
		return false
	} else if flag1 || flag2 {
		//情况二、一个点在圆内，一个点在圆外：一定相交
		return true
	} else {
		//将直线p1p2化为一般式：Ax+By+C=0的形式。先化为两点式，然后由两点式得出一般式
		A := p1.y - p2.y
		B := p2.x - p1.x
		C := p1.x*p2.y - p2.x*p1.y
		//使用距离公式判断圆心到直线ax+by+c=0的距离是否大于半径
		dist1 := A*c.x + B*c.y + C
		dist1 *= dist1
		dist2 := (A*A + B*B) * c.r * c.r
		if dist1 > dist2 {
			return false
		}
		//点到直线距离大于半径r  不相交
		angle1 := (c.x-p1.x)*(p2.x-p1.x) + (c.y-p1.y)*(p2.y-p1.y)
		angle2 := (c.x-p2.x)*(p1.x-p2.x) + (c.y-p2.y)*(p1.y-p2.y)
		if angle1 > 0 && angle2 > 0 {
			return true
		}
		//余弦都为正，则是锐角 相交
		return false
	}
}

type Line struct {
	id             int
	x1, y1, x2, y2 float64
}

type LineInt struct {
	id             int
	x1, y1, x2, y2 int
}

// 两线段是否相交
func intersection(l1, l2 *LineInt) bool {
	//快速排斥实验
	if (max(l1.x1, l1.x2) < (min(l2.x1, l2.x2))) ||
		(max(l1.y1, l1.y2) < (min(l2.y1, l2.y2))) ||
		(max(l2.x1, l2.x2) < (min(l1.x1, l1.x2))) ||
		(max(l2.y1, l2.y2) < (min(l1.y1, l1.y2))) {
		return false
	}
	//跨立实验
	if (((l1.x1-l2.x1)*(l2.y2-l2.y1)-(l1.y1-l2.y1)*(l2.x2-l2.x1))*
		((l1.x2-l2.x1)*(l2.y2-l2.y1)-(l1.y2-l2.y1)*(l2.x2-l2.x1))) > 0 ||
		(((l2.x1-l1.x1)*(l1.y2-l1.y1)-(l2.y1-l1.y1)*(l1.x2-l1.x1))*
			((l2.x2-l1.x1)*(l1.y2-l1.y1)-(l2.y2-l1.y1)*(l1.x2-l1.x1))) > 0 {
		return false
	}

	return true
}

func minFloat(a, b float64) float64 {
	if a < b {
		return a
	}
	return b
}

func maxFloat(a, b float64) float64 {
	if b > a {
		return b
	}
	return a
}

func helper(a, b, c []float64) float64 {
	w1, v1, w2, v2 := b[0]-a[0], b[1]-a[1], c[0]-a[0], c[1]-a[1]
	return w1*v2 - w2*v1
}

// 两线段是否相交
func check(l1, l2 *Line) bool {
	p1, p2 := []float64{l1.x1, l1.y1}, []float64{l1.x2, l1.y2}
	p3, p4 := []float64{l2.x1, l2.y1}, []float64{l2.x2, l2.y2}
	if maxFloat(p1[0], p2[0]) >= minFloat(p3[0], p4[0]) &&
		maxFloat(p3[0], p4[0]) >= minFloat(p1[0], p2[0]) &&
		maxFloat(p1[1], p2[1]) >= minFloat(p3[1], p4[1]) &&
		maxFloat(p3[1], p4[1]) >= minFloat(p2[1], p1[1]) {
		if helper(p1, p2, p3)*helper(p1, p2, p4) <= 0 &&
			helper(p3, p4, p1)*helper(p3, p4, p2) <= 0 {
			return true
		}
		return false
	} else {
		return false
	}
}

// 两圆是否相交
func isIntersection(c1, c2 *Circle) bool {
	r := c1.r + c2.r
	x := c1.x - c2.x
	y := c1.y - c2.y
	d := int(math.Sqrt(float64(x*x + y*y)))
	if abs(c1.r-c2.r) > d {
		return false
	}
	if d > r {
		return false
	}

	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

func cal(a, b []int) bool {
	switch len(a) {
	case 3:
		switch len(b) {
		case 3:
			return isIntersection(&Circle{
				x: int(a[0]),
				y: int(a[1]),
				r: int(a[2]),
			}, &Circle{
				x: int(b[0]),
				y: int(b[1]),
				r: int(b[2]),
			})
		case 4:
			return judge(&Point{
				x: int(b[0]),
				y: int(b[1]),
			}, &Point{
				x: int(b[2]),
				y: int(b[3]),
			}, &Circle{
				x: int(a[0]),
				y: int(a[1]),
				r: int(a[2]),
			})
		}
	case 4:
		switch len(b) {
		case 3:
			return judge(&Point{
				x: int(a[0]),
				y: int(a[1]),
			}, &Point{
				x: int(a[2]),
				y: int(a[3]),
			}, &Circle{
				x: int(b[0]),
				y: int(b[1]),
				r: int(b[2]),
			})
		case 4:
			return check(&Line{
				x1: float64(a[0]),
				y1: float64(a[1]),
				x2: float64(a[2]),
				y2: float64(a[3]),
			}, &Line{
				x1: float64(b[0]),
				y1: float64(b[1]),
				x2: float64(b[2]),
				y2: float64(b[3]),
			})
		}
	}
	return false
}

func antPass(geometry [][]int, path [][]int) []bool {
	g, vis := make(map[int][]int), make(map[int]bool)
	for i := 0; i < len(geometry); i++ {
		for j := i + 1; j < len(geometry); j++ {
			if cal(geometry[i], geometry[j]) {
				g[i] = append(g[i], j)
				g[j] = append(g[j], i)
			}
		}
	}

	gMap := make([]map[int]bool, 0)
	for k := range geometry {
		if vis[k] {
			continue
		}
		q, ans := []int{k}, make(map[int]bool)
		vis[k] = true
		ans[k] = true
		for len(q) > 0 {
			tmp := q
			q = nil
			for _, v := range tmp {
				for _, vv := range g[v] {
					if ans[vv] {
						continue
					}
					vis[vv] = true
					ans[vv] = true
					q = append(q, vv)
				}
			}
		}
		gMap = append(gMap, ans)
	}

	res := make([]bool, 0, len(path))
	for _, p := range path {
		found := false
		for _, g := range gMap {
			if g[p[0]] && g[p[1]] {
				res = append(res, true)
				found = true
				break
			}
		}
		if !found {
			res = append(res, false)
		}
	}
	return res
}
