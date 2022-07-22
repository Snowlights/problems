package _709

import (
	"fmt"
	"io"
)

func CF1709A(in io.Reader, out io.Writer) {

	var n int
	fmt.Fscan(in, &n)
	for i := 0; i < n; i++ {
		var k int
		fmt.Fscan(in, &k)
		arr := make([]int, 3)
		for j := 0; j < 3; j++ {
			fmt.Fscan(in, &arr[j])
		}

		now := 0
		for k != 0 {
			k = arr[k-1]
			now++
		}
		if now == 3 {
			fmt.Fprintln(out, "YES")
		} else {
			fmt.Fprintln(out, "NO")
		}
	}

}
