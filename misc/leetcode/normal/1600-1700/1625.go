package _600_1700

import "sort"

// 1630
func checkArithmeticSubarrays(nums []int, l []int, r []int) []bool {
	var n = len(l)
	var res = make([]bool, n)
	for i := 0; i < n; i++ {
		s, e := l[i], r[i]
		if s >= e {
			res[i] = false
			continue
		}
		tempArr := make([]int, e-s+1)
		copy(tempArr, nums[s:e+1])
		sort.Ints(tempArr)
		ans := true
		temp := tempArr[1] - tempArr[0]
		for j := 1; j < e-s+1; j++ {
			if tempArr[j]-tempArr[j-1] != temp {
				ans = false
				break
			}
		}
		res[i] = ans
	}
	return res
}

// 1631
func minimumEffortPath(heights [][]int) int {
	m, n := len(heights), len(heights[0])
	d := make([]int, 0, m*n)
	for _, v := range heights {
		d = append(d, v...)
	}
	fa := make([]int, m*n)
	for i := range fa {
		fa[i] = i
	}
	var find func(int) int
	find = func(from int) int {
		if fa[from] != from {
			fa[from] = find(fa[from])
		}
		return fa[from]
	}
	same := func(from, to int) bool {
		return find(from) == find(to)
	}
	merge := func(from, to int) {
		from, to = find(from), find(to)
		if from == to {
			return
		}
		fa[from] = to
	}

	type pair struct {
		from, to, val int
	}
	abs := func(a int) int {
		if a < 0 {
			return -a
		}
		return a
	}
	pairList := make([]pair, 0, len(heights))
	for i, v := range heights {
		for j := range v {
			num := i*n + j
			button, right := (i+1)*n+j, i*n+j+1
			if button < m*n {
				pairList = append(pairList, pair{
					from: num,
					to:   button,
					val:  abs(d[num] - d[button]),
				})
			}
			if right < (i+1)*n {
				pairList = append(pairList, pair{
					from: num,
					to:   right,
					val:  abs(d[num] - d[right]),
				})
			}
		}
	}

	sort.Slice(pairList, func(i, j int) bool {
		return pairList[i].val < pairList[j].val
	})

	for _, pair := range pairList {
		merge(pair.from, pair.to)
		if same(0, m*n-1) {
			return pair.val
		}
	}

	return 0
}
