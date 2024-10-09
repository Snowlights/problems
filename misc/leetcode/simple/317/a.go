package simple_317

import "sort"

// 1
func averageValue(nums []int) int {
	total, num := 0, 0
	for _, v := range nums {
		if v%3 == 0 && v%2 == 0 {
			total += v
			num++
		}
	}
	if num == 0 {
		return 0
	}
	return total / num
}

// 2
func mostPopularCreator(creators []string, ids []string, views []int) [][]string {
	type pair struct {
		name string
		view int
		vv   []struct {
			id   string
			view int
		}
	}
	pList, h := make([]*pair, 0), make(map[string]*pair)
	for i, c := range creators {
		p, ok := h[c]
		if !ok {
			p = &pair{
				name: c,
			}
			pList = append(pList, p)
			h[c] = p
		}
		p.vv = append(p.vv, struct {
			id   string
			view int
		}{id: ids[i], view: views[i]})
		p.view += views[i]
	}

	sort.Slice(pList, func(i, j int) bool {
		return pList[i].view > pList[j].view
	})
	v := pList[0].view
	ans := make([][]string, 0)
	for _, p := range pList {
		if p.view == v {
			sort.Slice(p.vv, func(i, j int) bool {
				return p.vv[i].view > p.vv[j].view || (p.vv[i].view == p.vv[j].view &&
					p.vv[i].id < p.vv[j].id)
			})
			ans = append(ans, []string{p.name, p.vv[0].id})
		}
	}
	return ans
}

// 3
func makeIntegerBeautiful(n int64, target int) int64 {
	for tail := int64(1); ; tail *= 10 {
		m := (tail-(n%tail))%tail + n
		tmp, res := int64(0), m
		for m > 0 {
			tmp += m % 10
			m /= 10
		}
		if tmp <= int64(target) {
			return res - n
		}
	}
}

// 4

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func treeQueries(root *TreeNode, queries []int) []int {
	h := make(map[*TreeNode]int)
	var height func(node *TreeNode) int
	height = func(node *TreeNode) int {
		if node == nil {
			return 0
		}
		h[node] = 1 + max(height(node.Left), height(node.Right))
		return h[node]
	}
	height(root)

	res := make([]int, len(h)+1)
	var dfs func(*TreeNode, int, int)
	dfs = func(node *TreeNode, depth, restH int) {
		if node == nil {
			return
		}
		depth++
		res[node.Val] = restH
		dfs(node.Left, depth, max(restH, depth+h[node.Right]))
		dfs(node.Right, depth, max(restH, depth+h[node.Left]))
	}
	dfs(root, -1, 0)

	for i, q := range queries {
		queries[i] = res[q]
	}

	return queries
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
