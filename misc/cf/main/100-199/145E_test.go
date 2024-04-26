//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/145/E
// https://codeforces.com/problemset/status/145/problem/E
func TestCF145E(t *testing.T) {
	t.Log("Current test is [CF145E]")
	testCases := [][2]string{
		{
			`2 3
			47
			count
			switch 1 2
			count`,
			`2
			1`,
		},
		{
			`3 5
			747
			count
			switch 1 1
			count
			switch 1 3
			count`,
			`2
			3
			2`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, CF145E, testCases, 0)
}
