package _00

import (
	"bufio"
	"fmt"
	"io"
)

func CF827B(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n, k int
	fmt.Fscan(r, &n, &k)
	// k条链，(n-1)%k条链长度为(n-1)/k+1 其余为(n-1)/k
	ans := (n-1)/k*2 + min(2, (n-1)%k)
	fmt.Fprintln(w, ans)
	for i := 2; i <= k; i++ {
		fmt.Fprintln(w, 1, " ", i)
	}
	for i := k + 1; i <= n; i++ {
		fmt.Fprintln(w, i, " ", i-k)
	}

}
