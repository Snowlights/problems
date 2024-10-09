package simple_312

import "sort"

// 1
func sortPeople(names []string, heights []int) []string {
	type pair struct {
		name   string
		height int
	}
	pairList := make([]*pair, 0)
	for i, name := range names {
		pairList = append(pairList, &pair{name: name, height: heights[i]})
	}
	sort.Slice(pairList, func(i, j int) bool {
		return pairList[i].height > pairList[j].height
	})
	ans := make([]string, 0)
	for _, v := range pairList {
		ans = append(ans, v.name)
	}
	return ans
}

// 2
func longestSubarray(nums []int) int {
	tmp := nums[0]
	for i := 1; i < len(nums); i++ {
		tmp = max(tmp, nums[i])
	}
	ans, val := 0, 0
	for _, v := range nums {
		if v == tmp {
			val++
			ans = max(ans, val)
		} else {
			val = 0
		}
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 3
func goodIndices(nums []int, k int) []int {
	l, r := make([]int, len(nums)), make([]int, len(nums))

	for i := 1; i < len(nums); i++ {
		if nums[i] <= nums[i-1] {
			l[i] = l[i-1] + 1
		}
	}
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] <= nums[i+1] {
			r[i] = r[i+1] + 1
		}
	}

	ans := make([]int, 0)
	for i := 1; i < len(nums)-1; i++ {
		if l[i-1] >= k-1 && r[i+1] >= k-1 {
			ans = append(ans, i)
		}
	}
	return ans
}

// 4

func numberOfGoodPaths(vals []int, edges [][]int) int {
	n := len(vals)
	g := make([][]int, n)
	for _, e := range edges {
		x, y := e[0], e[1]
		g[x] = append(g[x], y)
		g[y] = append(g[y], x) // 建图
	}

	// 并查集模板
	fa := make([]int, n)
	// size[x] 表示节点值等于 vals[x] 的节点个数，如果按照节点值从小到大合并，
	// size[x] 也是连通块内的等于最大节点值的节点个数
	size := make([]int, n)
	id := make([]int, n) // 后面排序用
	for i := range fa {
		fa[i] = i
		size[i] = 1
		id[i] = i
	}
	var find func(int) int
	find = func(x int) int {
		if fa[x] != x {
			fa[x] = find(fa[x])
		}
		return fa[x]
	}

	ans := n
	sort.Slice(id, func(i, j int) bool { return vals[id[i]] < vals[id[j]] })
	for _, x := range id {
		vx := vals[x]
		fx := find(x)
		for _, y := range g[x] {
			y = find(y)
			if y == fx || vals[y] > vx { // 只考虑最大节点值比 vx 小的连通块
				continue
			}
			if vals[y] == vx { // 可以构成好路径
				ans += size[fx] * size[y] // 乘法原理
				size[fx] += size[y]       // 统计连通块内节点值等于 vx 的节点个数
			}
			fa[y] = fx // 把小的节点值合并到大的节点值上
		}
	}
	return ans
}
