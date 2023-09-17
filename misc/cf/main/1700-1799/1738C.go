//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
)

func CF1738C(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, n, v int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		var odd, even int
		for i := 0; i < n; i++ {
			Fscan(in, &v)
			if v&1 == 1 {
				odd++
			} else {
				even++
			}
		}

		// 如果odd+1 >> 1 是偶数， alice必赢
		// 如果两者都为奇数，alice总能拿到偶数个奇数，
		if (odd+1)>>1&1 == 0 || (odd&1 == 1 && even&1 == 1) {
			Fprintln(out, "Alice")
		} else {
			Fprintln(out, "Bob")
		}
	}
}

// func main() { CF1738C(os.Stdin, os.Stdout) }
