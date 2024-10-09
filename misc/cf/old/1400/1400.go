package _400

import (
	"bufio"
	"fmt"
	"io"
)

func CF1400D(in io.Reader, out io.Writer) {
	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var T, n int
	// 6
	// 1 3 3 1 2 3
	for fmt.Fscan(r, &T); T > 0; T-- {
		fmt.Fscan(r, &n)
		a := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(r, &a[i])
		}
		ans := int64(0)
		h := make(map[int]int64)
		h[a[0]]++
		for j := 1; j < n-2; j++ {
			v := a[j]
			exit := h[a[j+1]] // k相等的个数
			// l
			for _, l := range a[j+2:] {
				if l == v {
					ans += exit
				}
				exit += h[l]
			}
			h[v]++
		}

		fmt.Fprintln(w, ans)
	}

}
