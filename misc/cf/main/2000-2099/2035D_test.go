//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/2035/D
// https://codeforces.com/problemset/status/2035/problem/D
func TestCF2035D(t *testing.T) {
	t.Log("Current test is [CF2035D]")
	testCases := [][2]string{
		{
			`3
			10
			1 2 3 4 5 6 7 8 9 10
			11
			1 6 9 4 7 4 4 10 3 2 3
			4
			527792568 502211460 850237282 374773208`,
			`1 3 8 13 46 59 126 149 1174 1311 
			1 7 22 26 70 74 150 1303 1306 1308 1568 
			527792568 83665723 399119771 773892979`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF2035D, testCases, 0)
}
