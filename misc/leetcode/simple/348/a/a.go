package main

func minimizedStringLength(s string) (ans int) {
	h := make(map[byte]struct{})
	for _, v := range s {
		h[byte(v)] = struct{}{}
	}
	return len(h)
}
