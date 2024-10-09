package _00_900

import (
	"container/list"
	"sort"
)

// / 851
func loudAndRich(richer [][]int, quiet []int) []int {
	g := make([][]int, len(quiet))
	for _, r := range richer {
		from, to := r[1], r[0]
		g[from] = append(g[from], to)
	}
	h := make(map[int]int)
	for i, q := range quiet {
		h[q] = i
	}
	dp := make([]int, len(quiet))
	for i := range dp {
		dp[i] = -1
	}

	ans := make([]int, len(quiet))
	for i := 0; i < len(quiet); i++ {
		ans[i] = quiet[i]
		q := []int{i}
		for len(q) > 0 {
			tmp := q
			q = nil
			for _, v := range tmp {
				for _, vv := range g[v] {
					val := dp[vv]
					if val != -1 {
						ans[i] = min(ans[i], val)
					} else {
						ans[i] = min(ans[i], quiet[vv])
						q = append(q, vv)
					}
				}
			}
		}
		dp[i] = ans[i]
	}

	for i, v := range ans {
		ans[i] = h[v]
	}
	return ans
}

// 852
func peakIndexInMountainArray(arr []int) int {
	return sort.Search(len(arr), func(i int) bool { return arr[i] > arr[i+1] })
}

// 854
func kSimilarity(s1 string, s2 string) int {

	next := func(s string) []string {
		i := 0
		res := make([]string, 0)
		for ; s[i] == s2[i]; i++ {
		}
		for j := i + 1; j < len(s1); j++ {
			if s[j] == s2[i] && s[j] != s2[j] {
				res = append(res, s[:i]+string(s[j])+s[i+1:j]+string(s[i])+s[j+1:])
			}
		}
		return res
	}

	q := []string{s1}
	vis := map[string]bool{s1: true}
	ans := 0
	for {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v == s2 {
				return ans
			}
			for _, vv := range next(v) {
				if vis[vv] {
					continue
				}
				q = append(q, vv)
				vis[vv] = true
			}
		}
		ans++
	}
}

// 855
type ExamRoom struct {
	seat *list.List // 表示坐着的同学的位置
	n    int
}

func Constructor855(N int) ExamRoom {
	return ExamRoom{
		seat: list.New(),
		n:    N - 1,
	}
}

func (this *ExamRoom) Seat() int {
	// 还没有人入座，直接将0插入
	if this.seat.Len() == 0 {
		this.seat.PushFront(0)
		return 0
	}
	e := this.seat.Front()
	pre := e.Value.(int)
	max := pre // 头部需要特殊判断
	addVal := 0
	addFront := e
	e = e.Next()
	for ; e != nil; e = e.Next() {
		val := e.Value.(int)
		distant := (val - pre) / 2 // 两点之间的最远距离
		if distant > max {
			max = distant
			addFront = e           // 需要插入的点的后一个元素。方便找到后直接插入
			addVal = pre + distant // 需要插入的位置
		}
		pre = val
	}
	distant := this.n - pre // 尾部特殊判断
	if distant > max {
		this.seat.PushBack(this.n) // 直接插入到链表尾部
		return this.n
	}
	this.seat.InsertBefore(addVal, addFront) // 插入
	return addVal
}

// 遍历知道对应的位置删除
func (this *ExamRoom) Leave(p int) {
	for e := this.seat.Front(); e != nil; e = e.Next() {
		if e.Value.(int) == p {
			this.seat.Remove(e)
			return
		}
	}
	return
}

// 862
func shortestSubarray(nums []int, k int) int {
	n := len(nums)
	s := make([]int, n+1)
	for i, x := range nums {
		s[i+1] = s[i] + x // 计算前缀和
	}
	ans := n + 1
	q := []int{}
	for i, curS := range s {
		for len(q) > 0 && curS-s[q[0]] >= k {
			ans = min(ans, i-q[0])
			q = q[1:]
		}
		for len(q) > 0 && s[q[len(q)-1]] >= curS {
			q = q[:len(q)-1]
		}
		q = append(q, i)
	}
	if ans > n {
		return -1
	}
	return ans
}

// 863
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func distanceK(root *TreeNode, target *TreeNode, k int) []int {
	g, h := make(map[int]int), make(map[int]*TreeNode)
	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			h[v.Val] = v
			if v.Left != nil {
				g[v.Left.Val] = v.Val
				q = append(q, v.Left)
			}
			if v.Right != nil {
				g[v.Right.Val] = v.Val
				q = append(q, v.Right)
			}
		}
	}

	idx := 0
	qL := []int{target.Val}
	vis := make(map[int]bool)
	vis[target.Val] = true
	for idx < k {
		tmp := qL
		qL = nil

		for _, v := range tmp {
			node, ok := h[v]
			if ok {
				if node.Left != nil && !vis[node.Left.Val] {
					qL = append(qL, node.Left.Val)
					vis[node.Left.Val] = true
				}
				if node.Right != nil && !vis[node.Right.Val] {
					qL = append(qL, node.Right.Val)
					vis[node.Right.Val] = true
				}
			}
			fa, ok := g[v]
			if ok && !vis[fa] {
				qL = append(qL, fa)
				vis[fa] = true
			}
		}
		idx++
	}

	return qL
}

// 864
func shortestPathAllKeys(grid []string) int {
	m, n := len(grid), len(grid[0])
	ans, key, mask := 0, 0, 0
	startx, starty := 0, 0
	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] >= 'a' && grid[i][j] <= 'f' {
				key |= 1 << (grid[i][j] - 'a')
			}
			if grid[i][j] == '@' {
				startx, starty = i, j
			}
		}
	}
	vis := make([][][1 << 7]bool, m)
	for i := range vis {
		vis[i] = make([][1 << 7]bool, n)
	}
	vis[startx][starty][mask] = true

	dir := [][]int{{0, -1}, {0, 1}, {1, 0}, {-1, 0}}
	q := [][]int{{startx, starty, mask}}
	for len(q) > 0 {
		tmp := q
		q = nil

		for _, v := range tmp {
			x, y, cur := v[0], v[1], v[2]
			if cur == key {
				return ans
			}
			for _, d := range dir {
				xx, yy := x+d[0], y+d[1]
				if 0 <= xx && xx < m && 0 <= yy && yy < n && grid[xx][yy] != '#' {
					if grid[xx][yy] == '.' || grid[xx][yy] == '@' {
						if !vis[xx][yy][cur] {
							q = append(q, []int{xx, yy, cur})
							vis[xx][yy][cur] = true
						}
					} else if grid[xx][yy] >= 'A' && grid[xx][yy] <= 'F' {
						if cur>>(grid[xx][yy]-'A')&1 == 1 {
							q = append(q, []int{xx, yy, cur})
							vis[xx][yy][cur] = true
						}
					} else if grid[xx][yy] >= 'a' && grid[xx][yy] <= 'f' {
						newMask := cur | 1<<(grid[xx][yy]-'a')
						if !vis[xx][yy][newMask] {
							q = append(q, []int{xx, yy, newMask})
							vis[xx][yy][newMask] = true
						}
					}
				}
			}
		}
		ans++
	}

	return -1
}

// 865
func subtreeWithAllDeepest(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	pre, q := []*TreeNode{}, []*TreeNode{root}
	g := make(map[int]*TreeNode)
	for len(q) > 0 {
		pre = q
		q = nil
		for _, v := range pre {
			if v.Left != nil {
				g[v.Left.Val] = v
				q = append(q, v.Left)
			}
			if v.Right != nil {
				g[v.Right.Val] = v
				q = append(q, v.Right)
			}
		}
	}

	var last map[*TreeNode]bool
	for _, v := range pre {
		ans := make(map[*TreeNode]bool)
		for {
			ans[v] = true
			v = g[v.Val]
			if v == nil {
				break
			}
		}
		if last == nil {
			last = ans
		} else {
			l := make(map[*TreeNode]bool)
			for v := range ans {
				if last[v] {
					l[v] = true
				}
			}
			last = l
		}
	}

	v := pre[0]
	for {
		if last[v] {
			return v
		}
		v = g[v.Val]
	}

}

// 870
func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	ans := make([]int, n)
	sort.Ints(nums1)
	ids := make([]int, 0, n)
	for i := range nums2 {
		ids = append(ids, i)
	}
	sort.Slice(ids, func(i int, j int) bool {
		return nums2[ids[i]] < nums2[ids[j]]
	})
	l, r := 0, n-1
	for _, v := range nums1 {
		if v > nums2[ids[l]] {
			ans[ids[l]] = v
			l++
		} else {
			ans[ids[r]] = v
			r--
		}
	}

	return ans
}
