package _022spring

import (
	"problems/testutil/leetcode"
	"testing"
)

func TestBuildBridge(t *testing.T) {
	examples := [][]string{
		{
			`10`, `[[1,2],[4,7],[8,9]]`,
			`3`,
		},
		{
			`10`, `[[1,5],[1,1],[10,10],[6,7],[7,8]]`,
			`10`,
		},
		{
			`5`, `[[1,2],[2,4]]`,
			`0`,
		},
	}

	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithExamples(t, buildBridge, examples, targetCaseNum); err != nil {
		t.Fatal(err)
	}
}

func TestGetMaxLayerSum(t *testing.T) {
	//examples := [][]string{
	//	{
	//		`[6,0,3,null,8]`,
	//		`11`,
	//	},
	//	{
	//		`[5,6,2,4,null,null,1,3,5]`,
	//		`9`,
	//	},
	//	{
	//		`[-5,1,7]`,
	//		`8`,
	//	},
	//}
	//
	//targetCaseNum := 0 // -1
	//if err := leetcode.RunLeetCodeFuncWithExamples(t, getMaxLayerSum, examples, targetCaseNum); err != nil {
	//	t.Fatal(err)
	//}

	targetCaseNum := 0 // -1
	if err := leetcode.RunLeetCodeFuncWithFile(t, getMaxLayerSum, "team_6.txt", targetCaseNum); err != nil {
		t.Fatal(err)
	}
}
