// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [D]")
	testCases := [][2]string{
		{
			`4
			5 3
			BBWBW
			5 5
			BBWBW
			5 1
			BBWBW
			1 1
			W`,
			`1
			2
			0
			1
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://codeforces.com/contest/1690/problem/D
