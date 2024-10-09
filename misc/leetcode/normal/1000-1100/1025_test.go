package _000_1100

import (
	"problems/testutil/leetcode"
	"testing"
)

// 1048
func TestLongestStrChain(t *testing.T) {
	examples := [][]string{
		{
			`["a","b","ba","bca","bda","bdca"]`,
			`4`,
		},
		{
			`["xbc","pcxbcf","xb","cxbc","pcxbc"]`,
			`5`,
		},
		{
			`["abcd","dbqca"]`,
			`1`,
		},
	}
	leetcode.RunLeetCodeFuncWithExamples(t, longestStrChain, examples, 0)
}
