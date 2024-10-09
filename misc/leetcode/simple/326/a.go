package simple_326

import (
	"sort"
	"strconv"
)

// 1
func countDigits(num int) int {
	s := strconv.Itoa(num)
	h := make(map[byte]int)
	for i := range s {
		h[s[i]]++
	}
	ans := 0
	for k, v := range h {
		if num%int(k-'0') == 0 {
			ans += v
		}
	}
	return ans
}

// 2
func distinctPrimeFactors(nums []int) int {
	h := make(map[int]bool)
	for _, v := range nums {
		i := 2
		for i*i <= v {
			for v%i == 0 {
				v /= i
				h[i] = true
			}
			i++
		}
		if v > 1 {
			h[v] = true
		}
	}

	return len(h)
}

// 3
func minimumPartition(s string, k int) int {
	sum, ans := 0, 0
	for i := range s {
		sum = sum*10 + int(s[i]-'0')
		if sum > k {
			sum = int(s[i] - '0')
			ans++
		}
		if int(s[i]-'0') > k {
			return -1
		}
	}
	if sum > 0 {
		ans++
	}
	return ans
}

// 4
const mx = 1e6 + 1

var primes = make([]int, 0, 78500)

func init() {
	np := [mx]bool{}
	for i := 2; i < mx; i++ {
		if !np[i] {
			primes = append(primes, i)
			for j := i * i; j < mx; j += i {
				np[j] = true
			}
		}
	}
	primes = append(primes, mx, mx) // 保证下面下标不会越界
}

func closestPrimes(left, right int) []int {
	p, q := -1, -1
	for i := sort.SearchInts(primes, left); primes[i+1] <= right; i++ {
		if p < 0 || primes[i+1]-primes[i] < q-p {
			p, q = primes[i], primes[i+1]
		}
	}
	return []int{p, q}
}
