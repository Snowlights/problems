package _50

import (
	"bufio"
	"fmt"
	"io"
)

func CF377A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, m, k int
	fmt.Fscan(r, &n, &m, &k)
	arr := make([][]byte, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &arr[i])
	}
	sx, sy, s := -1, -1, 0
	for i, v := range arr {
		for j, vv := range v {
			if vv == '.' {
				if sx == -1 {
					sx, sy = i, j
				}
				s++
			}
		}
	}
	vis := make([][]bool, n)
	for i := range vis {
		vis[i] = make([]bool, m)
	}

	k = s - k
	q := [][]int{{sx, sy}}
	vis[sx][sy] = true
	k--
	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	for len(q) > 0 && k > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			for _, d := range dir {
				x, y := v[0]+d[0], v[1]+d[1]
				if 0 <= x && x < n && 0 <= y && y < m && arr[x][y] == '.' && !vis[x][y] && k > 0 {
					q = append(q, []int{x, y})
					vis[x][y] = true
					k--
				}
			}
		}
	}

	for i, v := range arr {
		for j := range v {
			if arr[i][j] == '.' && !vis[i][j] {
				arr[i][j] = 'X'
			}
		}
		fmt.Fprintln(w, string(arr[i]))
	}
}
