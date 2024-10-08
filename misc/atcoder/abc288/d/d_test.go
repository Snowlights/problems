// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc288/submit?taskScreenName=abc288_d
// 对拍：https://atcoder.jp/contests/abc288/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc288_d&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{
			`7 3
			3 -1 1 -2 2 0 5
			2
			1 6
			2 7`,
			`Yes
			No`,
		},
		{
			`20 4
			-19 -66 -99 16 18 33 32 28 26 11 12 0 -16 4 21 21 37 17 55 -19
			5
			13 16
			4 11
			3 12
			13 18
			4 10`,
			`No
			Yes
			No
			Yes
			No`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc288/tasks/abc288_d
