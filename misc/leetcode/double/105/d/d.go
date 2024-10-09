package main

const mx = 1e5

var lpf [mx + 1]int

func init() {
	for i := 2; i <= mx; i++ {
		if lpf[i] == 0 {
			for j := i; j <= mx; j += i {
				if lpf[j] == 0 {
					lpf[j] = i
				}
			}
		}
	}
}

func canTraverseAllPairs(a []int) (ans bool) {
	n := len(a)
	fa := make([]int, n+mx)
	for i := range fa {
		fa[i] = i
	}
	var find func(from int) int
	find = func(from int) int {
		if fa[from] != from {
			fa[from] = find(fa[from])
		}
		return fa[from]
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
	}

	for i, v := range a {
		for v > 1 {
			p := lpf[v]
			for ; lpf[v] == p; v /= p {
			}
			merge(i, n+p)
		}
	}

	for i := 1; i < n; i++ {
		if find(i) != find(0) {
			return false
		}
	}

	return true
}
