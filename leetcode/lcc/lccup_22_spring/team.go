package lccup_22_spring

import (
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
// todo 第五题 https://leetcode.cn/contest/season/2022-spring/

// 6
// todo 第六题
