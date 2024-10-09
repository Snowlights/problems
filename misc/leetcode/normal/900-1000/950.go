package _00_1000

import "sort"

// 951
func flipEquiv(root1 *TreeNode, root2 *TreeNode) bool {
	var equal func(a, b *TreeNode) bool
	equal = func(a, b *TreeNode) bool {
		if a == nil && b == nil {
			return true
		}
		if a == nil || b == nil {
			return false
		}

		if a.Val != b.Val {
			return false
		}

		return (equal(a.Left, b.Left) && equal(a.Right, b.Right)) ||
			(equal(a.Left, b.Right) && equal(a.Right, b.Left))
	}
	return equal(root1, root2)
}

// 959
func regionsBySlashes(grid []string) int {
	n := len(grid)
	a := make([]int, 4*n*n)
	uf := newUnionFind(4*n*n, a)
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			num := i*n + j
			if i < n-1 {
				button := num + n
				uf.merge(num*4+2, 4*button)
			}
			if j < n-1 {
				right := num + 1
				uf.merge(4*num+1, 4*right+3)
			}
			switch grid[i][j] {
			case '/':
				uf.merge(4*num, 4*num+3)
				uf.merge(4*num+1, 4*num+2)
			case '\\':
				uf.merge(4*num, 4*num+1)
				uf.merge(4*num+2, 4*num+3)
			default:
				uf.merge(4*num, 4*num+1)
				uf.merge(4*num+1, 4*num+2)
				uf.merge(4*num+2, 4*num+3)
			}
		}
	}
	return uf.groups
}

type uf struct {
	fa, sz  []int
	groups  int
	maxSize int
}

func newUnionFind(n int, a []int) uf {
	fa := make([]int, n)
	sz := make([]int, n)
	for i := range fa {
		fa[i] = i
		sz[i] = a[i]
	}
	return uf{fa, sz, n, 1}
}

func (u uf) find(x int) int {
	if u.fa[x] != x {
		u.fa[x] = u.find(u.fa[x])
	}
	return u.fa[x]
}

func (u *uf) merge(from, to int) (isNewMerge int) {
	x, y := u.find(from), u.find(to)
	if x == y {
		return 0
	}
	u.fa[x] = y
	u.sz[y] += u.sz[x]
	if u.sz[y] > u.maxSize {
		u.maxSize = u.sz[y]
	}
	u.groups--
	return u.sz[y]
}

func (u uf) same(x, y int) bool { return u.find(x) == u.find(y) }
func (u uf) size(x int) int     { return u.sz[u.find(x)] }

// 968
var res int

func minCameraCover(root *TreeNode) int {

	res = 0
	if dfs(root) == 1 {
		// 如果最顶端的元素没有被安装，也没有被监视
		// 那么这个节点需要安装一个摄像头
		res++
	}
	return res
}

// 1未被监视, 2安装摄像头
func dfs(root *TreeNode) int {

	// 最后一个节点不关心
	if root == nil {
		return 0
	}

	left := dfs(root.Left)
	right := dfs(root.Right)

	// 如果子节点没有被监视
	if left == 1 || right == 1 {
		// 安装摄像头
		res++
		return 2
	}

	// 如果子节点安装摄像头
	if left == 2 || right == 2 {
		// 标记被监视
		return 3
	}
	return 1
}

// 971
func flipMatchVoyage(root *TreeNode, voyage []int) (ans []int) {
	var idx int
	var n int = len(voyage)
	var dfs func(root *TreeNode) bool
	dfs = func(root *TreeNode) bool {
		if root == nil {
			return true
		}

		if root.Val != voyage[idx] {
			return false
		}

		if idx+1 < n && root.Left != nil && root.Right != nil &&
			root.Left.Val != voyage[idx+1] && root.Right.Val == voyage[idx+1] {
			ans = append(ans, root.Val)
			root.Left, root.Right = root.Right, root.Left
		}
		idx++
		return dfs(root.Left) && dfs(root.Right)
	}
	if dfs(root) {
		return ans
	}
	return []int{-1}
}

// 973
func kClosest(points [][]int, k int) [][]int {

	type node struct {
		x, y, num int
	}
	nodeList := make([]*node, 0, len(points))
	for _, p := range points {
		nodeList = append(nodeList, &node{
			x:   p[0],
			y:   p[1],
			num: p[0]*p[0] + p[1]*p[1],
		})
	}
	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[i].num < nodeList[j].num
	})
	ans := make([][]int, 0, k)
	for i := 0; i < k; i++ {
		ans = append(ans, []int{nodeList[i].x, nodeList[i].y})
	}
	return ans
}
