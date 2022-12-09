package template

import "sort"

// 最长递增子序列
func lengthOfLIS(nums []int) int {

	//使得f按照升序排列
	f := []int{}
	for _, e := range nums {
		//h := e
		if i := sort.SearchInts(f, e); i < len(f) {
			f[i] = e
		} else {
			f = append(f, e)
		}
	}
	return len(f)
}

// 回文串中心扩展
func countSubstrings(s string) int {
	n := len(s)
	ans := 0
	for i := 0; i < 2*n-1; i++ {
		l, r := i/2, i/2+i%2
		for l >= 0 && r < n && s[l] == s[r] {
			l--
			r++
			ans++
		}
	}
	return ans
}
