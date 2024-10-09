package c

func extractMantra(matrix []string, mantra string) int {
	m, n := len(matrix), len(matrix[0])
	type pair struct {
		i, j, k int
	}
	q := []pair{pair{
		i: 0,
		j: 0,
		k: 0,
	}}
	vis := make(map[pair]bool)
	vis[q[0]] = true
	step := 1
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, p := range tmp {
			i, j, k := p.i, p.j, p.k
			if matrix[i][j] == mantra[k] {
				if k == len(mantra)-1 {
					return step
				}
				newP := pair{
					i: i,
					j: j,
					k: k + 1,
				}
				if !vis[newP] {
					vis[newP] = true
					q = append(q, newP)
				}
			}
			for _, val := range [][]int{{i + 1, j}, {i - 1, j}, {i, j + 1}, {i, j - 1}} {
				x, y := val[0], val[1]
				if 0 <= x && x < m && 0 <= y && y < n {
					newP := pair{
						i: x,
						j: y,
						k: k,
					}
					if !vis[newP] {
						vis[newP] = true
						q = append(q, newP)
					}
				}
			}
		}
		step++
	}

	return -1
}
