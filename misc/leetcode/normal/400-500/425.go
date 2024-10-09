package _00_500

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

// 429
type Node struct {
	Val      int
	Children []*Node
}

func levelOrder(root *Node) [][]int {
	if root == nil {
		return nil
	}
	ans, q := make([][]int, 0), []*Node{root}
	for len(q) > 0 {
		arr, tmp := make([]int, 0), q
		q = nil
		for _, n := range tmp {
			arr = append(arr, n.Val)
			q = append(q, n.Children...)
		}
		ans = append(ans, arr)
	}
	return ans
}

// 437
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func rootSum(root *TreeNode, targetSum int) (res int) {
	if root == nil {
		return
	}
	val := root.Val
	if val == targetSum {
		res++
	}
	res += rootSum(root.Left, targetSum-val)
	res += rootSum(root.Right, targetSum-val)
	return
}

func pathSum(root *TreeNode, targetSum int) int {
	if root == nil {
		return 0
	}
	res := rootSum(root, targetSum)
	res += pathSum(root.Left, targetSum)
	res += pathSum(root.Right, targetSum)
	return res
}

// 438
func findAnagrams(s string, p string) []int {
	if len(p) > len(s) {
		return nil
	}
	ph := make(map[byte]int)
	sh := make(map[byte]int)
	for _, ps := range p {
		ph[byte(ps)]++
	}
	ans := make([]int, 0)
	start, end := 0, len(p)
	for i := start; i < end; i++ {
		sh[byte(s[i])]++
	}
	for end < len(s) {
		if equal(sh, ph) {
			ans = append(ans, start)
		}
		sh[s[start]]--
		sh[s[end]]++
		start++
		end++
	}

	return ans
}

func equal(ph, sh map[byte]int) bool {
	for k, v := range ph {
		if sh[k] != v {
			return false
		}
	}
	return true
}

// 441
func arrangeCoins(n int) int {
	return sort.Search(n, func(i int) bool {
		i++
		return (i+1)*i > 2*n
	})
}

// 445
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	s1, s2 := []int{}, []int{}
	for l1 != nil {
		s1 = append(s1, l1.Val)
		l1 = l1.Next
	}
	for l2 != nil {
		s2 = append(s2, l2.Val)
		l2 = l2.Next
	}
	var head *ListNode
	val, flag := 0, 0
	for len(s1) > 0 || len(s2) > 0 {
		a, b := 0, 0
		if len(s1) > 0 {
			a = s1[len(s1)-1]
			s1 = s1[:len(s1)-1]
		}
		if len(s2) > 0 {
			b = s2[len(s2)-1]
			s2 = s2[:len(s2)-1]
		}
		val, flag = (a+b+flag)%10, (a+b+flag)/10
		if head == nil {
			head = &ListNode{
				Val: val,
			}
		} else {
			node := &ListNode{
				Val:  val,
				Next: head,
			}
			head = node
		}
	}
	if flag > 0 {
		node := &ListNode{
			Val:  flag,
			Next: head,
		}
		head = node
	}

	return head
}

// 449
type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}
	data := &strings.Builder{}
	var dfs func(*TreeNode)
	dfs = func(root *TreeNode) {
		if root == nil {
			return
		}
		data.WriteString(strconv.Itoa(root.Val))
		data.WriteByte(' ')
		dfs(root.Left)
		dfs(root.Right)
	}
	dfs(root)
	return data.String()[0 : data.Len()-1]
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	vals := strings.Split(data, " ")
	i := 0
	var dfs func(int, int) *TreeNode
	dfs = func(mi, mx int) *TreeNode {
		if i == len(vals) {
			return nil
		}
		x, _ := strconv.Atoi(vals[i])
		if x < mi || x > mx {
			return nil
		}
		i++
		root := &TreeNode{Val: x}
		root.Left = dfs(mi, x)
		root.Right = dfs(x, mx)
		return root
	}
	return dfs(math.MinInt64, math.MaxInt64)
}
