package _00_500

import "sort"

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
