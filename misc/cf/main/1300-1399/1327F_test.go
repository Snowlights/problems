//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1327/F
// https://codeforces.com/problemset/status/1327/problem/F
func TestCF1327F(t *testing.T) {
	t.Log("Current test is [CF1327F]")
	testCases := [][2]string{
		{
			`4 3 2
			1 3 3
			3 4 6
			`,
			`3`,
		},
		{
			`5 2 3
			1 3 2
			2 5 0
			3 3 3
			`,
			`33`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1327F, testCases, 0)
}
