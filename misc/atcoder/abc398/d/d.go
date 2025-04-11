package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func run(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	type pair struct{ x, y int }
	dir4 := []pair{'W': {1, 0}, 'E': {-1, 0}, 'S': {0, -1}, 'N': {0, 1}}
	var man, fire pair
	var s []byte
	Fscan(in, &s, &man.y, &man.x, &s)
	has := map[pair]bool{fire: true}
	for i, b := range s {
		d := dir4[b]
		man.x += d.x
		man.y += d.y
		fire.x += d.x
		fire.y += d.y
		has[fire] = true
		if has[man] {
			s[i] = '1'
		} else {
			s[i] = '0'
		}
	}
	Fprintf(out, "%s", s)

}

func main() { run(os.Stdin, os.Stdout) }
