// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc162/submit?taskScreenName=abc162_c
// 对拍：https://atcoder.jp/contests/abc162/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc162_c&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [c]")
	testCases := [][2]string{
		{
			`2`,
			`9`,
		},
		{
			`200`,
			`10813692`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc162/tasks/abc162_c