package _900_2000

// 1926
func nearestExit(maze [][]byte, entrance []int) int {
	m, n := len(maze), len(maze[0])
	out, vis := make([][]bool, m), make([][]bool, m)
	for i := range out {
		out[i] = make([]bool, n)
		vis[i] = make([]bool, n)
	}
	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	for i := 0; i < m; i++ {
		if maze[i][0] == '.' {
			out[i][0] = true
		}
		if maze[i][n-1] == '.' {
			out[i][n-1] = true
		}
	}
	for i := 0; i < n; i++ {
		if maze[0][i] == '.' {
			out[0][i] = true
		}
		if maze[m-1][i] == '.' {
			out[m-1][i] = true
		}
	}

	q, l := [][]int{entrance}, 0
	vis[entrance[0]][entrance[1]] = true
	out[entrance[0]][entrance[1]] = false
	for len(q) > 0 {
		t := q
		q = nil
		for _, v := range t {
			x, y := v[0], v[1]
			if out[x][y] {
				return l
			}
			for _, d := range dir {
				xx, yy := x+d[0], y+d[1]
				if 0 <= xx && xx < m && 0 <= yy && yy < n && maze[xx][yy] == '.' && !vis[xx][yy] {
					q = append(q, []int{xx, yy})
					vis[xx][yy] = true
				}
			}
		}
		l++
	}

	return -1
}
