package _300_1400

import "sort"

// 1333
func filterRestaurants(restaurants [][]int, veganFriendly, maxPrice, maxDistance int) []int {
	sort.Slice(restaurants, func(i, j int) bool {
		a, b := restaurants[i], restaurants[j]
		return a[1] > b[1] || a[1] == b[1] && a[0] > b[0]
	})
	ans := make([]int, 0)
	veg := func(v int) bool {
		if veganFriendly == 1 {
			return v == veganFriendly
		}
		return true
	}
	for _, r := range restaurants {
		if veg(r[2]) && r[3] <= maxPrice && r[4] <= maxDistance {
			ans = append(ans, r[0])
		}
	}
	return ans
}

// 1337
func kWeakestRows(mat [][]int, k int) []int {

	valMap := make(map[int][]int)
	for i, val := range mat {
		res := 0
		for _, v := range val {
			if v == 1 {
				res += 1
			} else {
				break
			}
		}
		valMap[res] = append(valMap[res], i)
	}

	type node struct {
		val int
	}
	nodeList := make([]*node, 0)
	for v, _ := range valMap {
		nodeList = append(nodeList, &node{val: v})
	}
	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[i].val < nodeList[j].val
	})
	vals := make([]int, 0)
	for _, v := range nodeList {
		tmpList, ok := valMap[v.val]
		if ok {
			length := len(tmpList)
			for length > 0 && k > 0 {
				vals = append(vals, tmpList[0])
				tmpList = tmpList[1:]
				length--
				k--
			}
		}
	}

	return vals
}

// 1346
func checkIfExist(arr []int) bool {
	h := make(map[int]bool)
	for _, v := range arr {
		if v%2 == 0 {
			if h[v/2] {
				return true
			}
		}
		if h[v*2] {
			return true
		}
		h[v] = true
	}
	return false
}
