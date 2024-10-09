package main

func makeSmallestPalindrome(s string) (ans string) {
	sb := []byte(s)
	n := len(s)
	for i := 0; i < n/2; i++ {
		if sb[i] == sb[n-1-i] {
			continue
		}
		if sb[i] > sb[n-1-i] {
			sb[i] = sb[n-1-i]
		} else {
			sb[n-1-i] = sb[i]
		}
	}

	return string(sb)
}
