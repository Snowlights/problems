package main

import (
	"maps"
	"slices"
	"sort"
)

type fenwick []int

func (f fenwick) add(i int) {
	for ; i < len(f); i += i & -i {
		f[i]++
	}
}

func (f fenwick) pre(i int) (res int) {
	for ; i > 0; i &= i - 1 {
		res += f[i]
	}
	return
}

func (f fenwick) query(l, r int) int {
	return f.pre(r) - f.pre(l-1)
}

func maxRectangleArea(points [][]int) int {
	xMap := map[int][]int{} // 同一列的所有点的纵坐标
	yMap := map[int][]int{} // 同一行的所有点的横坐标
	for _, p := range points {
		x, y := p[0], p[1]
		xMap[x] = append(xMap[x], y)
		yMap[y] = append(yMap[y], x)
	}

	// 预处理每个点的正下方的点
	type pair struct{ x, y int }
	below := map[pair]int{}
	for x, ys := range xMap {
		slices.Sort(ys)
		for i := 1; i < len(ys); i++ {
			below[pair{x, ys[i]}] = ys[i-1]
		}
	}

	// 预处理每个点的正左边的点
	left := map[pair]int{}
	for y, xs := range yMap {
		slices.Sort(xs)
		for i := 1; i < len(xs); i++ {
			left[pair{xs[i], y}] = xs[i-1]
		}
	}

	// 离散化用
	xs := slices.Sorted(maps.Keys(xMap))
	ys := slices.Sorted(maps.Keys(yMap))

	// 收集询问：矩形区域（包括边界）的点的个数
	type query struct {
		x1, x2, y1, y2 int
		area           int64
	}
	queries := []query{}
	// 枚举 (x2,y2) 作为矩形的右上角
	for x2, listY := range xMap {
		for i := 1; i < len(listY); i++ {
			// 计算矩形左下角 (x1,y1)
			y2 := listY[i]
			x1, ok := left[pair{x2, y2}]
			if !ok {
				continue
			}
			y1 := listY[i-1] // (x2,y2) 下面的点（矩形右下角）的纵坐标
			// 矩形右下角的左边的点的横坐标必须是 x1
			if x, ok := left[pair{x2, y1}]; !ok || x != x1 {
				continue
			}
			// 矩形左上角的下边的点的纵坐标必须是 y1
			if y, ok := below[pair{x1, y2}]; !ok || y != y1 {
				continue
			}
			queries = append(queries, query{
				sort.SearchInts(xs, x1), // 离散化
				sort.SearchInts(xs, x2),
				sort.SearchInts(ys, y1),
				sort.SearchInts(ys, y2),
				int64(x2-x1) * int64(y2-y1),
			})
		}
	}

	// 离线询问
	type data struct{ qid, sign, y1, y2 int }
	qs := make([][]data, len(xs))
	for i, q := range queries {
		if q.x1 > 0 {
			qs[q.x1-1] = append(qs[q.x1-1], data{i, -1, q.y1, q.y2})
		}
		qs[q.x2] = append(qs[q.x2], data{i, 1, q.y1, q.y2})
	}

	// 回答询问
	res := make([]int, len(queries))
	tree := make(fenwick, len(ys)+1)
	for i, x := range xs {
		// 把横坐标为 x 的所有点都加到树状数组中
		for _, y := range xMap[x] {
			tree.add(sort.SearchInts(ys, y) + 1) // 离散化
		}
		for _, q := range qs[i] {
			// 查询 [y1,y2] 中点的个数
			res[q.qid] += q.sign * tree.query(q.y1+1, q.y2+1)
		}
	}

	ans := int64(-1)
	for i, cnt := range res {
		if cnt == 4 { // 矩形区域（包括边界）恰好有 4 个点
			ans = max(ans, queries[i].area)
		}
	}
	return int(ans)
}

func maxRectangleArea2(points [][]int) int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0] < points[j][0] ||
			(points[i][0] == points[j][0] && points[i][1] < points[j][1])
	})
	has := map[[2]int]int{}
	for _, p := range points {
		has[[2]int{p[0], p[1]}]++
	}
	n := len(points)
	ans := -1
	//fmt.Println(points, has)
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			if points[j][0] == points[i][0] ||
				points[j][1] >= points[i][1] {
				//fmt.Println(j, i)
				continue
			}
			l, r := points[j][0], points[i][0]
			lo, hi := points[j][1], points[i][1]
			//fmt.Println(l, r)
			if has[[2]int{l, hi}] != 1 || has[[2]int{r, lo}] != 1 {
				continue
			}
			cnt := 0
			for c := 0; c < n; c++ {
				if points[c][0] >= l && points[c][0] <= r &&
					points[c][1] >= lo && points[c][1] <= hi {
					cnt++
				}
			}
			if cnt == 4 {
				ans = max(ans, (r-l)*(hi-lo))
			}
		}
	}
	return ans
}
