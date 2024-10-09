package interview

// 面试题 04-01
func findWhetherExistsPath(n int, graph [][]int, start int, target int) bool {
	g := make([][]int, n)
	vis := make(map[int]bool)

	for _, v := range graph {
		g[v[0]] = append(g[v[0]], v[1])
	}
	q := []int{start}
	vis[start] = true
	for len(q) > 0 {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v == target {
				return true
			}
			for _, to := range g[v] {
				if !vis[to] {
					q = append(q, to)
				}
			}
		}
	}
	return false
}

// 面试题 04-02
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	idx := len(nums) / 2
	root := &TreeNode{Val: nums[idx]}
	root.Left = sortedArrayToBST(nums[:idx])
	root.Right = sortedArrayToBST(nums[idx+1:])
	return root
}

// 面试题 04-03
func listOfDepth(tree *TreeNode) []*ListNode {
	ans := make([]*ListNode, 0)
	q := []*TreeNode{tree}
	for len(q) > 0 {
		tmp := q
		q = nil
		root := &ListNode{}
		cur := root
		for _, v := range tmp {
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}
			cur.Next = &ListNode{
				Val:  v.Val,
				Next: nil,
			}
			cur = cur.Next
		}
		ans = append(ans, root.Next)
	}
	return ans
}

// 面试题 04-04
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	var help func(node *TreeNode) int
	max := func(a, b int) int {
		if a > b {
			return a
		}
		return b
	}
	abs := func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}
	help = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		left, right := help(root.Left), help(root.Right)
		return max(left, right) + 1
	}
	return abs(help(root.Left)-help(root.Right)) <= 1 &&
		isBalanced(root.Left) && isBalanced(root.Right)
}

// 从左向右遍历一个数组，通过不断将其中的元素插入树中可以逐步地生成一棵二叉搜索树。
// 给定一个由不同节点组成的二叉搜索树 root，输出所有可能生成此树的数组。

// 面试题 04-09
var res [][]int

func BSTSequencesV2(root *TreeNode) [][]int {
	res = [][]int{}
	// 测试用例规定的，空的时候返回"[[]]"
	if root == nil {
		res = append(res, []int{})
		return res
	}

	// 开始时set中只有root， path未选取任何节点，为空。
	f(map[*TreeNode]struct{}{root: struct{}{}}, []int{})
	return res
}

// set：可供选择的节点集合
// path: 当前已经选择的节点
func f(set map[*TreeNode]struct{}, path []int) {
	// 可供选择的节点集合为空，表示已经将所有节点都选择完了
	// 可以将path添加到结果集中
	if len(set) == 0 {
		// 涉及到回溯，所以将path复制了一份
		copy := []int{}
		copy = append(copy, path...)
		res = append(res, copy)
		return
	}

	// 将每一个节点都选择一次
	for k, _ := range set {
		// 复制一份可供选择的节点集合，作为下一次选择的集合
		setNext := map[*TreeNode]struct{}{}
		for a, b := range set {
			setNext[a] = b
		}
		// 将k从可供选择的集合中拿走
		delete(setNext, k)
		// 我们已经选择了k，添加已经选择的序列中
		path = append(path, k.Val)
		// 向可供选择的集合中添加下一步可以被选择的新加入的节点
		if k.Left != nil {
			setNext[k.Left] = struct{}{}
		}
		if k.Right != nil {
			setNext[k.Right] = struct{}{}
		}
		// 来到下一步
		f(setNext, path)
		// path回溯
		path = path[:len(path)-1]
	}
}
