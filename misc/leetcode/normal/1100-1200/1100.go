package _100_1200

import . "problems/testutil/leetcode"

// 1106
func parseBoolExpr(expression string) bool {
	stack := []byte{}
	for i := range expression {
		v := expression[i]
		switch v {
		case ')':
			f, t, ans := 0, 0, 'f'
			for stack[len(stack)-1] == 'f' || stack[len(stack)-1] == 't' || stack[len(stack)-1] == ',' {
				switch stack[len(stack)-1] {
				case 'f':
					f++
				case 't':
					t++
				}
				stack = stack[:len(stack)-1]
			}
			op := stack[len(stack)-2]
			stack = stack[:len(stack)-2]
			switch op {
			case '!':
				if f > 0 {
					ans = 't'
				}
			case '&':
				if f == 0 {
					ans = 't'
				}
			case '|':
				if t > 0 {
					ans = 't'
				}
			}
			stack = append(stack, byte(ans))
		default:
			stack = append(stack, v)
		}
	}

	return stack[0] == 't'
}

// 1123
func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	ans, d := root, 0
	var dfs func(node *TreeNode, depth int) int
	dfs = func(node *TreeNode, depth int) int {
		if node == nil {
			d = max(d, depth)
			return depth
		}
		l, r := dfs(node.Left, depth+1), dfs(node.Right, depth+1)
		if l == r && d == l {
			ans = node
		}
		return max(l, r)
	}

	dfs(root, 0)
	return ans
}
