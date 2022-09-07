package offer

import "strings"

// offer 57
func twoSum(nums []int, target int) []int {

	h := make(map[int]int)
	for i, v := range nums {
		if _, ok := h[target-v]; ok {
			return []int{target - v, v}
		}
		h[v] = i
	}
	return nil
}

// offer 58
func reverseLeftWords(s string, n int) string {
	return s[n:] + s[:n]
}

// offer 58-1
func reverseWords(s string) string {
	s = strings.Trim(s, " ")
	parts := strings.Split(s, " ")
	realParts := make([]string, 0, len(parts))
	for _, p := range parts {
		if len(p) == 0 {
			continue
		}
		realParts = append(realParts, p)
	}
	st, e := 0, len(realParts)-1
	for st < e {
		realParts[st], realParts[e] = realParts[e], realParts[st]
		st++
		e--
	}
	return strings.Join(realParts, " ")
}

// offer 63
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	ans, preMin := 0, prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i]-preMin > ans {
			ans = prices[i] - preMin
		}
		if preMin > prices[i] {
			preMin = prices[i]
		}
	}
	return ans
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
