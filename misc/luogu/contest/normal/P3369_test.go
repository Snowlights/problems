//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://www.luogu.com.cn/problem/P3369

func TestP3369(t *testing.T) {
	t.Log("Current test is [P3369]")
	testCases := [][2]string{
		{
			`10
			1 106465
			4 1
			1 317721
			1 460929
			1 644985
			1 84185
			1 89851
			6 81968
			1 492737
			5 493598
			`,
			`106465
			84185
			492737`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}
