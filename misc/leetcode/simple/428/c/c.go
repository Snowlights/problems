package main

func beautifulSplits(nums []int) (ans int) {
	n := len(nums)
	// lcp[i][j] 表示 s[i:] 和 s[j:] 的最长公共前缀
	lcp := make([][]int, n+1)
	for i := range lcp {
		lcp[i] = make([]int, n+1)
	}
	for i := n - 1; i >= 0; i-- {
		for j := n - 1; j >= i; j-- {
			if nums[i] == nums[j] {
				lcp[i][j] = lcp[i+1][j+1] + 1
			}
		}
	}

	for i := 1; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if i <= j-i && lcp[0][i] >= i ||
				j-i <= n-j && lcp[i][j] >= j-i {
				ans++
			}
		}
	}
	return
}

func calcZ(s []int) []int {
	n := len(s)
	z := make([]int, n)
	boxL, boxR := 0, 0 // z-box 左右边界
	for i := 1; i < n; i++ {
		if i <= boxR {
			z[i] = min(z[i-boxL], boxR-i+1)
		}
		for i+z[i] < n && s[z[i]] == s[i+z[i]] {
			boxL, boxR = i, i+z[i]
			z[i]++
		}
	}
	return z
}

func beautifulSplits2(nums []int) (ans int) {
	n := len(nums)
	z0 := calcZ(nums)
	for i := 1; i < n-1; i++ {
		z := calcZ(nums[i:])
		for j := i + 1; j < n; j++ {
			if i <= j-i && z0[i] >= i ||
				j-i <= n-j && z[j-i] >= j-i {
				ans++
			}
		}
	}
	return
}
