package main

import "sort"

func maximumSumQueries(a []int, b []int, qs [][]int) (ans []int) {
	type node struct {
		x, y int
	}
	nodeList := make([]*node, 0, len(a))
	for i, v := range a {
		nodeList = append(nodeList, &node{v, b[i]})
	}
	for i := range qs {
		qs[i] = append(qs[i], i)
	}
	ans = make([]int, len(qs))

	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[i].x < nodeList[j].x
	})
	sort.Slice(qs, func(i, j int) bool {
		return qs[i][0] > qs[j][0]
	})

	q, i := []*node{}, len(a)-1
	for _, qu := range qs {
		for i >= 0 && nodeList[i].x >= qu[0] {
			for len(q) > 0 && q[len(q)-1].y < nodeList[i].x+nodeList[i].y {
				q = q[:len(q)-1]
			}
			if len(q) == 0 || q[len(q)-1].x < nodeList[i].y {
				q = append(q, &node{nodeList[i].y, nodeList[i].x + nodeList[i].y})
			}
			i--
		}
		j := sort.Search(len(q), func(i int) bool {
			return q[i].x >= qu[1]
		})
		if j < len(q) {
			ans[qu[2]] = q[j].y
		} else {
			ans[qu[2]] = -1
		}
	}

	return
}
