// Code generated by copypasta/gen/atcoder/generator_test.go
package main

import (
	"problems/testutil/codeforces"
	"testing"
)

// 提交：https://atcoder.jp/contests/practice2/submit?taskScreenName=practice2_l
// 对拍：https://atcoder.jp/contests/practice2/submissions?f.LanguageName=Go&f.Status=AC&f.Task=practice2_l&orderBy=source_length
func Test_run(t *testing.T) {
	t.Log("Current test is [l]")
	testCases := [][2]string{
		{
			`5 5
			0 1 0 0 1
			2 1 5
			1 3 4
			2 2 5
			1 1 3
			2 1 2`,
			`2
			0
			1`,
		},
		
	}
	codeforces.AssertEqualStringCase(t, run, testCases, 0)
}
// https://atcoder.jp/contests/practice2/tasks/practice2_l
