package simple_308

import (
	"sort"
	"strings"
)

// 1

func answerQueries(nums []int, queries []int) []int {
	sort.Ints(nums)
	ans := make([]int, 0, len(queries))
	for _, v := range queries {
		tmp, c := 0, 0
		for _, vv := range nums {
			if tmp+vv <= v {
				tmp += vv
				c++
			} else {
				break
			}
		}
		ans = append(ans, c)
	}
	return ans
}

// 2
func removeStars(s string) string {

	ans := make([]byte, 0)
	for _, v := range s {
		if v == '*' {
			ans = ans[:len(ans)-1]
			continue
		}
		ans = append(ans, byte(v))
	}

	return string(ans)
}

// 3
func garbageCollection(garbage []string, travel []int) int {

	pl, gl, ml := 0, 0, 0
	ans := 0
	for i, v := range garbage {
		m := strings.Count(v, "M")
		if m > 0 {
			ml = i
			ans += m
		}
		p := strings.Count(v, "P")
		if p > 0 {
			pl = i
			ans += p
		}
		g := strings.Count(v, "G")
		if g > 0 {
			gl = i
			ans += g
		}
	}
	for _, v := range travel[:gl] {
		ans += v
	}
	for _, v := range travel[:ml] {
		ans += v
	}
	for _, v := range travel[:pl] {
		ans += v
	}

	return ans
}

// 4
func topoSort(k int, edges [][]int) []int {
	g, in := make(map[int][]int), make(map[int]int)
	for _, v := range edges {
		g[v[0]] = append(g[v[0]], v[1])
		in[v[1]]++
	}
	q, ans := make([]int, 0, k), make([]int, 0)
	for i := 1; i <= k; i++ {
		if in[i] == 0 {
			q = append(q, i)
			ans = append(ans, i)
		}
	}
	for len(q) > 0 {
		v := q[0]
		q = q[1:]
		for _, v := range g[v] {
			if in[v]--; in[v] == 0 {
				q = append(q, v)
				ans = append(ans, v)
			}
		}
	}

	if len(q) > 0 || len(ans) != k {
		return nil
	}

	return ans
}

func buildMatrix(k int, rowConditions, colConditions [][]int) [][]int {
	row := topoSort(k, rowConditions)
	col := topoSort(k, colConditions)
	// fmt.Println(row, col)
	if row == nil || col == nil {
		return nil
	}
	pos := make([]int, k+1)
	for i, v := range col {
		pos[v] = i
	}
	ans := make([][]int, k)
	for i := range ans {
		ans[i] = make([]int, k)
	}
	for i, v := range row {
		ans[i][pos[v]] = v
	}

	return ans
}
