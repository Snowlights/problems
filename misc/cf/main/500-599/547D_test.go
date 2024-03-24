//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/547/D
// https://codeforces.com/problemset/status/547/problem/D
func TestCF547D(t *testing.T) {
	t.Log("Current test is [CF547D]")
	testCases := [][2]string{
		{
			`4
			1 1
			1 2
			2 1
			2 2
			`,
			`brrb`,
		},
		{
			`3
			1 1
			1 2
			2 1
			`,
			`brr`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF547D, testCases, 0)
}
