package _00_900

import (
	"sort"
)

// 926
func maxProfitAssignment(difficulty []int, profit []int, worker []int) int {
	type pair struct {
		diff, pro int
	}
	pairList := make([]*pair, 0)
	for i, v := range difficulty {
		pairList = append(pairList, &pair{
			diff: v,
			pro:  profit[i],
		})
	}
	sort.Slice(pairList, func(i, j int) bool {
		return pairList[i].diff < pairList[j].diff
	})

	ans := 0
	for _, w := range worker {
		tmp, idx := 0, 0
		for idx < len(pairList) && pairList[idx].diff <= w {
			tmp = max(tmp, pairList[idx].pro)
			idx++
		}
		ans += tmp
	}
	return ans
}

// 828
func uniqueLetterString(s string) (ans int) {
	sum, last := 0, [26][2]int{}
	for i := range last {
		last[i] = [2]int{-1, -1}
	}
	for i, c := range s {
		c -= 'A'

		sum += i - last[c][0]*2 + last[c][1]

		ans += sum

		last[c][1] = last[c][0]

		last[c][0] = i
	}
	return
}

// 841
func canVisitAllRooms(rooms [][]int) bool {
	vis := make(map[int]bool)
	vis[0] = true
	q := []int{0}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			for _, vv := range rooms[v] {
				if vis[vv] {
					continue
				}
				q = append(q, vv)
				vis[vv] = true
			}
		}
	}
	return len(rooms) == len(vis)
}
