package _00_900

import "sort"

// 852
func peakIndexInMountainArray(arr []int) int {
	return sort.Search(len(arr), func(i int) bool { return arr[i] > arr[i+1] })
}

// 854
func kSimilarity(s1 string, s2 string) int {

	next := func(s string) []string {
		i := 0
		res := make([]string, 0)
		for ; s[i] == s2[i]; i++ {
		}
		for j := i + 1; j < len(s1); j++ {
			if s[j] == s2[i] && s[j] != s2[j] {
				res = append(res, s[:i]+string(s[j])+s[i+1:j]+string(s[i])+s[j+1:])
			}
		}
		return res
	}

	q := []string{s1}
	vis := map[string]bool{s1: true}
	ans := 0
	for {
		tmp := q
		q = nil
		for _, v := range tmp {
			if v == s2 {
				return ans
			}
			for _, vv := range next(v) {
				if vis[vv] {
					continue
				}
				q = append(q, vv)
				vis[vv] = true
			}
		}
		ans++
	}
}

// 870
func advantageCount(nums1 []int, nums2 []int) []int {
	n := len(nums1)
	ans := make([]int, n)
	sort.Ints(nums1)
	ids := make([]int, 0, n)
	for i := range nums2 {
		ids = append(ids, i)
	}
	sort.Slice(ids, func(i int, j int) bool {
		return nums2[ids[i]] < nums2[ids[j]]
	})
	l, r := 0, n-1
	for _, v := range nums1 {
		if v > nums2[ids[l]] {
			ans[ids[l]] = v
			l++
		} else {
			ans[ids[r]] = v
			r--
		}
	}

	return ans
}
