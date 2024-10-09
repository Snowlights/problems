//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF628B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	var s []byte
	Fscan(in, &s)
	ans := 0
	for i, c := range s {
		if c%4 == 0 {
			ans++
		}
		if i > 0 && (s[i-1]*2+c)%4 == 0 {
			ans += i
		}
	}
	Fprint(out, ans)

}

func main() { CF628B(os.Stdin, os.Stdout) }
