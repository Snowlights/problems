// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

func Test_run(t *testing.T) {
	t.Log("Current test is [A]")
	testCases := [][2]string{
		{
			`6
			11
			6
			10
			100000
			7
			8`,
			`4 5 2
			2 3 1
			4 5 1
			33334 33335 33331
			2 4 1
			3 4 1
			`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://codeforces.com/contest/1690/problem/A