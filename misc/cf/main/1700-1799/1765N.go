//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1765N(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var T, k, j int
	var s, ans []byte
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &s, &k)
		ans, j = make([]byte, 0, len(s)), 0
		for i, b := range s[:k+1] {
			if b > '0' && b < s[j] {
				j = i
			}
		}
		ans = append(ans, s[j])
		k -= j
		for _, b := range s[j+1:] {
			for len(ans) > 1 && k > 0 && ans[len(ans)-1] > b {
				ans = ans[:len(ans)-1]
				k--
			}
			ans = append(ans, b)
		}
		ans = ans[:len(ans)-k]
		Fprintln(out, string(ans))
	}

}

func main() { CF1765N(os.Stdin, os.Stdout) }
