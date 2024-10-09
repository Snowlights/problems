package simple_302

import (
	"fmt"
	"math"
	"sort"
)

// 1
func numberOfPairs(nums []int) []int {

	h := make(map[int]int)
	for _, v := range nums {
		h[v]++
	}
	a1, a2 := 0, 0
	for _, v := range h {
		a1 += v / 2
		if v%2 == 1 {
			a2++
		}
	}
	return []int{a1, a2}

}

// 2

func maximumSum(nums []int) int {

	h := make(map[int][]int)

	for _, v := range nums {
		tmp, val := 0, v
		for val > 0 {
			tmp += val % 10
			val /= 10
		}
		h[tmp] = append(h[tmp], v)
	}
	ans := -1
	for _, v := range h {
		if len(v) < 2 {
			continue
		}
		sort.Ints(v)
		if v[len(v)-1]+v[len(v)-2] > ans {
			ans = v[len(v)-1] + v[len(v)-2]
		}
	}
	return ans
}

// 3

func smallestTrimmedNumbers(nums []string, queries [][]int) []int {

	type node struct {
		idx int
		val string
	}

	h := make(map[int][]*node)
	for i, n := range nums {
		for j := len(n) - 1; j >= 0; j-- {
			h[len(n)-j] = append(h[len(n)-j], &node{
				idx: i,
				val: n[j:],
			})
		}
	}
	for k, v := range h {
		sort.Slice(v, func(i, j int) bool {
			return v[i].val < v[j].val || (v[i].val == v[j].val && v[i].idx < v[j].idx)
		})
		h[k] = v
	}
	ans := make([]int, len(queries))
	for i, q := range queries {
		v, ok := h[q[1]]
		if ok {
			ans[i] = v[q[0]-1].idx
		}
	}
	return ans
}

// 4

func minOperations(nums []int, numsDivide []int) int {

	type Node struct {
		val int
		num int
	}
	nodeList := make([]*Node, 0)
	nMap, dMap := make(map[int]*Node), make(map[int]bool)
	dMin := math.MaxInt64
	for _, v := range nums {
		n, ok := nMap[v]
		if !ok {
			n = &Node{val: v, num: 0}
			nMap[v] = n
			nodeList = append(nodeList, n)
		}
		n.num++
	}
	for _, n := range numsDivide {
		dMap[n] = true
		if n < dMin {
			dMin = n
		}
	}
	sort.Slice(nodeList, func(i, j int) bool {
		return nodeList[i].val < nodeList[j].val
	})
	var g int
	if len(numsDivide) > 1 {
		g = gcd(numsDivide[0], numsDivide[1])
		for i := 2; i < len(numsDivide); i++ {
			g = gcd(numsDivide[i], g)
		}
	} else {
		g = numsDivide[0]
	}

	ans, found := 0, false
	idx := -1
	for i, n := range nodeList {
		if g > n.val && g%n.val != 0 {
			found = true
			idx = i
			ans += n.num
			continue
		}
		break
	}
	fmt.Println(g, idx)

	if idx == len(nodeList)-1 {
		return -1
	}

	if !found {
		if g%nodeList[0].val == 0 {
			return ans
		}
		return -1
	}
	return ans
}

func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}
