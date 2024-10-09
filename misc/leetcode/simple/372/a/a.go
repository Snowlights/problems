package main

func findMinimumOperations(s1, s2, s3 string) int {
	n := min(len(s1), len(s2), len(s3))
	i := 0
	for i < n && s2[i] == s1[i] && s3[i] == s1[i] {
		i++
	}
	if i == 0 {
		return -1
	}
	return len(s1) + len(s2) + len(s3) - i*3
}

func min(a ...int) int {
	res := a[0]
	for _, v := range a {
		if v < res {
			res = v
		}
	}
	return res
}
