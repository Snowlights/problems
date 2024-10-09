package _50

import (
	"bufio"
	"fmt"
	"io"
	"math/bits"
)

func CF484A(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T, left, right int
	for fmt.Fscan(r, &T); T > 0; T-- {
		fmt.Fscan(r, &left, &right)
		base, ans, bi := 1, left, bits.OnesCount(uint(left))
		for left|base <= right {
			left |= base
			base <<= 1
			if val := bits.OnesCount(uint(left)); val > bi {
				ans = left
			}
		}
		fmt.Fprintln(w, ans)
	}

}
