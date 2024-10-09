package _800_1900

import "sort"

// 1802
func maxValue1802(n int, index int, maxSum int) int {

	valsum := func(i, n int) int {
		if n > i {
			return (i+i*i)/2 + n - i
		} else if n < i {
			return (i+i*i)/2 - (((i - n) + (i-n)*(i-n)) / 2)
		} else {
			return (i + i*i) / 2
		}
	}

	return sort.Search(1e9+1, func(i int) bool {
		sum := valsum(i-1, index) + valsum(i-1, n-1-index) + i
		return sum > maxSum
	}) - 1
}

// 1818
func minAbsoluteSumDiff(nums1 []int, nums2 []int) int {
	vec := append(sort.IntSlice{}, nums1...)
	sort.Ints(vec)

	sum, maxn, n := 0, 0, len(nums1)
	for i, v := range nums2 {
		diff := abs(nums1[i] - v)

		sum += diff
		j := vec.Search(v)
		if j < n {
			maxn = max(maxn, diff-(vec[j]-v))
		}
		if j > 0 {
			maxn = max(maxn, diff-(v-vec[j-1]))
		}
	}

	return (sum - maxn) % (1e9 + 7)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 1823
func findTheWinner(n int, k int) int {
	q := make([]int, n)
	for i := range q {
		q[i] = i + 1
	}
	for len(q) != 1 {
		idx := (k - 1) % len(q)
		q = append(q[idx+1:], q[0:idx]...)
	}
	return q[0]
}
