package _00_300

// 226
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left, root.Right = root.Right, root.Left
	invertTree(root.Left)
	invertTree(root.Right)
	return root
}

// 227
func calculate227(s string) (ans int) {
	stack := []int{}
	preSign := '+'
	num := 0
	for i, ch := range s {
		isDigit := '0' <= ch && ch <= '9'
		if isDigit {
			num = num*10 + int(ch-'0')
		}
		if !isDigit && ch != ' ' || i == len(s)-1 {
			switch preSign {
			case '+':
				stack = append(stack, num)
			case '-':
				stack = append(stack, -num)
			case '*':
				stack[len(stack)-1] *= num
			default:
				stack[len(stack)-1] /= num
			}
			preSign = ch
			num = 0
		}
	}
	for _, v := range stack {
		ans += v
	}
	return
}

// 230
func kthSmallest(root *TreeNode, k int) int {
	stack := []*TreeNode{}
	for {
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		stack, root = stack[:len(stack)-1], stack[len(stack)-1]
		k--
		if k == 0 {
			return root.Val
		}
		root = root.Right
	}
}

// 235
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	stoP, valToNode := make(map[int]*TreeNode), make(map[int]*TreeNode)
	que := []*TreeNode{root}
	stoP[root.Val] = root
	for len(que) > 0 {
		tmp := que
		que = nil
		for _, v := range tmp {
			valToNode[v.Val] = v
			if v.Left != nil {
				stoP[v.Left.Val] = v
				que = append(que, v.Left)
			}
			if v.Right != nil {
				stoP[v.Right.Val] = v
				que = append(que, v.Right)
			}
		}
	}

	pn, qn := valToNode[p.Val], valToNode[q.Val]
	pMap := make(map[int]bool)
	pMap[pn.Val] = true
	for {
		parent := stoP[pn.Val]
		pMap[parent.Val] = true
		if parent == pn {
			break
		}
		pn = parent
	}
	if pMap[qn.Val] {
		return valToNode[qn.Val]
	}
	for {
		parent := stoP[qn.Val]
		if pMap[parent.Val] {
			return valToNode[parent.Val]
		}
		qn = parent
	}

}
