// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/abc162/submit?taskScreenName=abc162_b
// 对拍：https://atcoder.jp/contests/abc162/submissions?f.LanguageName=Go&f.Status=AC&f.Task=abc162_b&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [b]")
	testCases := [][2]string{
		{
			`15`,
			`60`,
		},
		{
			`1000000`,
			`266666333332`,
		},
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}

// https://atcoder.jp/contests/abc162/tasks/abc162_b
