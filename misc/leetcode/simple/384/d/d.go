package main

func countMatchingSubarrays(nums, pattern []int) (ans int) {
	m := len(pattern)
	btoi := func(a, b int) int {
		if a == b {
			return 0
		}
		if a > b {
			return 1
		}
		return -1
	}
	pattern = append(pattern, 2)
	for i := 1; i < len(nums); i++ {
		pattern = append(pattern, btoi(nums[i], nums[i-1]))
	}

	n := len(pattern)
	z := make([]int, n)
	l, r := 0, 0
	for i := 1; i < n; i++ {
		if i <= r {
			z[i] = min(z[i-l], r-i+1)
		}
		for i+z[i] < n && pattern[z[i]] == pattern[i+z[i]] {
			l, r = i, i+z[i]
			z[i]++
		}
	}

	for _, lcp := range z[m+1:] {
		if lcp == m {
			ans++
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
