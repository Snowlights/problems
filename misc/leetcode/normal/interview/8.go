package interview

// 08. 14
// https://www.luogu.com.cn/blog/BreakPlus/ou-jian-dp-zong-jie-ti-xie
// 区间DP模版题 https://leetcode.cn/problems/boolean-evaluation-lcci/
func countEval(s string, result int) int {
	n := len(s)
	dp := make([][][2]int, n)
	for i := range dp {
		dp[i] = make([][2]int, n)
	}

	for i := n - 1; i >= 0; i -= 2 {
		for j := i; j < n; j++ {
			if i == j {
				if s[i] == '0' {
					dp[i][j][0]++
				} else {
					dp[i][j][1]++
				}
				continue
			}

			/*
				操作符s[k]将s[i:j+1]分为s[i:k]和s[k+1:j+1]分为前后两块
				遍历各自dp[i][k-1][0/1]和dp[k+1][j][0/1]共四种情况，
				累计dp[i][j][0/1]的值
			*/

			for k := i + 1; k < j; k += 2 {
				for a := 0; a <= 1; a++ {
					for b := 0; b <= 1; b++ {
						if doBoolOp(a, b, s[k]) == 0 {
							dp[i][j][0] += dp[i][k-1][a] * dp[k+1][j][b]
						} else {
							dp[i][j][1] += dp[i][k-1][a] * dp[k+1][j][b]
						}
					}
				}
			}
		}
	}

	return dp[0][n-1][result]
}

func doBoolOp(arg1, arg2 int, operator byte) int {
	if operator == '&' {
		return arg1 & arg2
	} else if operator == '|' {
		return arg1 | arg2
	} else {
		return arg1 ^ arg2
	}
}
