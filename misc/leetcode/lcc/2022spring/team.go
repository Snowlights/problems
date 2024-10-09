package _022spring

import (
	"container/heap"
	. "problems/testutil/leetcode"
	"sort"
)

// 1
func getMinimumTime(time []int, fruits [][]int, limit int) int {
	var res int
	for _, fruit := range fruits {
		i := fruit[0]
		num := fruit[1]
		res += (num + limit - 1) / limit * time[i]
	}
	return res
}

// 2
func conveyorBelt(matrix []string, start []int, end []int) int {
	n := len(matrix)
	m := len(matrix[0])
	total := n * m
	dst := make([][]int, n)
	for i := range dst {
		dst[i] = make([]int, m)
		for j := range dst[i] {
			dst[i][j] = total
		}
	}
	type node struct{ x, y int }
	var dirs4 = []node{
		{-1, 0},
		{0, 1},
		{1, 0},
		{0, -1}}
	arrowCorrect := "v<^>"
	endNode := &node{end[0], end[1]}
	queue := []*node{endNode}
	dst[endNode.x][endNode.y] = 0
	for len(queue) > 0 {
		uNode := queue[0]
		queue = queue[1:]
		ux := uNode.x
		uy := uNode.y
		for i := 0; i < 4; i++ {
			nx := ux + dirs4[i].x
			ny := uy + dirs4[i].y
			if nx >= 0 && ny >= 0 && nx < n && ny < m {
				if matrix[nx][ny] == arrowCorrect[i] {
					if dst[nx][ny] > dst[ux][uy] {
						dst[nx][ny] = dst[ux][uy]
						queue = append(queue, &node{nx, ny})
					}
				} else {
					if dst[nx][ny] > dst[ux][uy]+1 {
						dst[nx][ny] = dst[ux][uy] + 1
						queue = append(queue, &node{nx, ny})
					}
				}
			}
		}
	}
	return dst[start[0]][start[1]]

}

// 3
func getMaximumNumber(moles [][]int) int {
	sort.Slice(moles, func(i, j int) bool {
		return moles[i][0] < moles[j][0]
	})
	type point struct{ x, y int }
	var molesByTime [][]*point
	var timeSort []int
	for _, mole := range moles {
		if len(timeSort) == 0 || mole[0] > timeSort[len(timeSort)-1] {
			timeSort = append(timeSort, mole[0])
			molesByTime = append(molesByTime, []*point{})
		}
		molesByTime[len(molesByTime)-1] = append(molesByTime[len(molesByTime)-1], &point{mole[1], mole[2]})
	}
	var preTime int
	var curTime int
	dp := [3][3]int{{-3, -3, -3}, {-3, 0, -3}, {-3, -3, -3}}

	if timeSort[0] == 0 {
		for _, cur := range molesByTime[0] {
			if cur.x == 1 && cur.y == 1 {
				dp[1][1] = 1
				break
			}
		}
		timeSort = timeSort[1:]
		molesByTime = molesByTime[1:]
	}

	for i, points := range molesByTime {
		curTime = timeSort[i]
		for j := 0; j < 4 && preTime+1+j <= curTime; j++ {
			var dpNew [3][3]int
			dpNew[0][0] = max(dp[0][0], max(dp[0][1], dp[1][0]))
			dpNew[0][1] = max(dp[0][1], max(dp[0][0], max(dp[0][2], dp[1][1])))
			dpNew[0][2] = max(dp[0][2], max(dp[0][1], dp[1][2]))
			dpNew[1][0] = max(dp[1][0], max(dp[0][0], max(dp[2][0], dp[1][1])))
			dpNew[1][1] = max(dp[1][1], max(dp[0][1], max(dp[1][0], max(dp[1][2], dp[2][1]))))
			dpNew[1][2] = max(dp[1][2], max(dp[0][2], max(dp[1][1], dp[2][2])))
			dpNew[2][0] = max(dp[2][0], max(dp[1][0], dp[2][1]))
			dpNew[2][1] = max(dp[2][1], max(dp[2][0], max(dp[2][2], dp[1][1])))
			dpNew[2][2] = max(dp[2][2], max(dp[1][2], dp[2][1]))
			dp = dpNew
		}
		for _, p := range points {
			dp[p.x][p.y]++
		}
		preTime = curTime
	}
	var res int
	for r := 0; r < 3; r++ {
		for c := 0; c < 3; c++ {
			res = max(res, dp[r][c])
		}
	}
	return res
}

// 4
func composeCube(shapes [][]string) bool {
	n := len(shapes[0][0])
	cube := make([][][]int, n)
	for i := range cube {
		cube[i] = make([][]int, n)
		for j := range cube[i] {
			cube[i][j] = make([]int, n)
		}
	}
	total := n*n*n - (n-2)*(n-2)*(n-2)
	pieces := make([][][]int, 6)
	for i := range pieces {
		pieces[i] = make([][]int, n)
		for j := range pieces[i] {
			pieces[i][j] = make([]int, n)
		}
	}
	for i, shape := range shapes {
		for r, row := range shape {
			for c, ch := range row {
				pieces[i][r][c] = int(ch - '0')
				if pieces[i][r][c] == 1 {
					total--
				}
			}
		}
	}
	if total != 0 {
		return false
	}
	var vis [6]bool
	vis[0] = true
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			cube[0][i][j] = pieces[0][i][j]
		}
	}
	var dfs func(curCube [][][]int, planeIdx int) bool
	dfs = func(curCube [][][]int, planeIdx int) bool {
		if planeIdx == 6 {
			return true
		}
		for i := 0; i < 6; i++ {
			if vis[i] == false {
				vis[i] = true
				for dIdx := 0; dIdx < 8; dIdx++ {
					cubeCopy := makeCubeCopy(curCube, n)
					if composePiece(cubeCopy, pieces[i], planeIdx, dIdx) && dfs(cubeCopy, planeIdx+1) {
						return true
					}
				}
				vis[i] = false
			}
		}
		return false
	}
	return dfs(cube, 1)
}

func convertPieceToCube(piece [][]int, plantIdx, dIdx int) [][][]int {
	n := len(piece)
	pieceCopy := make([][]int, n)
	for i := range pieceCopy {
		pieceCopy[i] = make([]int, n)
		copy(pieceCopy[i], piece[i])
	}
	var res [][][]int
	switch dIdx % 4 {
	case 1:
		for i := 0; i < n/2; i++ {
			pieceCopy[i], pieceCopy[n-1-i] = pieceCopy[n-1-i], pieceCopy[i]
		}
		for i := 0; i < n; i++ {
			for j := 0; j < n/2; j++ {
				pieceCopy[i][j], pieceCopy[i][n-1-j] = pieceCopy[i][n-1-j], pieceCopy[i][j]
			}
		}
	case 2:
		for i := 0; i < n/2; i++ {
			pieceCopy[i], pieceCopy[n-1-i] = pieceCopy[n-1-i], pieceCopy[i]
		}
		for i := 0; i < n; i++ {
			for j := 0; j < i; j++ {
				pieceCopy[i][j], pieceCopy[j][i] = pieceCopy[j][i], pieceCopy[i][j]
			}
		}
	case 3:
		for i := 0; i < n/2; i++ {
			pieceCopy[i], pieceCopy[n-1-i] = pieceCopy[n-1-i], pieceCopy[i]
		}
		for i := 0; i < n; i++ {
			for j := 0; n-1-j > i; j++ {
				pieceCopy[i][j], pieceCopy[n-1-j][n-1-i] = pieceCopy[n-1-j][n-1-i], pieceCopy[i][j]
			}
		}
	}
	if dIdx >= 4 {
		for i := 0; i < n; i++ {
			for j := 0; j < n/2; j++ {
				pieceCopy[i][j], pieceCopy[i][n-1-j] = pieceCopy[i][n-1-j], pieceCopy[i][j]
			}
		}
	}
	switch plantIdx {
	case 1:
		res = make([][][]int, n)
		for i := range res {
			res[i] = make([][]int, n)
			for j := range res[i] {
				res[i][j] = make([]int, n)
				copy(res[i][j], pieceCopy[j])
			}
			copy(res[i][0], pieceCopy[i])
		}
		for l := range res {
			copy(res[l][n-1], pieceCopy[l])
		}

	case 2:
		res = make([][][]int, n)
		for i := range res {
			res[i] = make([][]int, n)
			for j := range res[i] {
				res[i][j] = make([]int, n)
				res[i][j][0] = pieceCopy[i][j]
			}
		}
		for l := range res {
			for x := range res[l] {
				res[l][x][n-1] = pieceCopy[l][x]
			}
		}
	case 3:
		res = make([][][]int, n)
		for i := range res {
			res[i] = make([][]int, n)
			for j := range res[i] {
				res[i][j] = make([]int, n)
			}
		}
		for l := range res {
			copy(res[l][0], pieceCopy[l])
		}
	case 4:
		res = make([][][]int, n)
		for i := range res {
			res[i] = make([][]int, n)
			for j := range res[i] {
				res[i][j] = make([]int, n)
				res[i][j][0] = pieceCopy[i][j]
			}
		}
		for l := range res {
			for x := range res[l] {
				res[l][x][0] = pieceCopy[l][x]
			}
		}
	case 5:
		res = make([][][]int, n)
		for i := range res {
			res[i] = make([][]int, n)
			for j := range res[i] {
				res[i][j] = make([]int, n)
			}
		}
		for x := range res[n-1] {
			copy(res[n-1][x], pieceCopy[x])
		}
	}
	return res
}

func composePiece(curCube [][][]int, piece [][]int, plantIdx, dIdx int) bool {
	n := len(curCube)
	extCube := convertPieceToCube(piece, plantIdx, dIdx)
	switch plantIdx {
	case 1:
		x := n - 1
		for l := 0; l < n; l++ {
			for y := 0; y < n; y++ {
				if curCube[l][x][y] == 1 && extCube[l][x][y] == 1 {
					return false
				} else {
					curCube[l][x][y] |= extCube[l][x][y]
				}
			}
		}
	case 2:
		y := n - 1
		for l := 0; l < n; l++ {
			for x := 0; x < n; x++ {
				if curCube[l][x][y] == 1 && extCube[l][x][y] == 1 {
					return false
				} else {
					curCube[l][x][y] |= extCube[l][x][y]
				}
			}
		}
	case 3:
		x := 0
		for l := 0; l < n; l++ {
			for y := 0; y < n; y++ {
				if curCube[l][x][y] == 1 && extCube[l][x][y] == 1 {
					return false
				} else {
					curCube[l][x][y] |= extCube[l][x][y]
				}
			}
		}
	case 4:
		y := 0
		for l := 0; l < n; l++ {
			for x := 0; x < n; x++ {
				if curCube[l][x][y] == 1 && extCube[l][x][y] == 1 {
					return false
				} else {
					curCube[l][x][y] |= extCube[l][x][y]
				}
			}
		}
	case 5:
		l := n - 1
		for x := 0; x < n; x++ {
			for y := 0; y < n; y++ {
				if curCube[l][x][y] == 1 && extCube[l][x][y] == 1 {
					return false
				} else {
					curCube[l][x][y] |= extCube[l][x][y]
				}
			}
		}
	}
	return true
}

func makeCubeCopy(curCube [][][]int, n int) [][][]int {
	cubeCopy := make([][][]int, n)
	for i := range cubeCopy {
		cubeCopy[i] = make([][]int, n)
		for j := range cubeCopy[i] {
			cubeCopy[i][j] = make([]int, n)
			copy(cubeCopy[i][j], curCube[i][j])
		}
	}
	return cubeCopy
}

// 5
func buildBridge(num int, wood [][]int) int64 {
	n := len(wood)
	L, R := &hp2{}, &hp{}
	L.push(wood[0][0])
	R.push(wood[0][0])
	biasL, biasR, ans := 0, 0, 0
	for i := 1; i < n; i++ {
		biasL -= wood[i][1] - wood[i][0]
		biasR += wood[i-1][1] - wood[i-1][0]
		l0 := L.Top().(int) + biasL
		r0 := R.Top().(int) + biasR
		x := wood[i][0]
		if x < l0 {
			ans += l0 - x
			L.pop()
			L.push(x - biasL)
			L.push(x - biasL)
			R.push(l0 - biasR)
		} else if x > r0 {
			ans += x - r0
			R.pop()
			R.push(x - biasR)
			R.push(x - biasR)
			L.push(r0 - biasL)
		} else {
			L.push(x - biasL)
			R.push(x - biasR)
		}
	}

	return int64(ans)
}

type hp struct{ sort.IntSlice }

// func (h hp) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] } // 加上这行变成最大堆
func (h *hp) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *hp) Top() interface{} {
	a := h.IntSlice
	v := a[0]
	return v
}
func (h *hp) push(v int) { heap.Push(h, v) }
func (h *hp) pop() int   { return heap.Pop(h).(int) } // 稍微封装一下，方便使用

type hp2 struct{ sort.IntSlice }

func (h hp2) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp2) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp2) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}
func (h *hp2) Top() interface{} {
	a := h.IntSlice
	v := a[0]
	return v
}
func (h *hp2) push(v int) { heap.Push(h, v) }
func (h *hp2) pop() int   { return heap.Pop(h).(int) } // 稍微封装一下，方便使用

// 6
type pair struct {
	first, second int
}

type Result struct {
	mp      map[int]*pair
	pending map[int]struct{}
}

func (res *Result) merge(r *Result) {
	for k, v := range r.mp {
		res.mp[k].first += v.first
		res.mp[k].second += v.second
		res.pending[k] = struct{}{}
	}
	for v := range r.pending {
		res.pending[v] = struct{}{}
	}
	return
}

func getMaxLayerSum(root *TreeNode) int {

	sm, sz := make([]int, 0, 1e6), make([]int, 0, 1e6)
	ans := 0

	var dfs func(node *TreeNode, d int)
	dfs = func(node *TreeNode, d int) {
		if node == nil {
			return
		}
		if len(sm) <= d {
			sm = append(sm, 0)
			sz = append(sz, 0)
		}
		sm[d] += node.Val
		sz[d]++
		dfs(node.Left, d+1)
		dfs(node.Right, d+1)
	}

	var dfs2 func(node *TreeNode, d int) *Result
	dfs2 = func(node *TreeNode, d int) *Result {
		if node.Left == nil && node.Right == nil {
			ret := &Result{
				mp:      make(map[int]*pair),
				pending: make(map[int]struct{}),
			}
			if sz[d] > 1 {
				ans = max(ans, sm[d]-node.Val)
				ret.mp[d] = &pair{
					first:  -node.Val,
					second: -1,
				}
			}
			return ret
		} else if node.Left != nil && node.Right != nil {
			a, b := dfs2(node.Left, d+1), dfs2(node.Right, d+1)
			if len(a.mp) < len(b.mp) {
				a, b = b, a
			}
			a.merge(b)
			a.mp[d] = &pair{
				first:  node.Left.Val + node.Right.Val - node.Val,
				second: 1,
			}
			a.pending[d] = struct{}{}
			return a
		} else {
			ch := node.Left
			if ch == nil {
				ch = node.Right
			}
			ret := dfs2(ch, d+1)
			for k := range ret.pending {
				p := ret.mp[k]
				if sz[k]+p.second > 0 {
					ans = max(ans, sm[k]+p.first)
				}
			}
			ret.pending = make(map[int]struct{})

			del := ch.Val - node.Val
			ans = max(ans, sm[d]+del)
			ret.mp[d] = &pair{del, 0}
			return ret
		}
	}

	dfs(root, 0)
	ans = sm[0]
	for _, v := range sm {
		ans = max(ans, v)
	}
	dfs2(root, 0)
	return ans
}

// time limit version
//func getMaxLayerSum(root *TreeNode) int {
//	var (
//		n, cnt, ans int
//		down        [2e6 + 10]int
//		ch          [2][2e6 + 10]int
//		de, mx      [2e6 + 10]int
//		mp          [2e6 + 10]map[int]int
//	)
//	var dfs func(x *TreeNode, dep int) int
//	dfs = func(x *TreeNode, dep int) int {
//		if x == nil {
//			return 0
//		}
//		n = max(n, dep)
//		cnt++
//		o := cnt
//		l, r := dfs(x.Left, dep+1), dfs(x.Right, dep+1)
//		ch[0][o], ch[1][o] = l, r
//		down[o] = max(down[l], down[r])
//		down[o] = max(down[o], dep)
//		return o
//	}
//
//	var dfs2 func(x *TreeNode, dep, length, total int) int
//	dfs2 = func(x *TreeNode, dep, length, total int) int {
//		if x == nil {
//			return 0
//		}
//		de[dep] += x.Val
//		t := -x.Val
//		if x.Left != nil {
//			t += x.Left.Val
//		}
//		if x.Right != nil {
//			t += x.Right.Val
//		}
//		cnt++
//		out := cnt
//		l, r := 0, 0
//		if ch[0][out] > 0 && ch[1][out] > 0 {
//
//		} else {
//			length = 0
//		}
//		l = dfs2(x.Left, dep+1, max(length, down[ch[1][out]]), max(total, down[ch[1][out]]))
//		r = dfs2(x.Right, dep+1, max(length, down[ch[0][out]]), max(total, down[ch[0][out]]))
//		if len(mp[l]) > len(mp[r]) {
//			l, r = r, l
//		}
//		for k := range mp[l] {
//			val, ok := mp[r][k]
//			if ok {
//				mp[l][k] += val
//				if k > length && (k <= total || k != n) {
//					mx[k] = max(mx[k], mp[l][k])
//				}
//				delete(mp[r], k)
//			}
//		}
//		if r > 0 {
//			out = r
//		} else {
//			mp[out] = make(map[int]int)
//		}
//		for k, v := range mp[l] {
//			mp[out][k] = v
//		}
//		mp[out][dep] = t
//		if dep > length && (dep <= total || dep != n) {
//			mx[dep] = max(mx[dep], t)
//		}
//		return out
//	}
//
//	dfs(root, 1)
//	cnt = 0
//	for i := 0; i <= n; i++ {
//		de[i], mx[i] = 0, 0
//	}
//	dfs2(root, 1, math.MaxInt32, 0)
//	ans = math.MinInt32
//	for i := 1; i <= n; i++ {
//		ans = max(ans, de[i]+mx[i])
//	}
//	return ans
//}
