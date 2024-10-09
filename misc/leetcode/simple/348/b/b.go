package main

func semiOrderedPermutation(a []int) (ans int) {
	n, first, last := len(a), 0, 0
	for i, v := range a {
		if v == 1 {
			first = i
		}
		if v == n {
			last = i
		}
	}

	if first < last {
		return first + n - 1 - last
	}

	return first + n - 2 - last
}
