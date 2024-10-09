package algorithm

func _() {
	// 从 n 个元素中选择 r 个元素，按字典序生成所有组合，每个组合用下标表示  r <= n
	// 由于实现上直接传入了 indexes，所以在 do 中不能修改 ids。若要修改则代码在传入前需要 copy 一份
	// 参考 https://docs.python.org/3/library/itertools.html#itertools.combinations
	// https://stackoverflow.com/questions/41694722/algorithm-for-itertools-combinations-in-python
	combinations := func(n, r int, do func(ids []int) (Break bool)) {
		ids := make([]int, r)
		for i := range ids {
			ids[i] = i
		}
		if do(ids) {
			return
		}
		for {
			i := r - 1
			for ; i >= 0; i-- {
				if ids[i] != i+n-r {
					break
				}
			}
			if i == -1 {
				return
			}
			ids[i]++
			for j := i + 1; j < r; j++ {
				ids[j] = ids[j-1] + 1
			}
			if do(ids) {
				return
			}
		}
	}

	// 从一个长度为 n 的数组中选择 r 个元素，按字典序生成所有排列，每个排列用下标表示  r <= n
	// 由于实现上直接传入了 indexes，所以在 do 中不能修改 ids。若要修改则代码在传入前需要 copy 一份
	// 参考 https://docs.python.org/3/library/itertools.html#itertools.permutations
	permutations := func(n, r int, do func(ids []int) (Break bool)) {
		ids := make([]int, n)
		for i := range ids {
			ids[i] = i
		}
		if do(ids[:r]) {
			return
		}
		cycles := make([]int, r)
		for i := range cycles {
			cycles[i] = n - i
		}
		for {
			i := r - 1
			for ; i >= 0; i-- {
				cycles[i]--
				if cycles[i] == 0 {
					tmp := ids[i]
					copy(ids[i:], ids[i+1:])
					ids[n-1] = tmp
					cycles[i] = n - i
				} else {
					j := cycles[i]
					ids[i], ids[n-j] = ids[n-j], ids[i]
					if do(ids[:r]) {
						return
					}
					break
				}
			}
			if i == -1 {
				return
			}
		}
	}

	_ = []interface{}{combinations, permutations}
}
