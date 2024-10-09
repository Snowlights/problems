package main

import "sort"

func minimumTime(a []int, b []int, x int) (ans int) {

	type pair struct {
		x, y int
	}
	pList := make([]*pair, 0, len(a))
	suma, sumb := 0, 0
	for i := range a {
		pList = append(pList, &pair{
			x: a[i],
			y: b[i],
		})
		suma += a[i]
		sumb += b[i]
	}
	sort.Slice(pList, func(i, j int) bool {
		return pList[i].y < pList[j].y ||
			pList[i].y == pList[j].y && pList[i].x < pList[j].x
	})

	f := make([]int, len(a)+1)
	for _, p := range pList {
		for i := len(a) - 1; i >= 0; i-- {
			f[i+1] = max(f[i+1], f[i]+(i+1)*p.y+p.x)
		}
	}

	for i := 0; i <= len(a); i++ {
		if sumb*i+suma-f[i] <= x {
			return i
		}
	}

	return -1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
