package main

func minimumOperationsToMakeEqual(x, y int) int {
	if x <= y {
		return y - x
	}

	q, vis := []int{x}, make(map[int]bool)
	vis[x] = true
	for step := 0; ; step++ {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v == y {
				return step
			}
			if v%11 == 0 && !vis[v/11] {
				q = append(q, v/11)
				vis[v/11] = true
			}
			if v%5 == 0 && !vis[v/5] {
				q = append(q, v/5)
				vis[v/5] = true
			}
			if !vis[v+1] {
				q = append(q, v+1)
				vis[v+1] = true
			}
			if !vis[v-1] {
				q = append(q, v-1)
				vis[v-1] = true
			}
		}
	}
}
