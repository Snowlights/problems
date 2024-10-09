package _00

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
)

func CF837D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, k, ans int
	var v uint64
	fmt.Fscan(r, &n, &k)
	f := make([][]int, k+1)
	for i := range f {
		f[i] = make([]int, k*25+1)
		for j := range f[i] {
			f[i][j] = -1
		}
	}
	f[0][0] = 0
	for ; n > 0; n-- {
		fmt.Fscan(r, &v)
		cnt2 := bits.TrailingZeros64(v)
		cnt5 := 0
		for ; v%5 == 0; v /= 5 {
			cnt5++
		}
		for i := k; i > 0; i-- {
			for j := k * 25; j >= cnt5; j-- {
				if f[i-1][j-cnt5] >= 0 {
					f[i][j] = max(f[i][j], f[i-1][j-cnt5]+cnt2)
				}
			}
		}
	}
	for i, cnt2 := range f[k] {
		ans = max(ans, min(i, cnt2))
	}
	fmt.Fprintln(w, ans)
}
