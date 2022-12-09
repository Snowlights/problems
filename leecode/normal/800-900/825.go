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

// 837
func new21Game(n int, k int, maxPts int) float64 {
	dp := make([]float64, k+maxPts)
	s := 0.0
	for i := k; i < k+maxPts; i++ {
		if i > n {
			dp[i] = 0
		} else {
			dp[i] = 1
		}
		s += dp[i]
	}

	for i := k - 1; i >= 0; i-- {
		dp[i] = s / float64(maxPts)
		s -= dp[i+maxPts]
		s += dp[i]
	}

	return dp[0]
}

// 838
func longestMountain(arr []int) int {
	n := len(arr)
	l, r := make([]int, n), make([]int, n)

	for i, v := range arr {
		if i > 0 && v > arr[i-1] {
			l[i] = l[i-1] + 1
		} else {
			l[i] = 1
		}
	}
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && arr[i] > arr[i+1] {
			r[i] = r[i+1] + 1
		} else {
			r[i] = 1
		}
	}

	ans := 0
	for i := 1; i < n; i++ {
		if l[i] > 1 && r[i] > 1 {
			ans = max(ans, l[i]+r[i]-1)
		}
	}
	return ans
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
