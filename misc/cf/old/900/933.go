package _00

import (
	"bufio"
	"fmt"
	"io"
)

// 		f[1]+=(a==1);
//		f[2]=max(f[1],f[2]+(a==2));
//		f[3]=max(f[2],f[3]+(a==1));
//		f[4]=max(f[3],f[4]+(a==2));

func CF933A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()
	var n, v int64
	dp := make([]int, 4)
	for fmt.Fscan(r, &n); n > 0; n-- {
		fmt.Fscan(r, &v)
		one, two := 0, 0
		if v == 1 {
			one = 1
		} else if v == 2 {
			two = 1
		}
		dp[0] += one
		dp[1] = max(dp[0], dp[1]+two)
		dp[2] = max(dp[1], dp[2]+one)
		dp[3] = max(dp[2], dp[3]+two)
	}

	fmt.Fprintln(w, dp[3])
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
