package main

func incremovableSubarrayCount(a []int) int {
	n, i := len(a), 0
	for i < n-1 && a[i] < a[i+1] {
		i++
	}
	if i == n-1 {
		return n * (n + 1) / 2
	}
	j, ans := n-1, i+2
	for ; j == n-1 || a[j] < a[j+1]; j-- {
		for i >= 0 && a[i] >= a[j] {
			i--
		}
		ans += i + 2
	}
	return ans
}
