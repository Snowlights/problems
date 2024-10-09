package test

import (
	"problems/testutil/codeforces"
	"testing"
)

func TestCF1775E(t *testing.T) {
	// 通过下载测试用例来判断用例是否通过
	//dir, _ := filepath.Abs(".")
	//codeforces.AssertEqualFileCaseWithName(t, CF1775E, dir, "in*.txt", "ans*.txt", 0)
	//

	// just copy from website
	//	rawText := `
	//inputCopy
	//5
	//3
	//1 2 -3
	//5
	//1 0 0 -1 -1
	//6
	//2 -4 3 -5 4 1
	//5
	//1 -1 1 -1 1
	//7
	//0 0 0 0 0 0 0
	//outputCopy
	//3
	//2
	//6
	//1
	//0`
	//	codeforces.AssertEqualCase(t, CF1775E, rawText, 0)
	var customTestCases = [][2]string{
		{
			`5
			3
			1 2 -3
			5
			1 0 0 -1 -1
			6
			2 -4 3 -5 4 1
			5
			1 -1 1 -1 1
			7
			0 0 0 0 0 0 0`,
			`3
			2
			6
			1
			0`,
		},
	}
	codeforces.AssertEqualStringCase(t, CF1775E, customTestCases, 0)

}
