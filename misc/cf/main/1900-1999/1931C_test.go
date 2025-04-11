//go:build main
// +build main

package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// https://codeforces.com/problemset/problem/1931/C
// https://codeforces.com/problemset/status/1931/problem/C
func TestCF1931C(t *testing.T) {
	t.Log("Current test is [CF1931C]")
	testCases := [][2]string{
		{
			`11
			1 1 1 1
			1 2 5 10
			4 6 100 200
			900000 900000 900000 900000
			0 0 0 0
			0 0 566 239
			1 0 0 0
			100 0 100 0
			0 0 0 4
			5 5 0 2
			5 4 0 5`,
			`4
			66
			0
			794100779
			1
			0
			1
			0
			1
			36
			126`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1931C, testCases, 0)
}
