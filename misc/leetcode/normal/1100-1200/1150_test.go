package _100_1200

import (
	"problems/testutil/leetcode"
	"testing"
)

// 1172
func TestConstructor(t *testing.T) {
	examples := [][]string{
		{
			`["DinnerPlates","push","push","push","push","push","popAtStack","push","push", "popAtStack","popAtStack","pop","pop","pop","pop","pop"]`,
			`[[2],[1],[2],[3],[4],[5],[0],[20],[21],[0],[2],[],[],[],[],[]]`,
			`[null,null,null,null,null,null,2,null,null,20,21,5,4,3,1,-1]`,
		},
	}

	leetcode.RunLeetCodeFuncWithExamples(t, Constructor, examples, 0)
}
