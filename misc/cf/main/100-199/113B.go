//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

type node13 struct {
	son [26]uint32
	end bool
}

var memo13 = [2e6]node13{{}}

func CF113B(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	defer out.Flush()

	calcPi := func(s []byte) []int {
		pi := make([]int, len(s))
		match := 0
		for i := 1; i < len(pi); i++ {
			v := s[i]
			for match > 0 && s[match] != v {
				match = pi[match-1]
			}
			if s[match] == v {
				match++
			}
			pi[i] = match
		}
		return pi
	}
	kmpSearch := func(text, pattern []byte) (end []int) {
		pi := calcPi(pattern)
		match := 0
		for i := range text {
			v := text[i]
			for match > 0 && pattern[match] != v {
				match = pi[match-1]
			}
			if pattern[match] == v {
				match++
			}
			if match == len(pi) {
				end = append(end, i)
				match = pi[match-1]
			}
		}
		return
	}

	var s, p, q []byte
	Fscan(in, &s, &p, &q)
	sufEnd := make([]bool, len(s))
	for _, i := range kmpSearch(s, q) {
		sufEnd[i] = true
	}

	ans := 0
	nodes := memo13[:1]
	root := uint32(0)
	mn := max(len(p), len(q)) - 1
	for _, i := range kmpSearch(s, p) {
		i -= len(p) - 1
		o := root
		for j, b := range s[i:] {
			b -= 'a'
			if nodes[o].son[b] == 0 {
				nodes[o].son[b] = uint32(len(nodes))
				nodes = append(nodes, node13{})
			}
			o = nodes[o].son[b]
			if j >= mn && sufEnd[i+j] && !nodes[o].end {
				nodes[o].end = true
				ans++
			}
		}
	}
	Fprint(out, ans)

}

func main() { CF113B(os.Stdin, os.Stdout) }
