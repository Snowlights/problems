//go:build main
// +build main

package main

import (
	"bufio"
	. "fmt"
	"io"
	"os"
)

func CF1730E(_r io.Reader, out io.Writer) {
	in := bufio.NewReader(_r)
	const mx = 1000001
	divisors := [mx][]uint32{}
	for i := uint32(1); i < mx; i++ {
		for j := i * 2; j < mx; j += i {
			divisors[j] = append(divisors[j], i)
		}
	}

	var T, n int
	for Fscan(in, &T); T > 0; T-- {
		Fscan(in, &n)
		a := make([]int, n)
		pos := [mx][]int{}
		for i := range a {
			Fscan(in, &a[i])
			pos[a[i]] = append(pos[a[i]], i)
		}

		leftHi := make([]int, n)  // >= a[i]
		rightHi := make([]int, n) // > a[i]
		leftLo := make([]int, n)  // < a[i]
		rightLo := make([]int, n) // < a[i]
		s := []int{-1}
		t := []int{-1}
		for i, v := range a {
			for len(s) > 1 && v > a[s[len(s)-1]] {
				rightHi[s[len(s)-1]] = i
				s = s[:len(s)-1]
			}
			leftHi[i] = s[len(s)-1]
			s = append(s, i)

			for len(t) > 1 && v <= a[t[len(t)-1]] {
				t = t[:len(t)-1]
			}
			leftLo[i] = t[len(t)-1]
			t = append(t, i)
		}
		for _, i := range s[1:] {
			rightHi[i] = n
		}

		t = append(t[:0], n)
		for i := n - 1; i >= 0; i-- {
			v := a[i]
			for len(t) > 1 && v <= a[t[len(t)-1]] {
				t = t[:len(t)-1]
			}
			rightLo[i] = t[len(t)-1]
			t = append(t, i)
		}

		ans := 0
		for i, v := range a {
			ans += min(rightHi[i], rightLo[i]) - i
			for _, d := range divisors[v] {
				l, r := leftHi[i], rightHi[i]
				ps := pos[d]
				if len(ps) > 0 && ps[0] < i {
					x := ps[0]
					ps = ps[1:]
					ans += get(x-max(l, leftLo[x]), min(r, rightLo[x])-i)
					l = max(l, x)
				}
				if len(ps) > 0 {
					x := ps[0]
					ans += get(i-max(l, leftLo[x]), min(r, rightLo[x])-x)
				}
			}
			if len(pos[v]) > 1 && pos[v][1] == i {
				pos[v] = pos[v][1:]
			}
		}
		Fprintln(out, ans)
	}
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func get(x, y int) int {
	if x <= 0 || y <= 0 {
		return 0
	}
	return x * y
}

func main() { CF1730E(os.Stdin, os.Stdout) }
