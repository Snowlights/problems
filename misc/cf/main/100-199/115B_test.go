//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/115/B
// https://codeforces.com/problemset/status/115/problem/B
func TestCF115B(t *testing.T) {
	t.Log("Current test is [CF115B]")
	testCases := [][2]string{
		{
			`4 5
			GWGGW
			GGWGG
			GWGGG
			WGGGG`,
			`11`,
		},
		{
			`3 3
			GWW
			WWW
			WWG`,
			`7`,
		},
		{
			`1 1
			G`,
			`0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF115B, testCases, 0)
}
