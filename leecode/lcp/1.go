package lcp

import "sort"

// lcp 06
func minCount(coins []int) int {
	ans := 0
	for _, v := range coins {
		ans += v / 2
		if v%2 == 1 {
			ans++
		}
	}
	return ans
}

// lcp 07
func numWays(n int, relation [][]int, k int) int {

	g := make(map[int][]int)
	for _, v := range relation {
		g[v[0]] = append(g[v[0]], v[1])
	}
	ans := 0
	q := []int{0}
	for k > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			for _, v2 := range g[v] {
				q = append(q, v2)
				if v2 == n-1 && k == 1 {
					ans++
				}
			}
		}
		k--
	}
	return ans
}

// lcp 08
func getTriggerTime(increase [][]int, requirements [][]int) []int {

	type pair struct {
		c, r, h int
	}
	p := &pair{}
	pre := make([]*pair, 0)
	pre = append(pre, p)
	for _, v := range increase {
		now := &pair{
			c: v[0] + p.c,
			r: v[1] + p.r,
			h: v[2] + p.h,
		}
		pre = append(pre, now)
		p = now
	}
	ans := make([]int, len(requirements))
	for i := range ans {
		ans[i] = -1
	}

	for i, v := range requirements {
		idx := sort.Search(len(pre), func(i int) bool {
			return pre[i].c >= v[0] && pre[i].r >= v[1] && pre[i].h >= v[2]
		})
		if idx < len(pre) {
			ans[i] = idx
		}
	}

	return ans
}

// lcp 09
func minJump(jump []int) int {

	vis, ans := make(map[int]bool), 0
	q := []int{0}
	lastIdx := 0
	vis[0] = true
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v >= len(jump) {
				return ans
			}
			for i := lastIdx; i < v; i++ {
				if !vis[i] {
					q = append(q, i)
					vis[i] = true
				}
			}
			if v > lastIdx {
				lastIdx = v
			}

			r := v + jump[v]
			if !vis[r] {
				q = append(q, r)
				vis[r] = true
			}
		}
		ans++
	}

	return -1
}

// lcp 10
func minimalExecTime(root *TreeNode) float64 {
	var dfs func(node *TreeNode) (max, par float64)
	dfs = func(node *TreeNode) (max, par float64) {
		if node == nil {
			return 0.0, 0.0
		}
		max = float64(node.Val)
		a, b := dfs(node.Left)
		c, d := dfs(node.Right)
		max += a + c
		if a < c {
			a, c = c, a
			b, d = d, b
		}
		if a-c <= 2*b {
			par = (a + c) / 2
		} else {
			par = b + c
		}
		return max, par
	}
	max, par := dfs(root)
	return max - par
}

// lcp 11
func expectNumber(scores []int) int {
	if len(scores) == 0 {
		return 0
	}
	maps := make(map[int]struct{}, 0)
	for _, v := range scores {
		maps[v] = struct{}{}
	}
	return len(maps)
}

// lcp 12

func minTime(time []int, m int) int {

	pre := make([]int, len(time)+1)
	for i := range time {
		pre[i+1] = pre[i] + time[i]
	}
	finish := func(t int) bool {
		now := 0
		s, e := 0, 0

		for e < len(time) {
			now++
			rangeMax := time[s]
			for e < len(time) {
				rangeMax = max(rangeMax, time[e])
				if pre[e+1]-pre[s]-rangeMax > t {
					break
				}
				e++
			}
			s = e
		}
		return now <= m
	}

	return sort.Search(1e10, func(i int) bool {
		return finish(i)
	})

	return -1
}

// lcp 14
func splitArray(nums []int) int {
	length := len(nums)

	if length == 1 {
		return 1
	}
	// 1、动态规划-记录位置i的最少切分数
	divides := make([]int, length)

	// 2、难点：记录素数prime的最佳位置， 对于[6, 7, 11, 13, 2]: 6 的素数集: 2和3
	// 则表示约数中含有2和3的数字，可以和位置0的数字6组成子数组，即：
	// primePosMap[2] = 0, primePosMap[3] = 0，primePosMap[7] = 1
	primePosMap := make(map[int]int)

	// 3、难点：通过构造素数速查表：primeMinMap[num]表示任意数字num的最小质数
	max := 0
	for i := 0; i < length; i++ {
		if nums[i] > max {
			max = nums[i]
		}
	}
	primeMinMap := make([]int, max+10)
	for i := 2; i <= max; i++ { // get新知识点：素数筛构造速查表
		if primeMinMap[i] == 0 {
			for j := i; j <= max; j += i {
				if primeMinMap[j] == 0 {
					primeMinMap[j] = i
				}
			}
		}
	}

	// 初始化位置0
	divides[0] = 1
	num := nums[0]
	for {
		// a：通过速查表，快速找到数字num的最小素数。
		prime := primeMinMap[num]
		primePosMap[prime] = 0

		// b：num/prime后得到新的数字 num2, 继续执行步骤a
		for num%prime == 0 {
			num = num / prime
		}
		if num == 1 {
			break
		}
	}

	for i := 1; i < length; i++ {
		divides[i] = i + 1 // 动态更新
		num = nums[i]
		for {
			prime := primeMinMap[num]
			if _, ok := primePosMap[prime]; !ok {
				primePosMap[prime] = i
			} else if divides[primePosMap[prime]] > divides[i-1]+1 {
				// 难点2：动态更新primePosMap，[2, 3, 5, 7, 11, 13, 17, 4, 49, 41, 51]中，
				// 一开始primePosMap[7]=3，之后由于4和49的加入，primePosMap[7]=8
				primePosMap[prime] = i
			}
			currentPos := primePosMap[prime]
			if currentPos == 0 {
				divides[i] = 1
				break
			} else if divides[currentPos-1]+1 < divides[i] {
				divides[i] = divides[currentPos-1] + 1
			}
			for num%prime == 0 {
				num = num / prime
			}
			if num <= 1 {
				break
			}
		}
	}
	//for prime, pos := range primePosMap {
	//	fmt.Printf("%d -- %d \n", prime, pos)
	//}
	//for i, s := range divides {
	//	fmt.Printf("%d --- %d --- %d\n", i, nums[i], s)
	//}
	return divides[length-1]

}

// lcp 15
// 先找一个最外围的点，再根据下一次放下选择下一个点。
//
//怎么选择下一个点？如果下次要向左转，就选当前点指向该点的向量顺时针旋转最后的那条向量对应的点；
//

func visitOrder(points [][]int, direction string) []int {
	var res []int
	// 找最外围的一点作为起点，这里找最上边的点
	var cur int
	for i, v := range points {
		if v[1] > points[cur][1] {
			cur = i
		}
	}

	// 点 a 指向点 b 的向量
	sub := func(a, b []int) []int {
		return []int{b[0] - a[0], b[1] - a[1]}
	}

	// 向量 a 和向量 b 的叉乘
	cross := func(a, b []int) int {
		return a[0]*b[1] - a[1]*b[0]
	}

	used := make([]bool, len(points))
	used[cur] = true
	res = append(res, cur)
	for _, d := range direction {
		next := -1
		for i, p := range points {
			if used[i] {
				continue
			}
			if next == -1 {
				next = i
				continue
			}
			x := sub(points[cur], points[next])
			y := sub(points[cur], p)
			cs := cross(x, y)
			if d == 'L' && cs < 0 || d == 'R' && cs > 0 {
				next = i
			}
		}
		cur = next
		used[cur] = true
		res = append(res, cur)
	}
	for i := range points {
		if !used[i] {
			res = append(res, i)
			break
		}
	}
	return res
}

// lcp 16
func maxWeight(edges [][]int, value []int) int {
	n, m := len(value), len(edges)
	cnts := make([]int, n)

	// 对边按权值和排序，以便之后对每个点，直接获得按权值和排序的边
	sort.Slice(edges, func(i int, j int) bool {
		a, b := edges[i], edges[j]
		return value[a[0]]+value[a[1]] > value[b[0]]+value[b[1]]
	})
	// 统计各个点的度数（出边数量）
	for _, v := range edges {
		cnts[v[0]]++
		cnts[v[1]]++
	}
	type pair struct {
		first, second int
	}
	// 将无向图重建为有向图
	g := make(map[int][]*pair)
	for i, v := range edges {
		if cnts[v[0]] < cnts[v[1]] || (cnts[v[0]] == cnts[v[1]] && v[0] < v[1]) {
			g[v[0]] = append(g[v[0]], &pair{
				first:  v[1],
				second: i,
			})
		} else {
			g[v[1]] = append(g[v[1]], &pair{
				first:  v[0],
				second: i,
			})
		}
	}

	nodes := make([][]int, m)
	vis, idxs := make([]int, n), make([]int, n)
	for i := range vis {
		vis[i] = -1
	}
	// 求所有的三元环，并按边归类
	for i := 0; i < m; i++ {
		for _, p := range g[edges[i][0]] {
			vis[p.first] = i
			idxs[p.first] = p.second
		}
		for _, p := range g[edges[i][1]] {
			if vis[p.first] == i {
				nodes[p.second] = append(nodes[p.second], edges[i][0])
				nodes[idxs[p.first]] = append(nodes[idxs[p.first]], edges[i][1])
				nodes[i] = append(nodes[i], p.first)
			}
		}
	}

	// 将三元环按顶点归类，每个顶点自动获得按权值和排序的边
	c := make([][]int, n)
	for i := 0; i < m; i++ {
		for _, v := range nodes[i] {
			c[v] = append(c[v], i)
		}
	}

	ans := 0
	for i := 0; i < n; i++ {
		bound := len(c[i]) - 1
		for a := 0; a < min(3, len(c[i])) && bound >= a; a++ {
			for b := a; b <= bound; b++ {
				cur := value[i] + value[edges[c[i][a]][0]] + value[edges[c[i][a]][1]]
				cnt := 0
				if edges[c[i][b]][0] != edges[c[i][a]][0] && edges[c[i][b]][0] != edges[c[i][a]][1] {
					cur += value[edges[c[i][b]][0]]
					cnt++
				}
				if edges[c[i][b]][1] != edges[c[i][a]][0] && edges[c[i][b]][1] != edges[c[i][a]][1] {
					cur += value[edges[c[i][b]][1]]
					cnt++
				}
				ans = max(ans, cur)
				if cnt == 2 {
					bound = b - 1
					break
				}
			}
		}
	}

	return ans
}
