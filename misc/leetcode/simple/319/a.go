package simple_319

import "sort"

// 1
func convertTemperature(celsius float64) []float64 {
	return []float64{celsius + 273.15, celsius*1.80 + 32}
}

// 2
func subarrayLCM(nums []int, k int) int {
	ans := 0
	for i := range nums {
		l := 1
		for _, x := range nums[i:] {
			l = lcm(l, x)
			if k%l > 0 {
				break
			}
			if l == k {
				ans++
			}
		}
	}
	return ans
}

func gcd(a, b int) int {
	for a%b != 0 {
		a, b = b, a%b
	}
	return b
}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

// 3
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minimumOperations(root *TreeNode) int {
	ans := 0

	q := []*TreeNode{root}
	for len(q) > 0 {
		tmp := q
		q = nil
		num := []int{}
		for _, v := range tmp {
			if v.Left != nil {
				q = append(q, v.Left)
			}
			if v.Right != nil {
				q = append(q, v.Right)
			}

			num = append(num, v.Val)
		}
		h, vis := make(map[int]int), make(map[int]bool)
		for i, v := range num {
			h[v] = i
		}
		sort.Ints(num)

		for i := range num {
			if vis[i] {
				continue
			}
			tmp, l := i, 0
			for !vis[tmp] {
				l++
				vis[tmp] = true
				tmp = h[num[tmp]]
			}
			ans += l - 1
		}
	}
	return ans
}

// 4

func maxPalindromes(s string, k int) int {
	n := len(s)
	f := make([]int, n+1)

	for i := 1; i < 2*n-1; i++ {
		l, r := i/2, 1+i%2
		for 0 <= l && r < n && s[l] == s[r] {
			if r-l+1 >= k {
				f[r+1] = max(f[r+1], f[l]+1)
			}
			l--
			r++
		}
	}

	return f[n]
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
