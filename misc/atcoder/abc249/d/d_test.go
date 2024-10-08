// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc249/submit?taskScreenName=abc249_d
// 对拍：https://atcoder.jp/contests/abc249/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc249_d&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [d]")
	testCases := [][2]string{
		{
			`3
			6 2 3`,
			`2`,
		},
		{
			`1
			2`,
			`0`,
		},
		{
			`10
			1 3 2 4 6 8 2 2 3 7`,
			`62`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc249/tasks/abc249_d
