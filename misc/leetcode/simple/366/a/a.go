package main

func differenceOfSums(n int, m int) (ans int) {
	k := n / m
	return n*(n+1)/2 - k*(k+1)*m
}
