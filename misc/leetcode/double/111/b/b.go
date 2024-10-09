package main

func canMakeSubsequence(str1 string, str2 string) bool {
	idx, n := 0, len(str2)
	for i := range str1 {
		if idx < n && (str2[idx] == str1[i] || str2[idx] == str1[i]+1 || str1[i]-25 == str2[idx]) {
			idx++
		}
	}
	return idx == n
}
