package _00_200

import (
	"math"
	"sort"
)

// 127
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordId := map[string]int{}
	graph := [][]int{}
	addWord := func(word string) int {
		id, has := wordId[word]
		if !has {
			id = len(wordId)
			wordId[word] = id
			graph = append(graph, []int{})
		}
		return id
	}
	addEdge := func(word string) int {
		id1 := addWord(word)
		s := []byte(word)
		for i, b := range s {
			s[i] = '*'
			id2 := addWord(string(s))
			graph[id1] = append(graph[id1], id2)
			graph[id2] = append(graph[id2], id1)
			s[i] = b
		}
		return id1
	}

	for _, word := range wordList {
		addEdge(word)
	}
	beginId := addEdge(beginWord)
	endId, has := wordId[endWord]
	if !has {
		return 0
	}

	const inf int = math.MaxInt64
	dist := make([]int, len(wordId))
	for i := range dist {
		dist[i] = inf
	}
	dist[beginId] = 0
	queue := []int{beginId}
	for len(queue) > 0 {
		v := queue[0]
		queue = queue[1:]
		if v == endId {
			return dist[endId]/2 + 1
		}
		for _, w := range graph[v] {
			if dist[w] == inf {
				dist[w] = dist[v] + 1
				queue = append(queue, w)
			}
		}
	}
	return 0
}

func parseWord(word string) map[byte]int {
	h := make(map[byte]int)
	for i := range word {
		h[word[i]]++
	}
	return h
}

func equalWithOne(a, b map[byte]int) bool {
	ans := 0
	for i := 'a'; i <= 'z'; i++ {
		if a[byte(i)] != b[byte(i)] {
			ans += abs(a[byte(i)] - b[byte(i)])
		}
	}
	if ans == 2 || ans == 0 {
		return true
	}
	return false
}

// 131
func partition(s string) [][]string {
	ans := make([][]string, 0)

	isPartition := func(s string) bool {
		t, e := 0, len(s)-1
		for t <= e {
			if s[t] != s[e] {
				return false
			}
			t++
			e--
		}
		return true
	}
	var build func(s string, idx int, tmp []string)
	build = func(s string, idx int, tmp []string) {
		if idx == len(s) {
			ans = append(ans, append([]string{}, tmp...))
			return
		}

		for i := idx; i < len(s); i++ {
			if isPartition(s[idx : i+1]) {
				tmp = append(tmp, s[idx:i+1])
				build(s, i+1, tmp)
				tmp = tmp[:len(tmp)-1]
			}
		}
	}
	build(s, 0, []string{})

	return ans
}

// 131
type NodeN struct {
	Val       int
	Neighbors []*Node
}

//func cloneGraph(node *Node) *Node {
//	visited := map[*Node]*Node{}
//	var cg func(node *Node) *Node
//	cg = func(node *Node) *Node {
//		if node == nil {
//			return node
//		}
//
//		// 如果该节点已经被访问过了，则直接从哈希表中取出对应的克隆节点返回
//		if _, ok := visited[node]; ok {
//			return visited[node]
//		}
//
//		// 克隆节点，注意到为了深拷贝我们不会克隆它的邻居的列表
//		cloneNode := &Node{node.Val, []*Node{}}
//		// 哈希表存储
//		visited[node] = cloneNode
//
//		// 遍历该节点的邻居并更新克隆节点的邻居列表
//		for _, n := range node.Neighbors {
//			cloneNode.Neighbors = append(cloneNode.Neighbors, cg(n))
//		}
//		return cloneNode
//	}
//	return cg(node)
//}

// 134
func canCompleteCircuit(gas []int, cost []int) int {
	n := len(gas)
	for i := 0; i < len(gas); {
		sum, tmp := gas[i], i
		for {
			sum -= cost[tmp%n]
			if sum < 0 {
				tmp++
				i = tmp
				break
			}
			tmp++
			sum += gas[tmp%n]
			if tmp%n == i {
				return i
			}
		}
	}

	return -1
}

// 135
func candy(ratings []int) (ans int) {
	n := len(ratings)
	left := make([]int, n)
	for i, r := range ratings {
		if i > 0 && r > ratings[i-1] {
			left[i] = left[i-1] + 1
		} else {
			left[i] = 1
		}
	}
	right := 0
	for i := n - 1; i >= 0; i-- {
		if i < n-1 && ratings[i] > ratings[i+1] {
			right++
		} else {
			right = 1
		}
		ans += max(left[i], right)
	}
	return
}

// 138

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

var cachedNode map[*Node]*Node

func deepCopy(node *Node) *Node {
	if node == nil {
		return nil
	}
	if n, has := cachedNode[node]; has {
		return n
	}
	newNode := &Node{Val: node.Val}
	cachedNode[node] = newNode
	newNode.Next = deepCopy(node.Next)
	newNode.Random = deepCopy(node.Random)
	return newNode
}

func copyRandomList(head *Node) *Node {
	cachedNode = map[*Node]*Node{}
	return deepCopy(head)
}

// 139
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool)
	for _, w := range wordDict {
		wordDictSet[w] = true
	}
	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

// 142
type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	h := make(map[*ListNode]int)
	idx := 0
	for head != nil {
		if _, ok := h[head]; ok {
			return head
		}
		h[head] = idx
		idx++
		head = head.Next
	}
	return nil
}

// 143
func reorderList(head *ListNode) {
	h := make(map[int]*ListNode)
	cur, idx := head, 0
	var pre *ListNode
	for cur != nil {
		h[idx] = cur
		idx++
		pre = cur
		cur = cur.Next
		pre.Next = nil
	}
	s, e := 1, idx-1
	pre = head

	for s <= e {
		if s <= e {
			pre.Next = h[e]
			e--
			pre = pre.Next
		}
		if s <= e {
			pre.Next = h[s]
			pre = pre.Next
			s++
		}
	}
}

// 144
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	q := []*TreeNode{root}
	ans := make([]int, 0)
	for len(q) > 0 {
		node := q[len(q)-1]
		q = q[:len(q)-1]
		if node.Right != nil {
			q = append(q, node.Right)
		}
		if node.Left != nil {
			q = append(q, node.Left)
		}
		ans = append(ans, node.Val)
	}

	return ans
}

// 145
func postorderTraversal(root *TreeNode) (res []int) {
	stk := []*TreeNode{}
	var prev *TreeNode
	for root != nil || len(stk) > 0 {
		for root != nil {
			stk = append(stk, root)
			root = root.Left
		}
		root = stk[len(stk)-1]
		stk = stk[:len(stk)-1]
		if root.Right == nil || root.Right == prev {
			res = append(res, root.Val)
			prev = root
			root = nil
		} else {
			stk = append(stk, root)
			root = root.Right
		}
	}
	return
}

// 138
func sortList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}

	h, ll := make(map[int]*ListNode), []*ListNode{}
	var pre *ListNode
	for head != nil {
		h[head.Val] = head
		ll = append(ll, head)
		pre = head
		head = head.Next
		pre.Next = nil
	}
	sort.Slice(ll, func(i, j int) bool {
		return ll[i].Val < ll[j].Val
	})
	newHead := &ListNode{
		Val:  0,
		Next: nil,
	}
	pre = newHead
	for _, v := range ll {
		pre.Next = v
		pre = pre.Next
	}

	return newHead.Next
}

// 139
func maxPoints(points [][]int) (ans int) {
	n := len(points)
	if n <= 2 {
		return n
	}
	for i, p := range points {
		if ans >= n-i || ans > n/2 {
			break
		}
		cnt := map[int]int{}
		for _, q := range points[i+1:] {
			x, y := p[0]-q[0], p[1]-q[1]
			if x == 0 {
				y = 1
			} else if y == 0 {
				x = 1
			} else {
				if y < 0 {
					x, y = -x, -y
				}
				g := gcd(abs(x), abs(y))
				x /= g
				y /= g
			}
			cnt[y+x*20001]++
		}
		for _, c := range cnt {
			ans = max(ans, c+1)
		}
	}
	return
}

func gcd(a, b int) int {
	for a != 0 {
		a, b = b%a, a
	}
	return b
}
