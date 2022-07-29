package _281

import (
	"bufio"
	"fmt"
	"io"
)

func CF1281B(in io.Reader, out io.Writer) {

	r := bufio.NewReader(in)
	w := bufio.NewWriter(out)
	defer w.Flush()

	var n int
	var s, t []byte
	fmt.Fscan(r, &n)
	for ; n > 0; n-- {
		fmt.Fscan(r, &s, &t)
		last := make(map[byte]int)
		for i, b := range s {
			last[b] = i
		}
		for i, v := range s {
			if v == 'A' {
				continue
			}
			exit, idx := false, i
			for j := 'A'; j <= 'Z'; j++ {
				if last[byte(j)] > i && byte(j) < byte(v) {
					exit = true
					idx = last[byte(j)]
					break
				}
			}
			if exit {
				s[i], s[idx] = s[idx], s[i]
				break
			}
		}
		// fmt.Println(string(s), string(t))
		if string(s) < string(t) {
			fmt.Fprintln(w, string(s))
		} else {
			fmt.Fprintln(w, "---")
		}
	}

}
