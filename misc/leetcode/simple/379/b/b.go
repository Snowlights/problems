package main

func minMovesToCaptureTheQueen(a, b, c, d, e, f int) int {
	notInMiddle := func(l, m, r int) bool {
		return m < min(l, r) || m > max(r, l)
	}
	if (a == e && (c != e || notInMiddle(b, d, f))) ||
		(b == f && (d != f || notInMiddle(a, c, e))) ||
		(c+d == e+f && (a+b != e+f || notInMiddle(c, a, e))) ||
		(c-d == e-f && (a-b != e-f || notInMiddle(c, a, e))) {
		return 1
	}
	return 2
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
