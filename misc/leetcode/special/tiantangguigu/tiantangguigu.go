package tiantangguigu

import (
	"container/heap"
	"sort"
	"strconv"
)

// 1
func lastMaterial(material []int) int {
	h := hp2{
		material,
	}
	heap.Init(&h)
	for h.Len() > 1 {
		a, b := heap.Pop(&h).(int), heap.Pop(&h).(int)
		heap.Push(&h, abs(a-b))
	}

	return h.IntSlice[0]
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type hp2 struct{ sort.IntSlice }

func (h hp2) Less(i, j int) bool  { return h.IntSlice[i] > h.IntSlice[j] }
func (h *hp2) Push(v interface{}) { h.IntSlice = append(h.IntSlice, v.(int)) }
func (h *hp2) Pop() interface{} {
	a := h.IntSlice
	v := a[len(a)-1]
	h.IntSlice = a[:len(a)-1]
	return v
}

// 2
func longestESR(sales []int) int {
	n := len(sales)
	pre := make([]int, n+1)
	for i := 0; i < n; i++ {
		val := -1
		if sales[i] > 8 {
			val = 1
		}
		pre[i+1] = pre[i] + val
	}

	q := []int{}
	for i := 0; i <= n; i++ {
		if len(q) == 0 || pre[q[len(q)-1]] > pre[i] {
			q = append(q, i)
		}
	}
	ans := 0

	for i := n; i >= 0; i-- {
		for len(q) > 0 && pre[i] > pre[q[len(q)-1]] {
			val := q[len(q)-1]
			q = q[:len(q)-1]
			ans = max(ans, i-val)
		}
	}

	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 3

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lightDistribution(root *TreeNode) []*TreeNode {
	h := make(map[string]int)
	ans := make([]*TreeNode, 0)
	var bfs func(node *TreeNode) string
	bfs = func(n *TreeNode) string {
		if n == nil {
			return " "
		}
		val := "(" + strconv.Itoa(n.Val) + bfs(n.Left) + bfs(n.Right) + ")"
		if h[val] == 1 {
			ans = append(ans, n)
		}
		h[val]++
		return val
	}

	bfs(root)
	return ans
}

// 4
func minSupplyStationNumber(root *TreeNode) int {

	sonToP, h := make(map[*TreeNode]*TreeNode), make(map[*TreeNode]int)
	nMap := make(map[*TreeNode]bool)
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v.Left != nil {
				q = append(q, v.Left)
				sonToP[v.Left] = v
				h[v]++
			}
			if v.Right != nil {
				q = append(q, v.Right)
				sonToP[v.Right] = v
				h[v]++
			}
			nMap[v] = true
		}
	}

	for v := range nMap {
		if h[v] == 0 {
			q = append(q, v)
		}
	}

	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v.Val == 0 {
				if v.Left == nil && v.Right == nil {
					p, ok := sonToP[v]
					if ok {
						p.Val = 1
					} else {
						v.Val = 1
					}
				} else if v.Left != nil && v.Right != nil {
					if v.Left.Val == 1 || v.Right.Val == 1 {

					} else {
						p, ok := sonToP[v]
						if ok {
							p.Val = 1
						} else {
							v.Val = 1
						}
					}
				} else {
					if v.Left != nil {
						if v.Left.Val == 1 {

						} else {
							p, ok := sonToP[v]
							if ok {
								p.Val = 1
							} else {
								v.Val = 1
							}
						}
					}
					if v.Right != nil {
						if v.Right.Val == 1 {

						} else {
							p, ok := sonToP[v]
							if ok {
								p.Val = 1
							} else {
								v.Val = 1
							}
						}
					}
				}
			}

			if h[sonToP[v]]--; h[sonToP[v]] == 0 {
				q = append(q, sonToP[v])
			}
		}
	}
	ans := 0
	for v := range nMap {
		ans += v.Val
	}

	return ans
}
