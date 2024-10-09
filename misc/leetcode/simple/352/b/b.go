package main

const mx = 1e6

var primes []int
var isP = [mx + 1]bool{}

func init() {
	for i := 2; i <= mx; i++ {
		isP[i] = true
	}
	for i := 2; i <= mx; i++ {
		if isP[i] {
			primes = append(primes, i)
			for j := i * i; j <= mx; j += i {
				isP[j] = false
			}
		}
	}
}

func findPrimePairs(n int) (ans [][]int) {
	for _, x := range primes {
		y := n - x
		if y < x {
			break
		}
		if isP[y] {
			ans = append(ans, []int{x, y})
		}
	}
	return
}
