package offer

// offer 116
func findCircleNum(isConnected [][]int) int {

	n, group := len(isConnected), len(isConnected)
	fa := make([]int, n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(from int) int {
		if fa[from] != from {
			from = find(fa[from])
		}
		return fa[from]
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
		group--
	}

	for from, v := range isConnected {
		for to, val := range v {
			if val == 1 {
				merge(from, to)
			}
		}
	}

	return group
}
