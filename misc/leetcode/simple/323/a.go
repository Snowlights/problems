package simple_323

import "sort"

// 1
func deleteGreatestValue(grid [][]int) int {
	for i, v := range grid {
		sort.Ints(v)
		grid[i] = v
	}
	ans := 0
	n, m := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		tmp := grid[0][i]
		for j := 0; j < n; j++ {
			tmp = max(tmp, grid[j][i])
		}
		ans += tmp
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 2
func longestSquareStreak(nums []int) int {
	sort.Ints(nums)
	h := make(map[int]int)
	ans := -1
	for i := len(nums) - 1; i >= 0; i-- {
		val, ok := h[nums[i]*nums[i]]
		if ok {
			h[nums[i]] = val + 1
			ans = max(ans, val+1)
		} else {
			h[nums[i]] = 1
		}
	}

	return ans
}

// 3
type Allocator struct {
	mem []int
}

func Constructor(n int) Allocator {
	return Allocator{make([]int, n)}
}

func (this *Allocator) Allocate(size int, mID int) int {
	cnt, idx := 0, 0
	for i, v := range this.mem {
		if v > 0 {
			cnt = 0
		} else {
			cnt++
			if cnt == size {
				idx = i
				break
			}
		}
	}
	tmp := cnt
	if tmp != size {
		return -1
	}
	for ; idx >= 0 && tmp > 0; idx-- {
		this.mem[idx] = mID
		tmp--
	}
	if cnt == size {
		return idx + 1
	}
	return -1
}

func (this *Allocator) Free(mID int) int {
	cnt := 0
	for i, v := range this.mem {
		if v == mID {
			this.mem[i] = 0
			cnt++
		}
	}
	return cnt
}

// 4
func maxPoints(grid [][]int, queries []int) []int {
	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	type tuple struct {
		x, i, j int
	}
	n, m := len(grid), len(grid[0])
	tupleList := make([]tuple, 0)
	for i, v := range grid {
		for j, vv := range v {
			tupleList = append(tupleList, tuple{
				x: vv,
				i: i,
				j: j,
			})
		}
	}
	sort.Slice(tupleList, func(i, j int) bool {
		return tupleList[i].x < tupleList[j].x
	})

	ids := make([]int, 0)
	for i := range queries {
		ids = append(ids, i)
	}
	sort.Slice(ids, func(i, j int) bool {
		return queries[ids[i]] < queries[ids[j]]
	})

	fa, size := make([]int, n*m), make([]int, n*m)
	for i := range fa {
		fa[i] = i
		size[i] = 1
	}
	var find func(int) int
	find = func(i int) int {
		if fa[i] != i {
			fa[i] = find(fa[i])
		}
		return fa[i]
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
		size[to] += size[from]
	}
	j, ans := 0, make([]int, len(queries))
	for _, id := range ids {
		val := queries[id]
		for ; j < len(tupleList) && tupleList[j].x < val; j++ {
			tuple := tupleList[j]
			for _, d := range dir {
				x, y := tuple.i+d[0], tuple.j+d[1]
				if 0 <= x && x < n && 0 <= y && y < m && grid[x][y] < val {
					merge(tuple.i*m+tuple.j, x*m+y)
				}
			}
		}
		if grid[0][0] < val {
			ans[id] = size[find(0)]
		}
	}

	return ans
}

var dirs = []struct{ x, y int }{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
