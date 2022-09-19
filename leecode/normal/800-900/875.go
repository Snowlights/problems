package _00_900

import "sort"

// 875
func minEatingSpeed(piles []int, h int) int {
	return sort.Search(1e9, func(i int) bool {
		tmp := 0
		if i == 0 {
			return false
		}
		for _, p := range piles {
			tmp += p / i
			if p%i > 0 {
				tmp++
			}
		}
		return tmp <= h
	})
}

// 876
type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	pre, next := head, head
	for next != nil && next.Next != nil {
		pre = pre.Next
		next = next.Next.Next
	}
	return pre
}

// 886
const (
	UNCOLORED, RED, GREEN = 0, 1, 2
)

func possibleBipartition(n int, dislikes [][]int) bool {

	color := make([]int, n+1)
	g := make(map[int][]int)
	for _, dis := range dislikes {
		g[dis[0]] = append(g[dis[0]], dis[1])
		g[dis[1]] = append(g[dis[1]], dis[0])
	}

	for i := 1; i <= n; i++ {
		if color[i] == UNCOLORED {
			color[i] = RED
			q := []int{i}
			for len(q) > 0 {
				tmp := q
				q = nil
				for _, v := range tmp {
					c := RED
					if color[v] == RED {
						c = GREEN
					}
					for _, vv := range g[v] {
						if color[vv] == UNCOLORED {
							color[vv] = c
							q = append(q, vv)
						} else if color[vv] != c {
							return false
						}
					}
				}
			}
		}
	}

	return true
}

func buildGraph(dislikes [][]int, n int, graph [][]bool) {
	for _, d := range dislikes {
		graph[d[0]][d[1]] = true
		graph[d[1]][d[0]] = true
	}
}

// 896
func isMonotonic(nums []int) bool {
	return sort.IsSorted(sort.IntSlice(nums)) || sort.IsSorted(sort.Reverse(sort.IntSlice(nums)))
}
