package main

func goodSubsetofBinaryMatrix(g [][]int) (ans []int) {

	idx := make(map[int]int)
	for i, v := range g {
		mask := 0
		for j, vv := range v {
			mask |= vv << j
		}
		idx[mask] = i
	}

	if val, ok := idx[0]; ok {
		return []int{val}
	}

	for i, v := range idx {
		for j, vv := range idx {
			if i&j == 0 {
				if v > vv {
					v, vv = vv, v
				}
				return []int{v, vv}
			}
		}
	}

	return
}
