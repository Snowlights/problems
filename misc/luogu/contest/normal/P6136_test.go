//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://www.luogu.com.cn/problem/P6136

func TestP6136(t *testing.T) {
	t.Log("Current test is [P6136]")
	testCases := [][2]string{
		{
			`6 7
			1 1 4 5 1 4
			2 1
			1 9
			4 1
			5 8
			3 13
			6 7
			1 4
			`,
			`6`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}
